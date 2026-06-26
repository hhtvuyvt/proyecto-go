import { test, expect } from "@playwright/test";


test.describe("CRUD libros",()=>{


    test.beforeEach(async({page})=>{

        await page.goto("/");

    });




    test("crear libro", async({page})=>{


        await page.fill("#titulo","Libro E2E");

        await page.fill("#autor","Autor E2E");

        await page.fill("#isbn","111");

        await page.fill(
            "#imagen",
            "https://via.placeholder.com/150"
        );



        await page.click("#saveButton");



        await expect(
            page.locator(".book")
        )
            .toContainText("Libro E2E");

    });






    test("editar mantiene datos",async({page})=>{


        await page.fill("#titulo","Libro editar");

        await page.fill("#autor","Autor original");

        await page.fill("#isbn","222");

        await page.fill("#imagen","https://via.placeholder.com/150");



        await page.click("#saveButton");



        const book =
            page.locator(".book")
                .filter({
                    hasText:"Libro editar"
                });



        await book
            .getByTestId("edit-button")
            .click();



        await page.fill(
            "#titulo",
            "Libro cambiado"
        );



        await page.click("#saveButton");



        const updated =
            page.locator(".book")
                .filter({
                    hasText:"Libro cambiado"
                });



        await expect(updated)
            .toContainText("Autor original");



        await expect(updated)
            .toContainText("222");

    });






    test("borrar libro",async({page})=>{


        await page.fill("#titulo","Libro borrar");

        await page.fill("#autor","Eliminar");

        await page.fill("#isbn","333");

        await page.fill("#imagen","https://via.placeholder.com/150");



        await page.click("#saveButton");



        const book =
            page.locator(".book")
                .filter({
                    hasText:"Libro borrar"
                });



        await book
            .getByTestId("delete-button")
            .click();



        await expect(book)
            .not
            .toBeVisible();


    });


});