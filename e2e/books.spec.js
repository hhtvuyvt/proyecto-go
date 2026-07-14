import {
    test,
    expect
}
    from "@playwright/test";


test.describe(
    "CRUD libros",
    () => {

        // =========================================
        // Setup: Autenticación previo a cada test
        // =========================================

        test.beforeEach(
            async ({ page }) => {

                // Ir a la página principal
                await page.goto(
                    "/"
                );

                // Si no está autenticado, hacer login
                const loginPanel =
                    page.locator(
                        "#loginPanel"
                    );

                const isLoginVisible =
                    await loginPanel.isVisible();

                if (isLoginVisible) {

                    // Llenar formulario de login
                    await page.fill(
                        "#loginUsername",
                        "admin"
                    );

                    await page.fill(
                        "#loginPassword",
                        "admin"
                    );

                    // Hacer clic en el botón de login
                    // Esperar a que la aplicación sea visible (UI update)
                    await Promise.all([
                        page.waitForSelector(
                            "#appPanel:not(.d-none)",
                            { timeout: 10000 }
                        ),
                        page.click(
                            "#loginForm button[type='submit']"
                        )
                    ]);

                }

                // Esperar a que carguen los libros
                await page.waitForSelector(
                    "#book-list",
                    { timeout: 10000 }
                );

            }
        );

        // =========================================
        // Test: Crear libro
        // =========================================

        test(
            "crear libro",
            async ({ page }) => {

                // Llenar el formulario con datos únicos
                const uniqueTitle =
                    `Libro E2E ${Date.now()}`;

                await page.fill(
                    "#title",
                    uniqueTitle
                );

                await page.fill(
                    "#author",
                    "Autor E2E Test"
                );

                await page.fill(
                    "#ISBN",
                    "978-1234567890"
                );

                await page.fill(
                    "#image",
                    "https://via.placeholder.com/150"
                );

                // POST /api/books (PÚBLICO, pero con credentials)
                // Esperar respuesta y hacer clic
                await Promise.all([
                    page.waitForResponse(
                        response =>
                            response.url().endsWith("/api/books")
                            && response.request().method() === "POST"
                            && response.status() === 200,
                        { timeout: 10000 }
                    ),
                    page.click(
                        "#saveButton"
                    )
                ]);

                // Verificar que el libro aparece en la lista
                const book =
                    page.locator(
                        ".book"
                    ).filter(
                        {
                            hasText: uniqueTitle
                        }
                    );

                await expect(
                    book
                )
                    .toContainText(
                        uniqueTitle
                    );

                await expect(
                    book
                )
                    .toContainText(
                        "Autor E2E Test"
                    );

                await expect(
                    book
                )
                    .toContainText(
                        "978-1234567890"
                    );

                // Verificar que el formulario se limpió
                await expect(
                    page.locator(
                        "#title"
                    )
                )
                    .toHaveValue(
                        ""
                    );

                await expect(
                    page.locator(
                        "#author"
                    )
                )
                    .toHaveValue(
                        ""
                    );

                // Verificar que el botón volvió a "Agregar"
                await expect(
                    page.locator(
                        "#saveButton"
                    )
                )
                    .toContainText(
                        "Agregar"
                    );

            }
        );

        // =========================================
        // Test: Editar libro mantiene datos
        // =========================================

        test(
            "editar libro mantiene datos",
            async ({ page }) => {

                // Crear un libro primero (datos únicos)
                const originalTitle =
                    `Original ${Date.now()}`;

                await page.fill(
                    "#title",
                    originalTitle
                );

                await page.fill(
                    "#author",
                    "Autor Original"
                );

                await page.fill(
                    "#ISBN",
                    "111-1111111111"
                );

                // POST /api/books (PÚBLICO)
                await Promise.all([
                    page.waitForResponse(
                        response =>
                            response.url().endsWith("/api/books")
                            && response.request().method() === "POST"
                    ),
                    page.click(
                        "#saveButton"
                    )
                ]);

                // Buscar el libro que acabamos de crear
                const book =
                    page.locator(
                        ".book"
                    ).filter(
                        {
                            hasText: originalTitle
                        }
                    )
                        .first();

                // Click en botón Editar
                await book
                    .getByText(
                        "Editar"
                    )
                    .click();

                // Verificar que el formulario se llenó con los datos actuales
                await expect(
                    page.locator(
                        "#title"
                    )
                )
                    .toHaveValue(
                        originalTitle
                    );

                await expect(
                    page.locator(
                        "#author"
                    )
                )
                    .toHaveValue(
                        "Autor Original"
                    );

                await expect(
                    page.locator(
                        "#ISBN"
                    )
                )
                    .toHaveValue(
                        "111-1111111111"
                    );

                // Verificar que el botón cambió a "Guardar cambios"
                await expect(
                    page.locator(
                        "#saveButton"
                    )
                )
                    .toContainText(
                        "Guardar cambios"
                    );

                // Cambiar solo el título
                const modifiedTitle =
                    `Modificado ${Date.now()}`;

                await page.fill(
                    "#title",
                    modifiedTitle
                );

                // PUT /api/books/{id} (PROTEGIDA)
                // Esperar respuesta y hacer clic
                await Promise.all([
                    page.waitForResponse(
                        response =>
                            response.url().includes("/api/books/")
                            && response.request().method() === "PUT",
                        { timeout: 10000 }
                    ),
                    page.click(
                        "#saveButton"
                    )
                ]);

                // Verificar que se actualizó correctamente
                const updated =
                    page.locator(
                        ".book"
                    ).filter(
                        {
                            hasText: modifiedTitle
                        }
                    );

                await expect(
                    updated
                )
                    .toContainText(
                        "Autor Original"
                    );

                // El autor debe mantenerse sin cambios
                await expect(
                    updated
                )
                    .toContainText(
                        "111-1111111111"
                    );

                // Verificar que el botón volvió a "Agregar"
                await expect(
                    page.locator(
                        "#saveButton"
                    )
                )
                    .toContainText(
                        "Agregar"
                    );

                // Verificar que el formulario se limpió
                await expect(
                    page.locator(
                        "#title"
                    )
                )
                    .toHaveValue(
                        ""
                    );

            }
        );

        // =========================================
        // Test: Borrar libro
        // =========================================

        test(
            "borrar libro",
            async ({ page }) => {

                // Crear un libro primero (datos únicos)
                const bookToDelete =
                    `Eliminar ${Date.now()}`;

                await page.fill(
                    "#title",
                    bookToDelete
                );

                await page.fill(
                    "#author",
                    "Autor Temporal"
                );

                // POST /api/books (PÚBLICO)
                await Promise.all([
                    page.waitForResponse(
                        response =>
                            response.url().endsWith("/api/books")
                            && response.request().method() === "POST"
                    ),
                    page.click(
                        "#saveButton"
                    )
                ]);

                // Buscar el libro
                const book =
                    page.locator(
                        ".book"
                    ).filter(
                        {
                            hasText: bookToDelete
                        }
                    )
                        .first();

                // Interceptar diálogo de confirmación
                page.on(
                    "dialog",
                    dialog => dialog.accept()
                );

                // DELETE /api/books/{id} (PROTEGIDA)
                // Esperar respuesta y hacer clic
                await Promise.all([
                    page.waitForResponse(
                        response =>
                            response.url().includes("/api/books/")
                            && response.request().method() === "DELETE",
                        { timeout: 10000 }
                    ),
                    book
                        .getByText(
                            "Borrar"
                        )
                        .click()
                ]);

                // Verificar que desapareció de la lista
                await expect(
                    page.locator(
                        ".book"
                    ).filter(
                        {
                            hasText: bookToDelete
                        }
                    )
                )
                    .not
                    .toBeVisible();

            }
        );

    }
);
