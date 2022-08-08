package helper

import (
	"mini-project/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GeneratePaginationRequest(context *gin.Context) *dto.Pagination {
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(context.DefaultQuery("page", "0"))
	return &dto.Pagination{Limit: limit, Page: page}
}
