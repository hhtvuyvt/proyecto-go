import { defineConfig } from "@playwright/test";


export default defineConfig({

testDir: "./e2e",

timeout: 30000,


webServer: {

command: "go run main.go",

url: "http://localhost:8080",

reuseExistingServer: true

},


use: {

baseURL: "http://localhost:8080",

headless: true,

screenshot: "only-on-failure",

trace: "retain-on-failure"

}

});