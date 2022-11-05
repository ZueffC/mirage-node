package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zueffc/mirage-node/internal/data/actions/users"
)

type query struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

func DeleteController(ctx *gin.Context) {
	var query query

	if err := ctx.ShouldBindJSON(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	err := users.Delete(query.Id, query.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
