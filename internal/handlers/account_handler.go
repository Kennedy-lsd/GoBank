package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Kennedy-lsd/GoBank/data"
	"github.com/Kennedy-lsd/GoBank/utils"
)

type AccountHandler struct {
	AccountRepo data.AccountRepository
}

func NewAccountHandler(r data.AccountRepository) *AccountHandler {
	return &AccountHandler{
		AccountRepo: r,
	}
}

func (h *AccountHandler) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.AccountRepo.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch accounts", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	writeJSON(w, 200, accounts)
}

func (h *AccountHandler) GetOneAccount(w http.ResponseWriter, r *http.Request) {

	id, err := idParser(w, r)

	if err != nil {
		return
	}

	account, err := h.AccountRepo.GetOne(id)
	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}
	writeJSON(w, 200, account)

}

func (h *AccountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	id, err := idParser(w, r)

	if err != nil {
		return
	}

	if err := h.AccountRepo.Delete(id); err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	writeJSON(w, 200, map[string]interface{}{
		"message": "Account deleted successfully",
	})
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) { //for test

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body:", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	log.Println("Raw request body:", string(body))

	// Reassign body to enable re-reading
	r.Body = io.NopCloser(bytes.NewReader(body))

	account := new(data.Account)
	err = json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if account.Name == "" || account.Age == 0 {
		log.Println("Validation error: missing required fields")
		http.Error(w, "Name, Age, and Amount are required", http.StatusBadRequest)
		return
	}

	err = h.AccountRepo.Create(account)
	if err != nil {
		log.Println("Error creating account:", err)
		http.Error(w, "Failed to create account", http.StatusInternalServerError)
		return
	}

	writeJSON(w, 201, map[string]interface{}{
		"message": "Account created successfully",
		"id":      account.Id,
	})
}

func (h *AccountHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")

	if len(pathSegments) < 3 {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	name := pathSegments[3]

	_, err := h.AccountRepo.GetByName(name)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	token, err := utils.GenerateToken(name)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	writeJSON(w, 200, map[string]interface{}{
		"token": token,
	})
}
