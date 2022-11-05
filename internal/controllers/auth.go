package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zueffc/mirage-node/internal/data/actions/users"
)

type Query struct {
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthController(ctx *gin.Context) {
	var query Query

	if err := ctx.ShouldBindJSON(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := users.Create(query.Nick, query.Email, query.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
