package handler

import (
	"net/http"

	"github.com/vanchaooo/go-marketplace/internal/model"
	"github.com/vanchaooo/go-marketplace/internal/service"
)

type WalletHandler struct {
	wallet *service.WalletService
}

func NewWalletHandler(wallet *service.WalletService) *WalletHandler {
	return &WalletHandler{wallet: wallet}
}

type walletResponse struct {
	UserID int64 `json:"user_id"`
	Balance int64 `json:"balance"`
}

type withdrawRequest struct {
	Amount int64 `json:"amount"`
}

type topUpRequest struct {
	Amount int64 `json:"amount"`
}

func walletToResponse(wallet model.Wallet) walletResponse {
	return walletResponse{
		UserID: wallet.UserID,
		Balance: wallet.Balance,
	}
}

func (h *WalletHandler) TopUp(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "некорректный ID")
		return
	}

	var request topUpRequest
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}
	
	wallet, err := h.wallet.TopUpBalance(r.Context(), id, request.Amount)
	if err != nil {
		handleServiceErrors(w, err)
		return
	}

	_ = writeJSON(w, http.StatusOK, walletToResponse(wallet))
}

func (h *WalletHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "некореектный ID")
		return
	}

	var request withdrawRequest
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, "некореектный JSON")
		return
	}
	
	wallet, err := h.wallet.WithdrawBalance(r.Context(), id, request.Amount)
	if err != nil {
		handleServiceErrors(w, err)
		return
	}

	_ = writeJSON(w, http.StatusOK, walletToResponse(wallet))
}

func (h *WalletHandler) GetWallet(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "некорректный ID")
		return
	}

	wallet, err := h.wallet.GetWallet(r.Context(), id)
	if err != nil {
		handleServiceErrors(w, err)
		return
	}

	_ = writeJSON(w, http.StatusOK, walletToResponse(wallet))
}