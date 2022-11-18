package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zueffc/mirage-node/internal/data/actions/packages"
)

type InformationQuery struct {
	Type string `json:"type"`
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func PackageInformationController(ctx *gin.Context) {
	var information InformationQuery
	if err := ctx.ShouldBindJSON(&information); err != nil {
		panic(err)
	}

	fmt.Println(information.Type)

	if information.Type == "all" {
		res := packages.FindByID(information.Type, information.Id)
		ctx.JSON(http.StatusOK, res)
	} else if information.Type == "current" {
		res := packages.FindByName(information.Name)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unknown type"})
	}

}
