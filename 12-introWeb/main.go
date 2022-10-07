package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//handlers
func hi(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El metodo de esta peticion es: " + r.Method)
	fmt.Fprintln(rw, "Hola")
}

func dontExist(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func errorr(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Error en la pericion 404 ", http.StatusNotFound)
}

func dataa(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	i, _ := strconv.Atoi(age)
	fmt.Fprintln(rw, " Hola ", name)

	if i >= 18 {
		fmt.Fprint(rw, "Felicidades sos grande!")
	} else {
		fmt.Fprintln(rw, "Sos un puberto")
	}

}

func main() {
	//mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", hi)
	mux.HandleFunc("/dont", dontExist)
	mux.HandleFunc("/error", errorr)
	mux.HandleFunc("/hi", dataa)

	//Router
	/*
		http.HandleFunc("/", hi)
		http.HandleFunc("/dont", dontExist)
		http.HandleFunc("/error", errorr)
		http.HandleFunc("/hi", dataa)*/

	fmt.Println("Se esta ejecutando el servidor en el puerto 9000")
	fmt.Println("Servidor corriendo  http://localhost:9000/")
	//http.ListenAndServe("localhost:9000", mux)

	//Servidor
	server := &http.Server{
		Addr:    "localhost:9000",
		Handler: mux,
	}

	server.ListenAndServe()
}
