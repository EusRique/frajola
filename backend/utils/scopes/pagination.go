package scopes

import (
	"strconv"

	"github.com/EusRique/frajola/application/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Paginate(c *gin.Context, value interface{}, pagination model.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	sort := "id " + c.Query("sort")

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 50
	}

	pagination.Limit = pageSize // Quantidade de registros a retornar
	pagination.Page = page      // Pagina recebida no param
	pagination.Sort = sort      // Ordenação dos registros

	c.Set("pagination", pagination)

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
