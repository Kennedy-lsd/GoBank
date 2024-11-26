package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Kennedy-lsd/GoBank/data"
)

type BalanceHandler struct {
	BalanceRepo data.BalanceRepository
}

func NewBalanceHandler(r data.BalanceRepository) *BalanceHandler {
	return &BalanceHandler{
		BalanceRepo: r,
	}
}

func (h *BalanceHandler) GetAllBalances(w http.ResponseWriter, r *http.Request) {
	balances, err := h.BalanceRepo.GetAll()

	if len(balances) == 0 {
		http.Error(w, "Balances not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	writeJSON(w, 200, balances)
}

func (h *BalanceHandler) CreateBalance(w http.ResponseWriter, r *http.Request) {
	var balance data.Balance
	if err := json.NewDecoder(r.Body).Decode(&balance); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdBalance, err := h.BalanceRepo.Create(&balance)
	if err != nil {
		http.Error(w, "Failed to create balance: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, 201, createdBalance)
}
