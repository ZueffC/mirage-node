package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zueffc/mirage-node/internal/data/actions/packages"
)

type PackageQuery struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AutorID     uint   `json:"author_id"`
	GitURL      string `json:"git_url"`
}

func PackagesController(ctx *gin.Context) {
	var query PackageQuery

	if err := ctx.ShouldBindJSON(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if query.Type == "new_package" {
		err := packages.Create(query.Name, query.Description, query.GitURL, query.AutorID)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, gin.H{})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unknown error"})
	}
}
