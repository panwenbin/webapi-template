package pagination_srv

import (
	"app/structs/requests"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Bind binds pagination queries in request to a gorm Query
func Bind(c *gin.Context, query *gorm.DB) *gorm.DB {
	pagi := requests.Pagination{
		Page:    1,
		PerPage: 10,
	}
	_ = c.BindQuery(&pagi)

	query = query.Offset((pagi.Page - 1) * pagi.PerPage).Limit(pagi.PerPage)
	if pagi.Order != "" {
		query = query.Order(pagi.Order)
	}

	return query
}
