package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func getStrings(fileName string) []string {
	var lines []string

	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return lines
}

func viewHandlerGuestbook(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("guestbook.txt")

	fmt.Printf("%#v\n", signatures)

	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}

	html, err := template.ParseFiles("guestbookView.html")
	if err != nil {
		log.Fatal(err)
	}
	err = html.Execute(writer, guestbook)
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandlerNewSignature(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("guestbookNewSignatureView.html")
	if err != nil {
		log.Fatal(err)
	}
	err = html.Execute(writer, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func postHandlerCreateSignature(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")

	file, err := os.OpenFile("guestbook.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(0600))
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintln(file, signature)
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func main() {
	http.HandleFunc("/guestbook", viewHandlerGuestbook)
	http.HandleFunc("/guestbook/new", viewHandlerNewSignature)
	http.HandleFunc("/guestbook/create", postHandlerCreateSignature)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
