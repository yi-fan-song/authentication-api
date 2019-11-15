package main

import (
	"authentication-api/sqlBuilder"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"os"
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
	// Load environment variables and panic if load failed
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	fmt.Print(os.Getenv("dbusername"))
	fmt.Print(os.Getenv("dbpassword"))
	fmt.Print(os.Getenv("dbhostname"))
	fmt.Print(os.Getenv("dbport"))

	sqlBuilder.GetSqlDb(os.Getenv("dbusername"), os.Getenv("dbpassword"), os.Getenv("dbhostname"), os.Getenv("dbport"))

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/token", func (w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}

		var req tokenRequest

		reqbody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(reqbody, &req); err != nil {
			panic(err)
		}



		reqbyte, err := json.Marshal(req)
		_, _ = fmt.Fprint(w, reqbyte)
	})

	http.HandleFunc("/signup", func (w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}

		var req signupRequest

		reqbody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(reqbody, &req); err != nil {
			panic(err)
		}


		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		_, _ = fmt.Fprint(w, hash)
	})

	if port := os.Getenv("port"); port != "" {
		if err := http.ListenAndServe(port, nil); err != nil {
			panic(err)
		}
	} else {
		if err := http.ListenAndServe(":9990", nil); err != nil {
			panic(err)
		}
	}
}
