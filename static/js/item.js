//
// item.js
//
// CRUD del catálogo de Items (Productos, Habitaciones, etc.) para el SaaS.
//

"use strict";

// ===================================
// Estado
// ===================================

const API = "/api/items";

let items = [];
let editingItemId = null;

// ===================================
// Referencias del DOM
// ===================================

const form = document.getElementById("item-form");
const nameInput = document.getElementById("name"); // Antes: title
const categoryInput = document.getElementById("category"); // Antes: author
const codeInput = document.getElementById("code"); // Antes: ISBN
const imageInput = document.getElementById("image");
const saveButton = document.getElementById("saveButton");
const itemList = document.getElementById("item-list");

// Referencias para el primer detalle opcional (antes: capítulo inicial)
const initialDetailNameInput = document.getElementById("initialDetailName");
const initialDetailContentInput = document.getElementById(
  "initialDetailContent",
);

// ===================================
// Utilidades
// ===================================

function resetForm() {
  editingItemId = null;
  form.reset();
  saveButton.innerText = "Agregar";
}

function sessionExpired() {
  alert("La sesión ha expirado.");
  if (typeof showLogin === "function") {
    showLogin();
  }
}

// ===================================
// API
// ===================================

async function loadItems() {
  try {
    const response = await fetch(API, {
      credentials: "same-origin",
    });

    if (response.status === 401) {
      sessionExpired();
      return;
    }

    items = await response.json();
    renderItems();
  } catch (error) {
    console.error("Error cargando items:", error);
  }
}

// ===================================
// Render
// ===================================

function renderItems() {
  itemList.innerHTML = "";

  items.forEach(function (item) {
    const card = document.createElement("div");
    card.className = "item col-sm-6 col-md-4 col-lg-3";
    card.dataset.id = item.id;

    card.innerHTML = `
      <div class="card h-100 shadow-sm">
        <img
          class="card-img-top"
          src="${item.image || "https://via.placeholder.com/300x200?text=Sin+Imagen"}"
          alt="${item.name}"
          style="object-fit: cover; height: 200px;"
        >
        <div class="card-body d-flex flex-column">
          <h5 class="card-title">${item.name}</h5>
          <p class="card-text"><strong>Categoría:</strong> ${item.category}</p>
          <p class="card-text"><strong>Código:</strong> ${item.code}</p>

          <div class="mt-auto">
            <a
              href="/static/reader.html?item_id=${item.id}"
              class="btn btn-success btn-sm w-100 mb-2"
            >
              Ver Detalles
            </a>
            <button class="btn btn-primary btn-sm" data-action="edit">Editar</button>
            <button class="btn btn-danger btn-sm" data-action="delete">Borrar</button>
          </div>
        </div>
      </div>
    `;

    itemList.appendChild(card);
  });
}

// ===================================
// Eventos del catálogo
// ===================================

// Gestiona los botones Editar y Borrar mediante delegación de eventos.
itemList.addEventListener("click", async function (event) {
  const button = event.target.closest("button");
  if (!button) return;

  const card = button.closest(".item");
  if (!card) return;

  const id = Number(card.dataset.id);

  // ===========================
  // Editar
  // ===========================
  if (button.dataset.action === "edit") {
    const item = items.find(function (i) {
      return i.id === id;
    });

    if (!item) return;

    editingItemId = id;
    nameInput.value = item.name;
    categoryInput.value = item.category;
    codeInput.value = item.code || "";
    imageInput.value = item.image || "";
    saveButton.innerText = "Guardar cambios";
    nameInput.focus();
  }

  // ===========================
  // Borrar
  // ===========================
  if (button.dataset.action === "delete") {
    const confirmed = confirm("¿Eliminar este item?");
    if (!confirmed) return;

    const response = await fetch(`${API}/${id}`, {
      method: "DELETE",
      credentials: "same-origin",
    });

    if (response.status === 401) {
      sessionExpired();
      return;
    }

    if (!response.ok) {
      console.error("No fue posible eliminar el item.");
      alert("Error al eliminar el item.");
      return;
    }

    await loadItems();
  }
});

// ===================================
// Formulario
// ===================================

form.addEventListener("submit", async function (event) {
  event.preventDefault();

  const data = {
    name: nameInput.value.trim(),
    category: categoryInput.value.trim(),
    code: codeInput.value.trim(),
    image: imageInput.value.trim(),
  };

  let response;
  let createdItemId = null;

  if (editingItemId === null) {
    // Crear nuevo
    response = await fetch(API, {
      method: "POST",
      credentials: "same-origin",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    if (response.ok) {
      const createdItem = await response.json();
      createdItemId = createdItem.id;
    }
  } else {
    // Actualizar existente
    response = await fetch(`${API}/${editingItemId}`, {
      method: "PUT",
      credentials: "same-origin",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
  }

  if (response.status === 401) {
    sessionExpired();
    return;
  }

  if (!response.ok) {
    console.error("No fue posible guardar el item.");
    alert("Error al guardar el item.");
    return;
  }

  // Si es un item nuevo y se completó el nombre del detalle inicial, lo creamos automáticamente
  if (
    editingItemId === null &&
    createdItemId &&
    initialDetailNameInput &&
    initialDetailNameInput.value.trim() !== ""
  ) {
    const detailData = {
      name: initialDetailNameInput.value.trim(),
      content: initialDetailContentInput.value.trim(),
    };

    try {
      await fetch(`${API}/${createdItemId}/details`, {
        method: "POST",
        credentials: "same-origin",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(detailData),
      });
    } catch (err) {
      console.error("Error al registrar el detalle inicial:", err);
    }
  }

  resetForm();
  loadItems();
});

// Cargar items al iniciar
document.addEventListener("DOMContentLoaded", loadItems);
