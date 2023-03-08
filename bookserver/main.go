package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books []Book

func main() {
	// add some sample books to the library
	books = []Book{
		{ID: 1, Title: "The Catcher in the Rye", Author: "J.D. Salinger", Quantity: 10},
		{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee", Quantity: 5},
		{ID: 3, Title: "1984", Author: "George Orwell", Quantity: 3},
	}

	// register HTTP handlers
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/books", handleBooks)
	http.HandleFunc("/books/add", handleAddBook)
	http.HandleFunc("/books/return", handleIncBook)
	http.HandleFunc("/books/borrow", handleBorrowBook)
	http.HandleFunc("/books/delete", handleDeleteBook)

	// start the server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad request")
		return
	}
	name := query.Get("name")
	if len(name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "You must supply a name")
		return
	}
	//json.NewEncoder(w).Encode("Hello " + name)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello %s", name)
}

func handleBooks(w http.ResponseWriter, r *http.Request) {
	// return a list of all books in the library
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
	//fmt.Fprintf(w, "Hello %v", books)
}

func handleAddBook(w http.ResponseWriter, r *http.Request) {
	// add a new book to the library
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	books = append(books, newBook)
	json.NewEncoder(w).Encode(books)
}

func handleDeleteBook(w http.ResponseWriter, r *http.Request) {
	var deleteRequest struct {
		ID int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&deleteRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//fmt.Printf("id : %s", r.FormValue("id"))
	// we then need to loop through all our articles
	for i := range books {
		if books[i].ID == deleteRequest.ID {
			books = append(books[:i], books[i+1:]...)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func handleBorrowBook(w http.ResponseWriter, r *http.Request) {
	// borrow a book from the library
	var borrowRequest struct {
		ID int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&borrowRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i := range books {
		if books[i].ID == borrowRequest.ID {
			if books[i].Quantity > 0 {
				books[i].Quantity--
				json.NewEncoder(w).Encode(books[i])
			} else {
				http.Error(w, "Book not available", http.StatusNotFound)
			}
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func handleIncBook(w http.ResponseWriter, r *http.Request) {
	// borrow a book from the library
	var borrowRequest struct {
		ID int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&borrowRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i := range books {
		if books[i].ID == borrowRequest.ID {
			if books[i].Quantity > 0 {
				books[i].Quantity++
				json.NewEncoder(w).Encode(books[i])
			} else {
				http.Error(w, "Book not available", http.StatusNotFound)
			}
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}
