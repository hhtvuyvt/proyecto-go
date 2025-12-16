# 📚 Proyecto Go – Plataforma de Libros y Lectura

Este proyecto es una **aplicación web desarrollada en Go** que permite gestionar un catálogo de libros, subir portadas, autenticarse mediante JWT y consumir una API REST desde un frontend web simple.

El objetivo del proyecto es evolucionar desde un CRUD educativo hacia una **plataforma profesional de publicación y lectura**, similar a **Wattpad**, pero construida sobre la base sólida que ya existe.

---

## 🚀 ¿Qué hace el proyecto?

Actualmente la aplicación permite:

- 📖 Crear, listar y eliminar libros
- 🔍 Buscar libros por título
- 📄 Paginación de resultados
- 🖼️ Subir imágenes como portadas de libros
- 🔐 Autenticación con JWT
- 🧼 Sanitización de entradas para seguridad
- 🧪 Tests unitarios para handlers y repositorios
- 🗄️ Persistencia de datos en SQLite
- 🌐 Servir frontend estático desde el backend en Go

---

## 🧠 ¿Cómo funciona internamente?

La aplicación sigue una arquitectura clara y modular, separando responsabilidades para facilitar el mantenimiento y la escalabilidad.

### Flujo general

Frontend → Handlers HTTP → Repositorios → Base de datos (SQLite)

---

## 🗂️ Estructura del proyecto

proyecto-go/
├── data/
├── handlers/
├── internal/db/
├── middlewares/
├── models/
├── routes/
├── static/
├── uploads/
├── utils/
├── main.go
└── README.md

---

## 🛠️ Tecnologías utilizadas

- Go
- net/http
- SQLite
- HTML, JavaScript, Bootstrap
- JWT
- go test

---

## ▶️ Cómo ejecutar el proyecto

1. Crear archivo .env

PORT=8080
DB_PATH=./data/books.db
JWT_SECRET=una_clave_secreta_segura

2. Ejecutar

go run main.go

---

## 📈 Roadmap

- Edición de libros
- Capítulos por libro
- Lectura online tipo Wattpad
- Usuarios y roles
- Comentarios y favoritos
- Docker y CI/CD

---

Proyecto educativo y personal 🚀
