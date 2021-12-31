package main

import (
	"log"
	"net/http"
)

func sayXWorld(x string) []byte {
	phrase := x + " World!"
	return []byte(phrase)
}

func getHelloWorld(helloWorld func(string) []byte, word string) []byte {
	return helloWorld(word)
}

func viewHandlerHello(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write(getHelloWorld(sayXWorld, "Hello"))
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandlerSalut(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write(getHelloWorld(sayXWorld, "Salut"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandlerHello)
	http.HandleFunc("/salut", viewHandlerSalut)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
