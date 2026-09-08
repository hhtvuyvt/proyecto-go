import { expect, test } from "@playwright/test";

test.describe("Upload", () => {
  test.beforeEach(async ({ page }) => {
    await page.goto("/");
    await page.fill("#loginUsername", "admin");
    await page.fill("#loginPassword", "admin123");
    await page.click('button[type="submit"]');
    await expect(page.locator("#appPanel")).toBeVisible();
  });

  test("permite seleccionar y subir una imagen de libro", async ({ page }) => {
    await expect(page.locator("#image")).toBeVisible();

    // Simula la selección de un archivo local en el input
    await page.setInputFiles("#image", "path/to/test-image.png");

    // Opcional: validar que el nombre del archivo aparezca o se previsualice
    const inputValue = await page.locator("#image").inputValue();
    expect(inputValue).toContain("test-image.png");
  });
});
