import {
    defineConfig
}
    from "@playwright/test";



export default defineConfig({

    testDir:"./e2e",


    webServer:{

        command:"go run .",

        url:
            "http://localhost:8080",

        reuseExistingServer:true,

        env:{

            JWT_SECRET:
                "test_secret"

        }

    },


    use:{

        baseURL:
            "http://localhost:8080",

        trace:"on-first-retry"

    }


});