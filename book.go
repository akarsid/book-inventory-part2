package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
var err error
const DNS="root:siddharth@tcp(127.0.0.1:3306)/booksdb?charset=utf8&parseTime=True&loc=Local"
type Book struct{
	gorm.Model
	// ID int `gorm:"AUTO_INCREMENT"`
	Title string `json:"title"`
	Author string `json:"author"`
}
func InitialMigration(){
	DB,err=gorm.Open(mysql.Open(DNS),&gorm.Config{})
	if err!=nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Book{})
}
func GetBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var books []Book
	DB.Find(&books)
	json.NewEncoder(w).Encode(books)
}
func GetBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	bookId:=params["bookId"]
	i, _:= strconv.Atoi(bookId)
	
	var book []Book
	DB.First(&book,i)
	json.NewEncoder(w).Encode(book)
}
func CreateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	//convert to the type of var provided in Decode
	json.NewDecoder(r.Body).Decode(&book)
	DB.Create(&book)
	json.NewEncoder(w).Encode(book)

}
func getBooksCount(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var books []Book
	DB.Find(&books)
	json.NewEncoder(w).Encode(len(books))
}
func getBooksByAuthor(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	author:=params["authName"]
	var books []Book
	DB.Where("Author <> ?", author).Find(&books)
	// var booksByAuthor []Book
	// for _, book := range books{
	// 	if book.Author == author{
	// 		booksByAuthor = append(booksByAuthor, book)
	// 	}
	// }
	json.NewEncoder(w).Encode(books)
}
func GetAuthors(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var books []Book
	DB.Find(&books)
	var authors []string 
	for _, book := range books{
		authors = append(authors, book.Author)
	}
	json.NewEncoder(w).Encode(authors)
}
// func UpdateUser(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	params:=mux.Vars(r)
// 	var user []User
// 	DB.First(&user,params["id"])
// 	json.NewDecoder(r.Body).Decode(&user)
// 	DB.Save(&user)
// 	json.NewEncoder(w).Encode(user)
// }
// func DeleteUser(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	params:=mux.Vars(r)
// 	var user []User
// 	DB.Delete(&user,params["id"])
// 	json.NewEncoder(w).Encode("user is deleted successfully")
// }




