package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"my-module/postgres"
	"my-module/server"
	"net/http"
)

type dbHandler struct {
	db *sql.DB
}

func (h *dbHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		rows, err := postgres.GetRows(h.db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("HandleGet error: %v", err)
			return
		}
		var products []postgres.Product
		for rows.Next() {
			var p postgres.Product
			if err := rows.Scan(&p.ID, &p.Product, &p.Price, &p.CreatedAt); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Printf("Row scan error: %v", err)
				return
			}

			products = append(products, p)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Row iteration error: %v", err)
			return
		}

		json.NewEncoder(w).Encode(products)
	case "POST":
		var product postgres.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		postgres.InsertTable(h.db, product)
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer db.Close() // Close the connection when the program ends

	postgres.CreateTable(db)

	handler := &dbHandler{db: db}

	http.Handle("/", handler)
	server.Connect()

}
