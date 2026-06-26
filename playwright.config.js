import { defineConfig } from "@playwright/test";


export default defineConfig({

    testDir: "./e2e",

    timeout: 30000,


    webServer: {

        command: "go run main.go",

        url: "http://localhost:8080",

        reuseExistingServer: false,

        env: {

            PORT: "8080",

            JWT_SECRET: "test_secret_key",

            DB_PATH: "./data/test.db"

        }

    },


    use: {

        baseURL: "http://localhost:8080",

        headless: true,

        screenshot: "only-on-failure",

        trace: "retain-on-failure"

    }

});