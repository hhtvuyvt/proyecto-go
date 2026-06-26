import { test, expect } from "@playwright/test";



test(
    "pagina carga sistema de imagen",
    async({page})=>{


        await page.goto("/");



        await expect(
            page.locator("#image")
        )
            .toBeVisible();



        await expect(
            page.locator("#book-list")
        )
            .toBeVisible();



    });