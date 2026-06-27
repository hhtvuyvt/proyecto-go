import {
    defineConfig
}
    from "@playwright/test";



export default defineConfig({

    testDir:"./e2e",


    webServer: {
        command: "go run .",
        url: "http://localhost:8080",
        timeout: 120000,

        env: {
            ...process.env,
            E2E: "true",
            JWT_SECRET: "test-secret-key",
            DB_PATH: "./test.db"
        }
    },


    use:{

        baseURL:
            "http://localhost:8080",

        trace:"on-first-retry"

    }


});