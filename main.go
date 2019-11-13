package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	hash = `$2a$10$xgUMm6jqiLxldJc0ie0cMuIxDUbWM5YFU2HewhCFS8Xr6BFhGPT9W`
)

type signupRequest struct {
	Username, Email, Password string
}

type tokenRequest struct {
	GrantType, Username, Password, ClientId string
}

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/token", func (w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}
		//bodyDec := json.NewDecoder(r.Body)

		//var req tokenRequest
		//if err := bodyDec.Decode(&req); err != io.EOF && err != nil {
		//	log.Fatal(err)
		//}
		request, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w,"%s", string (request))

		//if req.GrantType == "password" {
		//	if err := bcrypt.CompareHashAndPassword([]byte (hash), []byte (req.Password)); err == nil{
		//		fmt.Fprintf(w, "logged in")
		//	}
		//}
		//
		//fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/signup", func (w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}
		bodyDec := json.NewDecoder(r.Body)

		var req signupRequest
		if err := bodyDec.Decode(&req); err != io.EOF && err != nil {
			log.Fatal(err)
		}
		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		fmt.Fprintf(w, req.Username)
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, req.Email)
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, string (hash))
	})

	err := http.ListenAndServe(":9990", nil)
	fmt.Print(err)
}