import { test, expect } from "@playwright/test";


const libro = {
    title: "Libro E2E " + Date.now(),
    author: "Autor original",
    isbn: "ISBN-12345",
    image: "https://via.placeholder.com/200x300"
};



test.describe(
    "CRUD libros",
    ()=>{


        test.beforeEach(
            async({page})=>{

                await page.goto("/");

                await expect(
                    page.locator("#book-form")
                )
                    .toBeVisible();

            }

        );





        test(
            "crear libro",
            async({page})=>{


                await page
                    .locator("#titulo")
                    .fill(libro.title);



                await page
                    .locator("#autor")
                    .fill(libro.author);



                await page
                    .locator("#isbn")
                    .fill(libro.isbn);



                await page
                    .locator("#image")
                    .fill(libro.image);




                await page
                    .getByTestId("saveButton")
                    .click();





                await expect(
                    page.locator(".book")
                        .filter({
                            hasText: libro.title
                        })
                )
                    .toBeVisible();



            });









        test(
            "editar mantiene datos",
            async({page})=>{



                const book =
                    page.locator(".book")
                        .filter({
                            hasText: libro.title
                        });



                await expect(book)
                    .toBeVisible();




                await book
                    .getByTestId("edit-book")
                    .click();





                await page
                    .locator("#titulo")
                    .fill(
                        "Libro cambiado E2E"
                    );




                await page
                    .getByTestId("saveButton")
                    .click();






                const updated =
                    page.locator(".book")
                        .filter({
                            hasText:"Libro cambiado E2E"
                        });




                await expect(updated)
                    .toContainText(
                        libro.author
                    );



                await expect(updated)
                    .toContainText(
                        libro.isbn
                    );



            });









        test(
            "borrar libro",
            async({page})=>{


                const book =
                    page.locator(".book")
                        .filter({
                            hasText:"Libro cambiado E2E"
                        });



                await expect(book)
                    .toBeVisible();





                await book
                    .getByTestId("delete-book")
                    .click();





                await expect(
                    page.locator(".book")
                        .filter({
                            hasText:"Libro cambiado E2E"
                        })
                )
                    .toHaveCount(0);



            });





    });