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

<<<<<<< HEAD
    // Sube una imagen al endpoint existente de upload con la sesión autenticada
    const response = await page.request.post("/api/upload", {
      multipart: {
        image: {
          name: "test-image.png",
          mimeType: "image/png",
          buffer: Buffer.from("fake-image-bytes"),
        },
      },
    });
=======
    // Simula la selección de un archivo local en el input
    await page.fill("#image", "path/to/test-image.png");
>>>>>>> ff572c17a8505179bb124b4cee455c5fcf55861c

    expect(response.ok()).toBeTruthy();
    const data = await response.json();
    expect(data.path).toContain("test-image.png");

    // Asigna la ruta de la imagen subida al input del formulario
    await page.fill("#image", data.path);

    // Validar que el nombre del archivo aparezca en el input
    const inputValue = await page.locator("#image").inputValue();
    expect(inputValue).toContain("test-image.png");
  });
});
