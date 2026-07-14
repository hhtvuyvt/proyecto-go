import {
    test,
    expect
}
    from "@playwright/test";


test.describe(
    "Autenticación",
    () => {

        test(
            "login genera token en cookie",
            async ({ request }) => {

                // POST con credenciales
                const response =
                    await request.post(
                        "/api/login",
                        {
                            data: {
                                username: "admin",
                                password: "admin"
                            }
                        }
                    );

                // ✅ Debe retornar 200 OK
                expect(
                    response.ok()
                )
                    .toBeTruthy();

                // ✅ La respuesta debe contener Set-Cookie header
                const setCookieHeader =
                    response.headers()["set-cookie"];

                expect(
                    setCookieHeader
                )
                    .toBeDefined();

                // ✅ La cookie debe llamarse "token"
                expect(
                    setCookieHeader
                )
                    .toContain(
                        "token="
                    );

                // ✅ La cookie debe ser HttpOnly
                expect(
                    setCookieHeader
                )
                    .toContain(
                        "HttpOnly"
                    );

            }
        );

        test(
            "login con credenciales incorrectas falla",
            async ({ request }) => {

                const response =
                    await request.post(
                        "/api/login",
                        {
                            data: {
                                username: "usuario_incorrecto",
                                password: "password_incorrecta"
                            }
                        }
                    );

                // ❌ Debe retornar 401 Unauthorized
                expect(
                    response.ok()
                )
                    .toBeFalsy();

                expect(
                    response.status()
                )
                    .toBe(
                        401
                    );

                // ✅ Debe contener mensaje de error
                const text =
                    await response.text();

                expect(
                    text
                )
                    .toContain(
                        "usuario o contraseña incorrectos"
                    );

            }
        );

        test(
            "logout borra la cookie de sesión",
            async ({ request }) => {

                // Primero hacer login
                await request.post(
                    "/api/login",
                    {
                        data: {
                            username: "admin",
                            password: "admin"
                        }
                    }
                );

                // Luego hacer logout
                const response =
                    await request.post(
                        "/api/logout"
                    );

                // ✅ Debe retornar 200 OK
                expect(
                    response.ok()
                )
                    .toBeTruthy();

                // ✅ Debe contener Set-Cookie para borrar el token
                const setCookieHeader =
                    response.headers()["set-cookie"];

                expect(
                    setCookieHeader
                )
                    .toBeDefined();

                // ✅ La cookie debe tener MaxAge=-1 para borrarla
                expect(
                    setCookieHeader
                )
                    .toContain(
                        "Max-Age=-1"
                    );

            }
        );

    }
);
