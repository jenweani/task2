package main

import (
	"database/sql"
	"encoding/json"
	"jonnedu/task2/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"os"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
	_ "github.com/lib/pq"
)

// Json response struct
type JsonResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data map[string]any `json:"data"`
}

// Request body struct
type JsonRequestBody struct {
	Name string `json:"name"`
}

func main(){
	// Load environment variables from the .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

	//sql database connection
	connection_url := os.Getenv("POSTGRES_CONNECTION_URL")
	if connection_url == ""{
		log.Fatal("connection url not set in env")
		return
	}
	db, err := sql.Open("postgres", connection_url)
	if err != nil {
        log.Fatal(err)
		return
    }
 
	// rate limiter with a rate of 10 requests per second
	limiter := rate.NewLimiter(10, 1)

	// create new persons table if it does not exist
	createPersonTable(db)

	//api endpoints
	m := mux.NewRouter()
	// POST Request endpoint
	m.HandleFunc("/api", func (w http.ResponseWriter, r *http.Request){
		if limiter.Allow() {
			w.Header().Set("Content-Type", "application/json")
			postFunc(w, r, db)
		}else {
			// user exceeds the rate limit
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Rate limit exceeded!"))
		}
	}).Methods("POST")

	// GET Request endpoint
	m.HandleFunc("/api/{user_id}", func (w http.ResponseWriter, r *http.Request){
		if limiter.Allow() {	
			user_id, e := strconv.Atoi(mux.Vars(r)["user_id"])
			if e != nil {
				// Handle the error if the string cannot be converted to an int
				http.Error(w, e.Error(), http.StatusInternalServerError)
				return
			}

			var body JsonRequestBody
			err := json.NewDecoder(r.Body).Decode(&body)
			// failing the request if both user_id and name are not parsed correctly 
			if err != nil && user_id == 0 {
				http.Error(w, "Failed to read body and no user_id param was passed", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			getFunc(w, r, db, user_id, body.Name)
		}else {
			// user exceeds the rate limit
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Rate limit exceeded!"))
		}
	}).Methods("GET")

	// PATCH Request endpoint
	m.HandleFunc("/api/{user_id}", func (w http.ResponseWriter, r *http.Request){
		if limiter.Allow() {	
			user_id, e := strconv.Atoi(mux.Vars(r)["user_id"])
			if e != nil {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				return
			}
			var body JsonRequestBody
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil || body.Name == "" {
				http.Error(w, "Failed to read body", http.StatusInternalServerError)
				return
			}
			defer r.Body.Close() // close body io reader
			w.Header().Set("Content-Type", "application/json")
			patchFunc(w, r, db, user_id, body.Name)
		}else {
			// user exceeds the rate limit
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Rate limit exceeded!"))
		}
	}).Methods("PATCH")

	// DELETE Request endpoint
	m.HandleFunc("/api/{user_id}", func (w http.ResponseWriter, r *http.Request){
		if limiter.Allow() {	
			user_id, e := strconv.Atoi(mux.Vars(r)["user_id"])
			if e != nil {
				fmt.Println("Error:", err)
				return
			}

			deleteFunc(w, r, db, user_id)
		}else {
			// user exceeds the rate limit
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Rate limit exceeded!"))
		}
	}).Methods("DELETE")

	defer db.Close()

	http.ListenAndServe(":80", m)
}

// endpoint functions 
// Post Request Fuction
func postFunc(w http.ResponseWriter, r *http.Request, db *sql.DB){
	var person models.Person
	e := json.NewDecoder(r.Body).Decode(&person);
	if   e != nil || person.Name == "" { // make sure that the correct type is passed and is not empty
		http.Error(w, "Failed to read body or name cannot be empty string", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close() // close body io reader

	err := person.CreatePerson(db)
	if err != nil {
		response := JsonResponse{Message: err.Error(), Status: http.StatusNoContent,}
		json.NewEncoder(w).Encode(response)
		return
	}

	mapResponse := map[string]any {"id": person.Id, "name": person.Name}
	response := JsonResponse{Data: mapResponse, Message: "successfully created new user", Status: http.StatusCreated,}
	json.NewEncoder(w).Encode(response)
}
// GET Request Function
func getFunc(w http.ResponseWriter, r *http.Request, db *sql.DB, id int, name string){
	var person models.Person
	err := person.ReadPerson(db, id, name)
	if err != nil {
		response := JsonResponse{Message: err.Error(), Status: http.StatusNoContent,}
		json.NewEncoder(w).Encode(response)
		return
	}
	mapResponse := map[string]any {"id": person.Id, "name": person.Name}
	response := JsonResponse{Data: mapResponse, Message: "successfully retrieved user from db", Status: http.StatusOK,}
	json.NewEncoder(w).Encode(response)
}
// PATCH Request Function
func patchFunc(w http.ResponseWriter, r *http.Request, db *sql.DB, id int, name string){
	var person models.Person
	err := person.UpdatePerson(db, id, name)
	if err != nil {
		response := JsonResponse{Message: err.Error(), Status: http.StatusNoContent,}
		json.NewEncoder(w).Encode(response)
		return
	}
	mapResponse := map[string]any {"id": person.Id, "name": person.Name}
	response := JsonResponse{Data: mapResponse, Message: "successfully updated user", Status: http.StatusOK,}
	json.NewEncoder(w).Encode(response)
}
// DELETE Request Function
func deleteFunc(w http.ResponseWriter, r *http.Request, db *sql.DB, id int){
	var person models.Person
	err := person.DeletePerson(db, id)
	if err != nil {
		response := JsonResponse{Message: err.Error(), Status: http.StatusNotFound,}
		json.NewEncoder(w).Encode(response)
		return
	}
	response := JsonResponse{Message: ("successfully deleted user with id: " + fmt.Sprintf("%d", id)), Status: http.StatusNoContent,}
	json.NewEncoder(w).Encode(response)
}

func createPersonTable(db *sql.DB){
	query := `CREATE TABLE IF NOT EXISTS persons ( 
		id INT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name VARCHAR(255) NOT NULL UNIQUE
		);`
	if _, err := db.Exec(query); err != nil {log.Fatal(err)}
}