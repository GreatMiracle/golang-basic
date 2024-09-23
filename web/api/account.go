package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	db "simpleBank/db/sqlc"
	"simpleBank/token"
	"strconv"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,valid_currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.Queries.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) getAccount(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	account, err := server.store.GetAccount(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	payload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if payload.Username != account.Owner {
		err = errors.New("account doesn't belong to the authentication user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID   int32   `form:"page_id" binding:"required,min=1"`
	PageSize int32   `form:"page_size" binding:"required,min=1"`
	Owner    *string `form:"owner"`
}

func (server *Server) getListAccount(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil { // Sử dụng ShouldBindQuery
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := db.ListAccountsParams{
		Limit:   req.PageSize,
		Offset:  (req.PageID - 1) * req.PageSize,
		Column1: req.Owner,
	}

	// Truyền owner vào params
	if req.Owner != nil && *req.Owner != "" {
		params.Column1 = "%" + *req.Owner + "%"
	} else {
		params.Column1 = "" // Hoặc không gán gì để giữ nguyên
	}

	accounts, err := server.store.ListAccounts(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusNoContent, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
