package helpers

import (
	"strconv"

	"github.com/cryptosalamander/gorm_crud_example/dtos"
	"github.com/gin-gonic/gin"
)

func GeneratePaginationRequest(context *gin.Context) *dtos.Pagination {
	// convert query parameter string to int
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(context.DefaultQuery("page", "0"))
	sort := context.DefaultQuery("sort", "created_at desc")

	return &dtos.Pagination{Limit: limit, Page: page, Sort: sort}
}
