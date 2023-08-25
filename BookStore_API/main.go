package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type BookDetails struct {
	BookID   int    `json:"bookId"`
	BookName string `json:"bookName"`
	Author   string `json:"author"`
	Price    int    `json:"price"`
}

var books = []BookDetails{}

func addBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newBook BookDetails
		err := json.NewDecoder(r.Body).Decode(&newBook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		books = append(books, newBook)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Book Added Successfully")
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(books)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		var updatedBook BookDetails
		err := json.NewDecoder(r.Body).Decode(&updatedBook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for i, book := range books {
			if book.BookID == updatedBook.BookID {
				books[i] = updatedBook
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, "Book details updated successfully")
				return
			}

		}
		http.Error(w, "Book not found", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		bookIDStr := r.URL.Query().Get("bookId")
		bookId, err := strconv.Atoi(bookIDStr)
		if err != nil {
			http.Error(w, "Invalid Book Id Parameter", http.StatusBadRequest)
			return
		}
		for i, book := range books {
			if book.BookID == bookId {
				books = append(books[:i], books[i+1:]...)
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, "Book Deleted Successfully")
				return
			}
		}
		http.Error(w, "Book not found", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/addBook", addBook)
	http.HandleFunc("/getbooks", getBook)
	http.HandleFunc("/updateBook", updateBook)
	http.HandleFunc("/deleteBook", deleteBook)
	log.Println("Listening...!")
	log.Fatal(http.ListenAndServe(":8181", nil))
}
