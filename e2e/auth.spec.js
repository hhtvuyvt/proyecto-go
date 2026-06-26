import { test, expect } from "@playwright/test";



test(
    "login genera token",
    async({request})=>{


        const response =
            await request.get(
                "/api/login"
            );



        expect(
            response.ok()
        )
            .toBeTruthy();



        const data =
            await response.json();



        expect(
            data.token
        )
            .toBeDefined();


    });