//
// auth.js
//
// Gestiona toda la autenticación de la aplicación.
//

"use strict";

// =============================
// Elementos del DOM
// =============================

let loginPanel;
let appPanel;
let userPanel;
let loginForm;
let usernameInput;
let passwordInput;
let loginError;
let usernameLabel;
let logoutButton;

// Inicializar referencias al DOM cuando esté listo
function initDOMElements() {
  loginPanel = document.getElementById("loginPanel");
  appPanel = document.getElementById("appPanel");
  userPanel = document.getElementById("userPanel");
  loginForm = document.getElementById("loginForm");
  usernameInput = document.getElementById("loginUsername");
  passwordInput = document.getElementById("loginPassword");
  loginError = document.getElementById("loginError");
  usernameLabel = document.getElementById("usernameLabel");
  logoutButton = document.getElementById("logoutButton");
}

// =============================
// Interfaz
// =============================

function showLogin() {
  if (loginPanel) loginPanel.classList.remove("d-none");
  if (appPanel) appPanel.classList.add("d-none");
  if (userPanel) userPanel.classList.add("d-none");

  clearLoginError();

  if (passwordInput) passwordInput.value = "";
  if (usernameInput) usernameInput.focus();
}

function showApplication() {
  if (loginPanel) loginPanel.classList.add("d-none");
  if (appPanel) appPanel.classList.remove("d-none");
  if (userPanel) userPanel.classList.remove("d-none");
}

function showLoginError(message) {
  if (!loginError) return;
  loginError.textContent = message;

  loginError.classList.remove("d-none");
}

function clearLoginError() {
  if (!loginError) return;
  loginError.textContent = "";

  loginError.classList.add("d-none");
}

// =============================
// API
// =============================

async function login(username, password) {
  clearLoginError();

  const response = await fetch("/api/login", {
    method: "POST",

    credentials: "same-origin",

    headers: {
      "Content-Type": "application/json",
    },

    body: JSON.stringify({
      username,

      password,
    }),
  });

  if (!response.ok) {
    showLoginError("Usuario o contraseña incorrectos.");

    return false;
  }

  if (usernameLabel) {
    usernameLabel.textContent = username;
  }

  return true;
}

async function logout() {
  await fetch("/api/logout", {
    method: "POST",

    credentials: "same-origin",
  });

  if (typeof showLogin === "function" && loginPanel) {
    showLogin();
  } else {
    window.location.href = "/static/index.html";
  }
}

async function checkSession() {
  try {
    const response = await fetch("/api/me", {
      credentials: "same-origin",
    });

    if (!response.ok) {
      return false;
    }

    const user = await response.json();

    if (usernameLabel) {
      usernameLabel.textContent = user.username;
    }

    if (userPanel) {
      userPanel.classList.remove("d-none");
    }

    return true;
  } catch {
    return false;
  }
}

// =============================
// Eventos (Protegidos por condición)
// =============================

function attachEventListeners() {
  if (loginForm) {
    loginForm.addEventListener(
      "submit",

      async function (event) {
        event.preventDefault();

        const username = usernameInput.value.trim();

        const password = passwordInput.value;

        const ok = await login(username, password);

        if (!ok) {
          return;
        }

        showApplication();

        if (typeof loadBooks === "function") {
          await loadBooks();
        }
      },
    );
  }

  if (logoutButton) {
    logoutButton.addEventListener(
      "click",

      async function () {
        await logout();
      },
    );
  }
}

// Comprobar la sesión automáticamente al cargar cualquier página
document.addEventListener("DOMContentLoaded", async () => {
  initDOMElements();
  attachEventListeners();

  const isLogged = await checkSession();
  if (isLogged && typeof showApplication === "function" && loginPanel) {
    showApplication();
    if (typeof loadBooks === "function") {
      await loadBooks();
    }
  } else if (!isLogged && loginPanel) {
    showLogin();
  }
});
