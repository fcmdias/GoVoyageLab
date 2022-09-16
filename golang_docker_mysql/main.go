package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers() []*User {
	// Open up our database connection.
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

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

func addUsers()  {
	// Open up our database connection.
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()






	/// start 


	sql := "INSERT INTO users(name) VALUES ('John Doe')"
	res, err := db.Exec(sql)

	if err != nil {
			panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
			log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)


	/// end 
	// Execute the query
	// result, err := db.Query("SELECT * FROM users")
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }

	// var users []*User
	// for results.Next() {
	// 	var u User
	// 	// for each row, scan the result into our tag composite object
	// 	err = results.Scan(&u.ID, &u.Name)
	// 	if err != nil {
	// 		panic(err.Error()) // proper error handling instead of panic in your app
	// 	}

	// 	users = append(users, &u)
	// }

	// return users
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func userPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()
	
	fmt.Println("Endpoint Hit: usersPage")
	json.NewEncoder(w).Encode(users)
}
func adduserPage(w http.ResponseWriter, r *http.Request) {
	addUsers()
	
	fmt.Println("Endpoint Hit: adduserPage")
	fmt.Fprintf(w, "users added!")
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", userPage)
	http.HandleFunc("/users/add", adduserPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}