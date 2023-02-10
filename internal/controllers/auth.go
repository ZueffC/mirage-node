package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zueffc/mirage-node/internal/data/actions/users"
)

type UserQuery struct {
	Type     string `json:"type"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthController(ctx *gin.Context) {
	var query UserQuery
	low := strings.ToLower

	if err := ctx.ShouldBindJSON(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if query.Type == "registration" {
		err := users.Create(query.Nick, query.Email, low(query.Password))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, gin.H{})
	} else if query.Type == "login" {
		res := users.Find(query.Nick, low(query.Password), query.Email)
		if len(res.Nick) > 0 {
			ctx.JSON(http.StatusOK, res)
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid type"})
	}
}
