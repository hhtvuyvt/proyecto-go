import { test, expect } from "@playwright/test";


test.describe("CRUD libros", ()=>{


    test("crear libro", async({page})=>{


        await page.goto("/");


        await page.fill(
            "#title",
            "Libro prueba E2E"
        );


        await page.fill(
            "#author",
            "Autor prueba"
        );


        await page.fill(
            "#isbn",
            "99999"
        );


        await page.click(
            "#saveButton"
        );



        await expect(
            page.locator(".book")
        )
            .toContainText(
                "Libro prueba E2E"
            );


    });






    test("editar mantiene información", async({page})=>{


        await page.goto("/");



        const book =
            page.locator(".book")
                .filter({
                    hasText:
                        "Libro prueba E2E"
                })
                .first();



        await book
            .getByText("Editar")
            .click();




        await page.fill(
            "#title",
            "Libro editado E2E"
        );



        await page.click(
            "#saveButton"
        );




        const updated =
            page.locator(".book")
                .filter({
                    hasText:
                        "Libro editado E2E"
                });



        await expect(updated)
            .toContainText(
                "Autor prueba"
            );


        await expect(updated)
            .toContainText(
                "99999"
            );



    });








    test("borrar libro", async({page})=>{


        await page.goto("/");



        const book =
            page.locator(".book")
                .filter({
                    hasText:
                        "Libro editado E2E"
                })
                .first();




        await book
            .getByText("Borrar")
            .click();




        await expect(
            page.locator(".book")
        )
            .not
            .toContainText(
                "Libro editado E2E"
            );


    });


});