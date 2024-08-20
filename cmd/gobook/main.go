package main

import (
	"database/sql"
	"net/http"

	"github.com/enrick-dev/gobooks.git/intermal/service"
	"github.com/enrick-dev/gobooks.git/intermal/web"
)

func main() {
	db, err := sql.Open("sqlite34", "./books.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookService := service.NewBookService(db)

	bookHandlers := web.NewBookHandlers((bookService))

	router := http.NewServeMux()
	router.HandleFunc("GET /books", bookHandlers.GetBooks)
	router.HandleFunc("POST /books", bookHandlers.CreateBook)
	router.HandleFunc("GET /books/{ID}", bookHandlers.GetBookByID)
	router.HandleFunc("PUT /books/{ID}", bookHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{ID}", bookHandlers.DeleteBook)

	http.ListenAndServe(":8080", router)
}
