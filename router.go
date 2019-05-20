package awsom_pcakage

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Id      int    `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Id: 1, Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: 2, Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func StartServer() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}

func My() {
	fmt.Print("from myrouter")
}
