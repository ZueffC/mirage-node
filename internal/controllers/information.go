package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zueffc/mirage-node/internal/data/actions/packages"
)

type InformationQuery struct {
	Type string `json:"type"`
	Id   uint   `json:"id"`
}

func PackageInformationController(ctx *gin.Context) {
	var information InformationQuery
	if err := ctx.ShouldBindJSON(&information); err != nil {
		panic(err)
	}

	if information.Type != "all" || information.Type != "current" {
		res := packages.Find(information.Type, information.Id)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unknown type"})
	}

}
