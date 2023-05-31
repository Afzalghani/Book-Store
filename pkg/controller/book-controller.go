package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/afzal/bookstore/pkg/model"
	"github.com/afzal/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook model.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := model.GetAllBooks()

	resp, _ := json.Marshal(newBooks)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

func GetBooksById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := model.GetbookById(ID)

	resp, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &model.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()

	resp, _ := json.Marshal(b)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error whiel parshing")
	}
	book := model.DeleteBook(ID)

	resp, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "appliation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var UpdateBook = &model.Book{}

	utils.ParseBody(r, UpdateBook)

	params := mux.Vars(r)

	//Get the book Id from the params
	bookId := params["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error on parsing ")
	}

	book, db := model.GetbookById(ID)

	if UpdateBook.Name != "" {
		book.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		book.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		book.Publication = UpdateBook.Publication
	}

	db.Save(&book)

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)

}
