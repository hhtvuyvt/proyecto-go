import {
    test,
    expect
} from "@playwright/test";

test.describe(
    "CRUD libros",
    () => {

        test.beforeEach(
            async ({ page }) => {

                await page.goto("/");

                await page.fill(
                    "#loginUsername",
                    "admin",
                );

                await page.fill(
                    "#loginPassword",
                    "admin123",
                );

                await page.click(
                    'button[type="submit"]',
                );

                await expect(
                    page.locator("#appPanel"),
                ).toBeVisible();

            },
        );

        test(
            "crear libro",
            async ({ page }) => {

                await page.fill(
                    "#title",
                    "Libro E2E",
                );

                await page.fill(
                    "#author",
                    "Autor E2E",
                );

                await page.click(
                    "#saveButton",
                );

                await expect(

                    page
                        .locator(".book")
                        .filter({
                            hasText: "Libro E2E",
                        }),

                ).toHaveCount(1);

            },
        );

        test(
            "editar mantiene datos",
            async ({ page }) => {

                const book =
                    page
                        .locator(".book")
                        .filter({
                            hasText: "Libro E2E",
                        })
                        .first();

                await expect(book)
                    .toBeVisible();

                await book
                    .getByText("Editar")
                    .click();

                await page.fill(
                    "#title",
                    "Libro cambiado",
                );

                await page.click(
                    "#saveButton",
                );

                await expect(

                    page
                        .locator(".book")
                        .filter({
                            hasText: "Libro cambiado",
                        }),

                ).toContainText(
                    "Autor E2E",
                );

            },
        );

        test(
            "borrar libro",
            async ({ page }) => {

                const book =
                    page
                        .locator(".book")
                        .filter({
                            hasText: "Libro cambiado",
                        })
                        .first();

                await expect(book)
                    .toBeVisible();

                page.once(
                    "dialog",
                    dialog => dialog.accept(),
                );

                console.log(
                    "Click borrar",
                );

                await book
                    .getByText("Borrar")
                    .click();

                // Esperar a que desaparezca del DOM
                await expect(

                    page
                        .locator(".book")
                        .filter({
                            hasText: "Libro cambiado",
                        }),

                ).toHaveCount(0);

            },
        );

    },
);