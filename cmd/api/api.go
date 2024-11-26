package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Kennedy-lsd/GoBank/config"
	"github.com/Kennedy-lsd/GoBank/database"
	"github.com/Kennedy-lsd/GoBank/internal/handlers"
	"github.com/Kennedy-lsd/GoBank/internal/repos"
	"github.com/Kennedy-lsd/GoBank/utils"
)

func Api() {
	config.InitConfig()
	db, err := database.Init()
	if err != nil {
		log.Fatal(err)
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	accountRepo := repos.NewAccoutRepo(db)
	accountHandler := handlers.NewAccountHandler(accountRepo)

	balanceRepo := repos.NewBalanceRepo(db)
	balanceHandler := handlers.NewBalanceHandler(balanceRepo)

	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("/api/accounts/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			accountHandler.GetAllAccounts(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			accountHandler.GetOneAccount(w, r)
		} else if r.Method == http.MethodDelete {
			accountHandler.DeleteAccount(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/login/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			accountHandler.AuthenticateUser(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			accountHandler.CreateAccount(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	// Protected routes (JWT-protected)
	mux.Handle("/api/balances/", utils.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			balanceHandler.GetAllBalances(w, r)
		} else if r.Method == http.MethodPost {
			balanceHandler.CreateBalance(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})))

	log.Printf("Server starts on %v", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil); err != nil {
		log.Fatalf("Error starting server on %v", PORT)
	}
}
