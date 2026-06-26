import {
    test,
    expect
}
    from "@playwright/test";


test.describe(
    "CRUD libros",
    ()=>{


        test(
            "crear libro",
            async({page})=>{


                await page.goto("/");



                await page.fill(
                    "#title",
                    "Libro E2E"
                );


                await page.fill(
                    "#author",
                    "Autor E2E"
                );


                await page.click(
                    "#saveButton"
                );



                await expect(
                    page.locator(".book")
                )
                    .toContainText(
                        "Libro E2E"
                    );


            });





        test(
            "editar mantiene datos",
            async({page})=>{


                await page.goto("/");



                const book=
                    page.locator(".book")
                        .filter(
                            {
                                hasText:"Libro E2E"
                            }
                        )
                        .first();



                await book
                    .getByText("Editar")
                    .click();



                await page.fill(
                    "#title",
                    "Libro cambiado"
                );



                await page.click(
                    "#saveButton"
                );



                const updated=
                    page.locator(".book")
                        .filter(
                            {
                                hasText:"Libro cambiado"
                            }
                        );



                await expect(updated)
                    .toContainText(
                        "Autor E2E"
                    );



            });






        test(
            "borrar libro",
            async({page})=>{


                await page.goto("/");



                const book=
                    page.locator(".book")
                        .filter(
                            {
                                hasText:"Libro cambiado"
                            }
                        )
                        .first();



                await book
                    .getByText("Borrar")
                    .click();



                await expect(book)
                    .not
                    .toBeVisible();



            });



    });