const API = "/api/books";

let token = null;
let editingBookId = null;
let books = [];


const titleInput = document.getElementById("title");
const authorInput = document.getElementById("author");
const isbnInput = document.getElementById("isbn");
const imageInput = document.getElementById("image");
const form = document.getElementById("book-form");
const saveButton = document.getElementById("saveButton");
const bookList = document.getElementById("book-list");



async function loadToken() {

    try {

        const res = await fetch("/api/login");


        if (!res.ok) {

            console.error(
                "Login no disponible:",
                res.status
            );

            token = null;

            return;
        }



        const data = await res.json();


        if (!data.token) {

            console.error(
                "Respuesta de login sin token"
            );

            token = null;

            return;
        }



        token = data.token;



    } catch (e) {

        console.error(
            "No se pudo obtener token",
            e
        );

        token = null;

    }

}


function authHeaders() {

    const headers = {
        "Content-Type": "application/json"
    };


    if (token) {
        headers.Authorization =
            `Bearer ${token}`;
    }


    return headers;
}




async function loadBooks() {

    const res = await fetch(API);


    if (!res.ok) {
        console.error("Error cargando libros");
        books = [];
        renderBooks();
        return;
    }


    const data = await res.json();


    books = Array.isArray(data)
        ? data
        : [];


    renderBooks();

}




function renderBooks() {


    bookList.innerHTML = "";


    books.forEach(book => {


        const div =
            document.createElement("div");


        div.className =
            "book col-sm-6 col-md-4 col-lg-3";


        div.dataset.id =
            String(book.id);



        div.innerHTML = `

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
data-action="edit"
class="btn btn-primary btn-sm">

Editar

</button>


<button
data-action="delete"
class="btn btn-danger btn-sm">

Borrar

</button>


</div>

</div>

`;


        bookList.appendChild(div);


    });


}





form.addEventListener(
    "submit",
    async e => {


        e.preventDefault();



        const data = {


            title:
                titleInput.value.trim(),


            author:
                authorInput.value.trim(),


            isbn:
                isbnInput.value.trim(),


            image:
                imageInput.value.trim()

        };



        let response;



        if (editingBookId !== null) {


            const oldBook =
                books.find(
                    b => Number(b.id) === Number(editingBookId)
                );



            if (oldBook) {

                data.author =
                    data.author || oldBook.author;

                data.isbn =
                    data.isbn || oldBook.isbn;

                data.image =
                    data.image || oldBook.image;

            }



            response =
                await fetch(
                    `${API}/${editingBookId}`,
                    {

                        method:"PUT",

                        headers:
                            authHeaders(),

                        body:
                            JSON.stringify(data)

                    }
                );


        } else {


            response =
                await fetch(
                    API,
                    {

                        method:"POST",

                        headers:{
                            "Content-Type":
                                "application/json"
                        },

                        body:
                            JSON.stringify(data)

                    }
                );

        }



        if (!response.ok) {

            console.error(
                "Error guardando libro"
            );

            return;

        }



        editingBookId = null;


        saveButton.innerText =
            "Agregar";


        form.reset();


        await loadBooks();


    });






bookList.addEventListener(
    "click",
    async e => {


        const card =
            e.target.closest(".book");


        if (!card)
            return;



        const id =
            Number(card.dataset.id);



        if (
            e.target.dataset.action === "delete"
        ) {


            const res =
                await fetch(
                    `${API}/${id}`,
                    {

                        method:"DELETE",

                        headers:
                            authHeaders()

                    }
                );


            if(res.ok){

                await loadBooks();

            }


            return;

        }






        if (
            e.target.dataset.action === "edit"
        ) {


            const book =
                books.find(
                    b => Number(b.id) === id
                );


            if(!book)
                return;



            editingBookId =
                id;



            titleInput.value =
                book.title || "";


            authorInput.value =
                book.author || "";


            isbnInput.value =
                book.isbn || "";


            imageInput.value =
                book.image || "";



            saveButton.innerText =
                "Guardar cambios";


            titleInput.focus();

        }


    });







window.onload =
    async () => {


        await loadToken();


        await loadBooks();


    };