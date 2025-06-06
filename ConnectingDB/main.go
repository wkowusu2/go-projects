package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type User struct {
	ID         int    `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
}

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	dbString := os.Getenv("POSTGRES_URL")
	if dbString == "" {
		fmt.Println("DB_CONNECTION_STRING is not set in .env file")
		return
	}
	db, err = sql.Open("postgres", dbString)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Unable to connect to DB:", err)
	}

	fmt.Println("✅ Successfully connected to database!")
	defer db.Close()
	//setting up the server
	mu := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mu,
	}
	mu.HandleFunc("/users", handleHome)
	server.ListenAndServe()
	// fmt.Println("Server is running on port 8080")
	// query := "SELECT * FROM test"
	// insertQuery := "INSERT INTO test (one) VALUES (1)"
	// _, err = db.Exec(insertQuery)
	// updateQuery := "UPDATE test SET one = 2 WHERE one = 1"
	// _, err = db.Exec(updateQuery)
	// deleteQuery := "DELETE FROM test WHERE one = 2"
	// _, err = db.Exec(deleteQuery)
	// rows, err := db.Query(query)
	// if err != nil {
	// 	log.Fatal("Error executing query:", err)
	// }
	// defer rows.Close()
	// fmt.Println("Query executed successfully!")
	// // fmt.Println("Rows returned:", rows)
	// for rows.Next() {
	// 	var one int
	// 	if err := rows.Scan(&one); err != nil {
	// 		log.Fatal("Error scanning row:", err)
	// 	}
	// 	fmt.Printf("ONE: %d\n", one)
	// }

}
func handleHome(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM users"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close()
	fmt.Println("Query executed successfully!")
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.First_Name, &user.Last_Name, &user.Email, &user.Gender); err != nil {
			log.Fatal("Error scanning row:", err)
		}
		users = append(users, user)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
