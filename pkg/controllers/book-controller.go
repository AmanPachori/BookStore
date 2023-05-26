package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/AmanPachori/bookstore/pkg/utils"
	"github.com/AmanPachori/bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter,r *http.Request){
	newbooks := models.GetBooks()
	res,_ :=json.Marshal(newbooks)
	w.Header().Set("Content-Type", "pkglication/json")
  w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter,r *http.Request){
	vars:= mux.Vars(r)
	bookId:= vars["bookId"]
	Id,err:= strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing book")
	}
	bookDetails,_:= models.GetBookById(Id)
	res,_ :=json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
  w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter,r *http.Request){

	CreateBook := &models.Book{}
	utils.ParseBody(r,CreateBook)
	b:= CreateBook.CreateBook()
	res,_:=json.Marshal(b)
  w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook (w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	bookId:= vars["bookId"]
	Id,err := strconv.ParseInt(bookId,0,0)
	if(err != nil){
		fmt.Println("error parsing book")
	}
	book:= models.DeleteBook(Id)
	res,_:=json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var UpdateBook = models.Book{}
	utils.ParseBody(r,UpdateBook)
	vars:=mux.Vars(r)
	bookId := vars["bookId"]
	Id,err := strconv.ParseInt(bookId,0,0)
	if(err != nil){
		fmt.Println("error parsing book")
	}

	bookDetails, db:=models.GetBookById(Id)
	if UpdateBook.Name !=""{
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author !=""{
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication !=""{
		bookDetails.Publication = UpdateBook.Publication
	}
	db.Save(&bookDetails)
	res,_:= json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
  w.WriteHeader(http.StatusOK)
	w.Write(res)

}
