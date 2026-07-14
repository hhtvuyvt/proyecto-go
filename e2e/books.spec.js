import {
    test,
    expect
}
    from "@playwright/test";


test.describe(
    "CRUD libros",
    () => {

        // =========================================
        // Setup: Autenticación con cookies
        // =========================================

        test.beforeEach(
            async ({ page }) => {

                // Ir a la página principal
                await page.goto(
                    "/"
                );

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
                // Esperar a que se complete la solicitud de autenticación
                await Promise.all([
                    page.waitForResponse(
                        response =>
                            response.url().includes(
                                "/api/login"
                            ) && response.status() === 200
                    ),
                    page.click(
                        "#loginForm button[type='submit']"
                    )
                ]);

                // Esperar a que la aplicación cargue
                // (loginPanel se oculta, appPanel se muestra)
                await page.waitForSelector(
                    "#appPanel:not(.d-none)"
                );

                // Esperar a que carguen los libros
                await page.waitForSelector(
                    "#book-list"
                );

            }
        );

        // =========================================
        // Test: Crear libro
        // =========================================

        test(
            "crear libro",
            async ({ page }) => {

                // Verificar que el formulario es visible
                await expect(
                    page.locator(
                        "#book-form"
                    )
                )
                    .toBeVisible();

                // Verificar estado inicial del botón
                await expect(
                    page.locator(
                        "#saveButton"
                    )
                )
                    .toContainText(
                        "Agregar"
                    );

                // Llenar el formulario
                await page.fill(
                    "#title",
                    "Libro E2E Test"
                );

                await page.fill(
                    "#author",
                    "Autor E2E"
                );

                await page.fill(
                    "#ISBN",
                    "978-1234567890"
                );

                // Hacer clic en guardar y esperar respuesta de API
                await Promise.all([
                    page.waitForResponse(
                        response =>
                            response.url().includes(
                                "/api/books"
                            ) && response.request().method() === "POST"
                            && response.status() === 200
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
                            hasText: "Libro E2E Test"
                        }
                    );

                await expect(
                    book
                )
                    .toContainText(
                        "Libro E2E Test"
                    );

                await expect(
                    book
                )
                    .toContainText(
                        "Autor E2E"
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

                // Crear un libro primero
                await page.fill(
                    "#title",
                    "Libro Original"
                );

                await page.fill(
                    "#author",
                    "Autor Original"
                );

                await page.fill(
                    "#ISBN",
                    "111-1111111111"
                );

                await Promise.all([
                    page.waitForResponse(
                        response =>
                            response.url().includes(
                                "/api/books"
                            ) && response.request().method() === "POST"
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
                            hasText: "Libro Original"
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
                        "Libro Original"
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
                await page.fill(
                    "#title",
                    "Libro Modificado"
                );

                // Guardar cambios y esperar respuesta de API
                await Promise.all([
                    page.waitForResponse(
                        response =>
                            response.url().includes(
                                "/api/books"
                            ) && response.request().method() === "PUT"
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
                            hasText: "Libro Modificado"
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

                // Crear un libro primero
                await page.fill(
                    "#title",
                    "Libro para borrar"
                );

                await page.fill(
                    "#author",
                    "Autor Temporal"
                );

                await Promise.all([
                    page.waitForResponse(
                        response =>
                            response.url().includes(
                                "/api/books"
                            ) && response.request().method() === "POST"
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
                            hasText: "Libro para borrar"
                        }
                    )
                        .first();

                // Interceptar diálogo de confirmación
                page.on(
                    "dialog",
                    dialog => dialog.accept()
                );

                // Click en Borrar y esperar respuesta de API
                await Promise.all([
                    page.waitForResponse(
                        response =>
                            response.url().includes(
                                "/api/books"
                            ) && response.request().method() === "DELETE"
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
                            hasText: "Libro para borrar"
                        }
                    )
                )
                    .not
                    .toBeVisible();

            }
        );

    }
);
