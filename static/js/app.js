//
// app.js
//

window.addEventListener(
    "DOMContentLoaded",
    async () => {

        const logged =
            await checkSession();

        if (logged) {

            showApplication();

            await loadBooks();

        } else {

            showLogin();

        }

    },
);