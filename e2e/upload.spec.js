import {
    test,
    expect
} from "@playwright/test";

test.describe(
    "Upload",
    () => {

        test.beforeEach(
            async ({ page }) => {

                await page.goto("/");

                await page.fill(
                    "#loginUsername",
                    "admin"
                );

                await page.fill(
                    "#loginPassword",
                    "admin123"
                );

                await page.click(
                    'button[type="submit"]'
                );

                await expect(
                    page.locator("#appPanel")
                ).toBeVisible();

            }
        );

        test(
            "pagina carga sistema de imagen",
            async ({ page }) => {

                await expect(
                    page.locator("#image")
                ).toBeVisible();

            }
        );

    }
);