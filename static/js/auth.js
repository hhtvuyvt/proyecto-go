//
// auth.js
//
// Gestiona toda la autenticación de la aplicación.
//

"use strict";

// =============================
// Elementos del DOM
// =============================

const loginPanel =
    document.getElementById(
        "loginPanel",
    );

const appPanel =
    document.getElementById(
        "appPanel",
    );

const userPanel =
    document.getElementById(
        "userPanel",
    );

const loginForm =
    document.getElementById(
        "loginForm",
    );

const usernameInput =
    document.getElementById(
        "loginUsername",
    );

const passwordInput =
    document.getElementById(
        "loginPassword",
    );

const loginError =
    document.getElementById(
        "loginError",
    );

const usernameLabel =
    document.getElementById(
        "usernameLabel",
    );

const logoutButton =
    document.getElementById(
        "logoutButton",
    );

// =============================
// Interfaz
// =============================

function showLogin() {

    loginPanel.classList.remove(
        "d-none",
    );

    appPanel.classList.add(
        "d-none",
    );

    userPanel.classList.add(
        "d-none",
    );

    clearLoginError();

    passwordInput.value = "";

    usernameInput.focus();

}

function showApplication() {

    loginPanel.classList.add(
        "d-none",
    );

    appPanel.classList.remove(
        "d-none",
    );

    userPanel.classList.remove(
        "d-none",
    );

}

function showLoginError(
    message,
) {

    loginError.textContent =
        message;

    loginError.classList.remove(
        "d-none",
    );

}

function clearLoginError() {

    loginError.textContent = "";

    loginError.classList.add(
        "d-none",
    );

}

// =============================
// API
// =============================

async function login(
    username,
    password,
) {

    clearLoginError();

    const response =
        await fetch(
            "/api/login",
            {

                method: "POST",

                credentials:
                    "same-origin",

                headers: {
                    "Content-Type":
                        "application/json",
                },

                body:
                    JSON.stringify({

                        username,

                        password,

                    }),

            },
        );

    if (!response.ok) {

        showLoginError(
            "Usuario o contraseña incorrectos.",
        );

        return false;

    }

    usernameLabel.textContent =
        username;

    return true;

}

async function logout() {

    await fetch(
        "/api/logout",
        {

            method: "POST",

            credentials:
                "same-origin",

        },
    );

    showLogin();

}

async function checkSession() {

    try {

        const response =
            await fetch(
                "/api/me",
                {

                    credentials:
                        "same-origin",

                },
            );

        if (!response.ok) {

            return false;

        }

        const user =
            await response.json();

        usernameLabel.textContent =
            user.username;

        return true;

    } catch {

        return false;

    }

}

// =============================
// Eventos
// =============================

loginForm.addEventListener(

    "submit",

    async function (
        event,
    ) {

        event.preventDefault();

        const username =
            usernameInput.value.trim();

        const password =
            passwordInput.value;

        const ok =
            await login(
                username,
                password,
            );

        if (!ok) {

            return;

        }

        showApplication();

        if (
            typeof loadBooks ===
            "function"
        ) {

            await loadBooks();

        }

    },

);

logoutButton.addEventListener(

    "click",

    async function () {

        await logout();

    },

);