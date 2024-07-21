package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Blog struct {
	BlogID string `json:"blogId"`
	Title string `json:"title"`
	Content string `json:"content"`
	Saved bool `json:"saved"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fetchBlogs/{userId}", handleFetchBlogs).Methods("GET")
	r.HandleFunc("/createBlog/{userId}", handleCreateBlog).Methods("POST")
	r.HandleFunc("/saveBlog/{userId}/{id}", handleSaveBlog).Methods("PUT")
	r.HandleFunc("/editBlog/{userId}/{id}", handleEditBlog).Methods("PATCH")

	fmt.Println("Listening on localhost:4040")
	if err := http.ListenAndServe("localhost:4040", r); err != nil {
		log.Fatal(err)
	}
}

func handleFetchBlogs(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]
	fmt.Println("Edit: ",userId);

	file, err := os.Open("blogData.json")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer file.Close()
	
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error reading file")
		return
	}
	fmt.Println("data of file", data)
	
	var result map[string][]Blog
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("error unmarshalling data")
		return
	}
	fmt.Println("result:", result)
	fmt.Println("Here are your blogs:")
	for index, value:= range result[userId] { // if we don't want to use index we can use _ for a variable name
		fmt.Printf("Blog %d:\n", index + 1)
		fmt.Println("Blog id: ", value.BlogID)
		fmt.Println("Title: ", value.Title)
		fmt.Println("Content: ", value.Content)
		fmt.Println("Is Saved: ", value.Saved)
	}
	blogs, ok := result[userId]
	if !ok {
		http.Error(writer, "No blogs found for user", http.StatusNotFound)
		return
	}
	jsonResponse, err := json.Marshal(blogs)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}

func handleCreateBlog(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]
	fmt.Println("Edit: ",userId);
}

func handleSaveBlog(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]
	fmt.Println("Save: ",userId);
}

func handleEditBlog(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]
	fmt.Println("Edit: ",userId);
}

