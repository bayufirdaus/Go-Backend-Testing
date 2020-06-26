package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// Book struct (Model)
type Publisher struct {
	ID     string `json:"publisher_id"`
	Name   string `json:"publisher_name"`
	Date   string `json:"input_date"`
	Update string `json:"last_update"`
}

// Get all orders

func getPublishers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var publishers []Publisher

	sql := `SELECT
				publisher_id,
				IFNULL(publisher_name,''),
				IFNULL(input_date,'') input_date, 
				IFNULL(last_update,'') last_update
			FROM publisher`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var publisher Publisher
		err := result.Scan(&publisher.ID, &publisher.Name,
			&publisher.Date, &publisher.Update)

		if err != nil {
			panic(err.Error())
		}
		publishers = append(publishers, publisher)
	}

	json.NewEncoder(w).Encode(publishers)
}

func createPublisher(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		ID := r.FormValue("publisher_id")
		Name := r.FormValue("publisher_name")
		Date := r.FormValue("input_date")
		Update := r.FormValue("last_update")

		stmt, err := db.Prepare("INSERT INTO publisher (publisher_id,publisher_name,input_date,last_update) VALUES(?,?,?,?)")

		if err != nil {
			panic(err.Error())
		}

		_, err = stmt.Exec(ID, Name, Date, Update)

		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintf(w, "Publisher Created")
		//http.Redirect(w, r, "/", 301)
	}
}

func getPublisher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var publishers []Publisher
	params := mux.Vars(r)

	sql := `SELECT
				publisher_id,
				IFNULL(publisher_name,''),
				IFNULL(input_date,'') input_date, 
				IFNULL(last_update,'') last_update
			FROM publisher WHERE publisher_id = ?`

	result, err := db.Query(sql, params["id"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var publisher Publisher

	for result.Next() {

		err := result.Scan(&publisher.ID, &publisher.Name,
			&publisher.Date, &publisher.Update)

		if err != nil {
			panic(err.Error())
		}

		publishers = append(publishers, publisher)
	}

	json.NewEncoder(w).Encode(publishers)
}

func updatePublisher(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		params := mux.Vars(r)
		Name := r.FormValue("publisher_name")
		Date := r.FormValue("input_date")
		Update := r.FormValue("last_update")

		stmt, err := db.Prepare("UPDATE publisher SET publisher_name = ?, input_date = ?,  last_update = ? WHERE publisher_id = ?")

		_, err = stmt.Exec(Name, Date, Update, params["id"])

		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintf(w, "Publisher with ID = %s was updated", params["id"])
	}
}

func deletePublisher(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM publisher WHERE publisher_id = ?")

	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Publisher with ID = %s was deleted", params["id"])
}

func getPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var publishers []Publisher

	ID := r.FormValue("publisher_id")

	sql := `SELECT
				publisher_id,
				IFNULL(publisher_name,''),
				IFNULL(input_date,'') input_date, 
				IFNULL(last_update,'') last_update
			FROM publisher WHERE publisher_id = ?`

	result, err := db.Query(sql, ID)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var publisher Publisher

	for result.Next() {
		err := result.Scan(&publisher.ID, &publisher.Name,
			&publisher.Date, &publisher.Update)

		if err != nil {
			panic(err.Error())
		}

		publishers = append(publishers, publisher)
	}

	json.NewEncoder(w).Encode(publishers)

}

func getPut(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var publishers []Publisher

	params := mux.Vars(r)

	sql := `SELECT
				publisher_id,
				IFNULL(publisher_name,''),
				IFNULL(input_date,'') input_date, 
				IFNULL(last_update,'') last_update
			FROM publisher WHERE publisher_id = ?`

	result, err := db.Query(sql, params["id"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var publisher Publisher

	for result.Next() {
		err := result.Scan(&publisher.ID, &publisher.Name,
			&publisher.Date, &publisher.Update)

		if err != nil {
			panic(err.Error())
		}

		publishers = append(publishers, publisher)
	}

	json.NewEncoder(w).Encode(publishers)

}

// Main function
func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_testing")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Server listened with port 8181")
	defer db.Close()

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/publisher", getPublishers).Methods("GET")
	r.HandleFunc("/publisher/{id}", getPublisher).Methods("GET")
	r.HandleFunc("/publisher", createPublisher).Methods("POST")
	r.HandleFunc("/publisher/{id}", updatePublisher).Methods("PUT")
	r.HandleFunc("/publisher/{id}", deletePublisher).Methods("DELETE")
	r.HandleFunc("/publisher/", getPost).Methods("POST", "GET")
	r.HandleFunc("/publisher/{id}", getPut).Methods("PUT", "GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8181", r))
}
