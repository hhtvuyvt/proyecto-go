// playwright.config.js
// noinspection JSUnusedGlobalSymbols

import { defineConfig, devices } from "@playwright/test";

/*
|--------------------------------------------------------------------------
| Configuración E2E - Playwright
|--------------------------------------------------------------------------
|
| Este archivo controla los tests end-to-end del frontend.
|
| Objetivos:
|
| 1. Levantar automáticamente el backend Go antes de probar.
| 2. Usar una base de datos independiente para tests.
| 3. Evitar modificar datos reales de producción.
| 4. Asegurar que el frontend pueda comunicarse con la API.
|
|--------------------------------------------------------------------------
*/


export default defineConfig({

    /*
    |--------------------------------------------------------------------------
    | Carpeta donde viven los tests E2E
    |--------------------------------------------------------------------------
    |
    | Los archivos aquí prueban comportamientos completos:
    |
    | - Crear libros
    | - Editar libros
    | - Eliminar libros
    | - Carga de imágenes
    |
    */
    testDir: "./e2e",



    /*
    |--------------------------------------------------------------------------
    | Tiempo máximo por prueba
    |--------------------------------------------------------------------------
    |
    | Si una operación del navegador tarda más que esto,
    | Playwright considera que la prueba falló.
    |
    */
    timeout: 30000,



    /*
    |--------------------------------------------------------------------------
    | Evidencia de fallos
    |--------------------------------------------------------------------------
    |
    | Cuando un test falla:
    |
    | - Guarda screenshot
    | - Guarda trace para depuración
    |
    | Esto permite revisar qué vio el navegador.
    |
    */
    use: {

        baseURL: "http://localhost:8080",

        screenshot: "only-on-failure",

        trace: "retain-on-failure"

    },



    /*
    |--------------------------------------------------------------------------
    | servidor antes de ejecutar pruebas
    |--------------------------------------------------------------------------
    |
    | Playwright inicia el backend Go automáticamente.
    |
    | Importante:
    |
    | Se pasan variables de entorno propias de pruebas.
    |
    | Esto evita usar:
    |
    | - .env de desarrollo
    | - base de datos real
    |
    */
    webServer: {

        command: "go run .",

        url: "http://localhost:8080",

        reuseExistingServer: false,


        env: {

            PORT: "8080",

            JWT_SECRET:
                "test-secret-key",

            DB_PATH:
                "./test.db",

            ADMIN_USERNAME:
                "admin",

            ADMIN_PASSWORD:
                "admin123"

        }

    },



    /*
    |--------------------------------------------------------------------------
    | Navegadores
    |--------------------------------------------------------------------------
    |
    | Actualmente se prueba Chromium.
    |
    | Se puede extender a Firefox y WebKit.
    |
    */
    projects: [

        {
            name: "chromium",
            use: {
                ...devices["Desktop Chrome"]
            }
        }

    ]

});