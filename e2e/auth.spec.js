import { test, expect } from "@playwright/test";

test(
    "login genera cookie de autenticación",
    async ({ request }) => {

        const response =
            await request.post(
                "/api/login",
                {
                    data: {
                        username: "admin",
                        password: "admin123",
                    },
                },
            );

        console.log({
            status: response.status(),
            body: await response.text(),
        });

        expect(
            response.ok(),
        ).toBeTruthy();

        const cookie =
            response.headers()["set-cookie"];

        expect(
            cookie,
        ).toBeDefined();

        expect(
            cookie,
        ).toContain("token=");

    },
);