package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Posts struct {
	Messages []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getPosts(filename string) []string {
	var posts []string
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		posts = append(posts, scanner.Text())
	}
	check(scanner.Err())
	return posts
}

func handleHome(writer http.ResponseWriter, request *http.Request) {
	messages := getPosts("assets/posts.txt")
	html, err := template.ParseFiles("assets/home.html")
	check(err)
	posts := Posts{
		Messages: messages,
	}
	err = html.Execute(writer, posts)
	check(err)
}

func main() {
	http.HandleFunc("/", handleHome)

	port := os.Getenv("PORT")
	httpPort := ":" + port
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
