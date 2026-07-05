//
// books.js
//
// Gestiona todas las operaciones relacionadas con
// el catálogo de libros.
//
// Responsabilidades:
//
// - Obtener los libros desde la API.
// - Renderizar el catálogo.
// - Crear nuevos libros.
// - Editar libros existentes.
// - Eliminar libros.
//
// La autenticación NO pertenece a este archivo.
// Toda la gestión de sesión se encuentra en auth.js.
//

"use strict";

// =============================
// Estado de la aplicación
// =============================

// Endpoint base del recurso libros.
const API =
    "/api/books";

// Copia local del catálogo.
//
// Se mantiene sincronizada con el backend
// después de cada operación CRUD.
let books = [];

// Identificador del libro que actualmente
// está siendo editado.
//
// Cuando es null el formulario trabaja
// en modo "crear".
let editingBookId =
    null;

// =============================
// Referencias al DOM
// =============================

// Formulario principal.
const form =
    document.getElementById(
        "book-form",
    );

// Campos del formulario.
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

// Botón que cambia entre
// "Agregar"
// y
// "Guardar cambios".
const saveButton =
    document.getElementById(
        "saveButton",
    );

// Contenedor donde se renderiza
// el catálogo.
const bookList =
    document.getElementById(
        "book-list",
    );

// =============================
// API
// =============================

// Obtiene el catálogo completo.
//
// Utiliza la cookie HttpOnly creada
// durante el inicio de sesión.
// No es necesario enviar tokens
// manualmente.
async function loadBooks() {

    const response =
        await fetch(
            API,
            {

                credentials:
                    "same-origin",

            },
        );

    if (!response.ok) {

        console.error(
            "No fue posible cargar el catálogo.",
        );

        books = [];

        renderBooks();

        return;

    }

    const data =
        await response.json();

    books =
        Array.isArray(data)
            ? data
            : [];

    renderBooks();

}

// =============================
// Render
// =============================

// Construye visualmente el catálogo.
//
// Cada libro se representa mediante
// una tarjeta Bootstrap que contiene:
//
// • Imagen.
// • Título.
// • Autor.
// • ISBN.
// • Botón Editar.
// • Botón Borrar.
function renderBooks() {

    bookList.innerHTML = "";

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
                String(
                    book.id,
                );

            card.innerHTML = `

<div class="card h-100 shadow-sm">

<img
class="card-img-top"
src="${book.image || ""}"
alt="Portada de ${book.title}">

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