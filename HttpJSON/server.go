package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func decoder(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Fprintf(w, "Tu usuario: %s %s tiene %d año/s", user.Firstname, user.Lastname, user.Age)
}

func encoder(w http.ResponseWriter, r *http.Request) {
	juan := User{
		Firstname: "Juanito",
		Lastname:  "Sanchez",
		Age:       99,
	}
	json.NewEncoder(w).Encode(juan)
}

func main() {

	http.HandleFunc("/decode", decoder)
	http.HandleFunc("/encode", encoder)

	err := http.ListenAndServe(":3333", nil) // Lee peticiones de cualquier IP al puerto 3333

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
