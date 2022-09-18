package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"os"
)

var version = "0.1.5"
var db *sql.DB

//Connect creates MySQL connection
func Connect() error {
	var err error
		user := os.Getenv("mysql_user")
		secret := os.Getenv("mysql_secret")
		port := os.Getenv("mysql_port")
		dbname := os.Getenv("mysql_db")
		db, err = sql.Open("mysql", user +":" + secret + "@tcp(db:" + port + ")/" + dbname)	
		if err != nil {
			return err
		}
		return nil
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers() []*User {

	// Execute the query
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var users []*User
	for results.Next() {
		var u User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, &u)
	}

	return users
}


func getUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
  result, err := db.Query("SELECT id, name FROM users WHERE id = ?", params["id"])
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  var user User
  for result.Next() {
    err := result.Scan(&user.ID, &user.Name)
    if err != nil {
      panic(err.Error())
    }
  }
	json.NewEncoder(w).Encode(user)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
  if err != nil {
    panic(err.Error())
  }
  _, err = stmt.Exec(params["id"])
 if err != nil {
    panic(err.Error())
  }
	fmt.Fprintf(w, "User with ID = %s was deleted", params["id"])
}

func updateUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  stmt, err := db.Prepare("UPDATE users SET name = ? WHERE id = ?")
  if err != nil {
    panic(err.Error())
  }
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err.Error())
  }
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)
  newName := keyVal["name"]
  _, err = stmt.Exec(newName, params["id"])
  if err != nil {
    panic(err.Error())
  }
  fmt.Fprintf(w, "User with ID = %s was updated with name: %s", params["id"], newName)
}

func createUser(w http.ResponseWriter, r *http.Request) {
  stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
  if err != nil {
    panic(err.Error())
  }
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err.Error())
  }
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)
  name := keyVal["name"]
  _, err = stmt.Exec(name)
  if err != nil {
    panic(err.Error())
  }
  fmt.Fprintf(w, "New user was created")
}

func homePage(w http.ResponseWriter, r *http.Request) {

	env := os.Getenv("env")
	fmt.Fprintf(w, "Welcome to the HomePage! version: %s, env: %s", version, env)
	fmt.Println("Endpoint Hit: homePage")
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Error connecting to db %s", globalError.Error())
	fmt.Println("Error connecting to db %s", globalError.Error())
}

func usersPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()
	
	fmt.Println("Endpoint Hit: usersPage")
	json.NewEncoder(w).Encode(users)
}

var globalError error

func main() {
	globalError = Connect()

	router := mux.NewRouter()
	if globalError != nil {
		http.HandleFunc("/", errorPage)
	} else {
		router.HandleFunc("/", homePage).Methods("GET")

		router.HandleFunc("/user", createUser).Methods("POST")
		router.HandleFunc("/user/{id}", getUser).Methods("GET")
		router.HandleFunc("/user/{id}", updateUser).Methods("PUT")
		router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
		router.HandleFunc("/users", usersPage).Methods("GET")
		

		defer db.Close()
	}
	http.ListenAndServe(":8080", router)
}