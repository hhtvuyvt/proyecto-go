//
// books.js
//
// CRUD del catálogo de libros.
//

"use strict";

// ===================================
// Estado
// ===================================

const API =
    "/api/books";

let books = [];

let editingBookId =
    null;

// ===================================
// Referencias del DOM
// ===================================

const form =
    document.getElementById(
        "book-form",
    );

const titleInput =
    document.getElementById(
        "title",
    );

const authorInput =
    document.getElementById(
        "author",
    );

const isbnInput =
    document.getElementById(
        "ISBN",
    );

const imageInput =
    document.getElementById(
        "image",
    );

const saveButton =
    document.getElementById(
        "saveButton",
    );

const bookList =
    document.getElementById(
        "book-list",
    );

// ===================================
// Utilidades
// ===================================

function resetForm() {

    editingBookId = null;

    form.reset();

    saveButton.innerText =
        "Agregar";

}

function sessionExpired() {

    alert(
        "La sesión ha expirado.",
    );

    if (
        typeof showLogin ===
        "function"
    ) {

        showLogin();

    }

}

// ===================================
// API
// ===================================

async function loadBooks() {

    try {

        const response =
            await fetch(
                API,
                {

                    credentials:
                        "same-origin",

                },
            );

        if (
            response.status === 401
        ) {

            sessionExpired();

            return;

        }

        books =
            await response.json();

        renderBooks();

    } catch (error) {

        console.error(
            error,
        );

    }

}

// ===================================
// Render
// ===================================

function renderBooks() {

    bookList.innerHTML =
        "";

    books.forEach(

        function (
            book,
        ) {

            const card =
                document.createElement(
                    "div",
                );

            card.className =
                "book col-sm-6 col-md-4 col-lg-3";

            card.dataset.id =
                book.id;

            card.innerHTML = `

<div class="card h-100 shadow-sm">

<img
class="card-img-top"
src="${book.image || ""}"
alt="${book.title}">

<div class="card-body">

<h5 class="card-title">

${book.title}

</h5>

<p class="card-text">

${book.author}

</p>

<p class="card-text">

${book.isbn}

</p>

<button
class="btn btn-primary btn-sm"
data-action="edit">

Editar

</button>

<button
class="btn btn-danger btn-sm"
data-action="delete">

Borrar

</button>

</div>

</div>

`;

            bookList.appendChild(
                card,
            );

        },

    );

}

// ===================================
// Eventos del catálogo
// ===================================

// Gestiona los botones Editar y Borrar
// mediante delegación de eventos.
bookList.addEventListener(

    "click",

    async function (
        event,
    ) {

        const button =
            event.target.closest(
                "button",
            );

        if (!button) {

            return;

        }

        const card =
            button.closest(
                ".book",
            );

        if (!card) {

            return;

        }

        const id =
            Number(
                card.dataset.id,
            );

        // ===========================
        // Editar
        // ===========================

        if (
            button.dataset.action ===
            "edit"
        ) {

            const book =
                books.find(

                    function (b) {

                        return b.id === id;

                    },

                );

            if (!book) {

                return;

            }

            editingBookId =
                id;

            titleInput.value =
                book.title;

            authorInput.value =
                book.author;

            isbnInput.value =
                book.isbn || "";

            imageInput.value =
                book.image || "";

            saveButton.innerText =
                "Guardar cambios";

            titleInput.focus();

        }

    },

);

// ===================================
// Formulario
// ===================================

form.addEventListener(

    "submit",

    async function (
        event,
    ) {

        event.preventDefault();

        const data = {

            title:
                titleInput.value.trim(),

            author:
                authorInput.value.trim(),

            isbn:
                isbnInput.value.trim(),

            image:
                imageInput.value.trim(),

        };

        let response;

        if (
            editingBookId ===
            null
        ) {

            response =
                await fetch(

                    API,

                    {

                        method:
                            "POST",

                        credentials:
                            "same-origin",

                        headers: {

                            "Content-Type":
                                "application/json",

                        },

                        body:
                            JSON.stringify(
                                data,
                            ),

                    },

                );

        } else {

            response =
                await fetch(

                    `${API}/${editingBookId}`,

                    {

                        method:
                            "PUT",

                        credentials:
                            "same-origin",

                        headers: {

                            "Content-Type":
                                "application/json",

                        },

                        body:
                            JSON.stringify(
                                data,
                            ),

                    },

                );

        }

        if (
            response.status ===
            401
        ) {

            sessionExpired();

            return;

        }

        if (
            !response.ok
        ) {

            console.error(
                "No fue posible guardar el libro.",
            );

            return;

        }

        resetForm();

        await loadBooks();

    },

);