package v0_1_0

import (
	"context"
	"fmt"
	"net/http"

	"com.ddabadi/estock/models"
	"com.ddabadi/estock/services"
	"github.com/gin-gonic/gin"
)

type ProductCategoryController struct {
}

func (controller *ProductCategoryController) GetAll(ctx *gin.Context) {
	fmt.Println(">>> ProductCategoryController - Get All <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}

	res = services.InitProductCategoryServiceInterface().GetAll()

	ctx.JSON(http.StatusOK, res)
}
