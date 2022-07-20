package scopes

import (
	"math"
	"strconv"

	"github.com/EusRique/frajola/application/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Paginate(c *gin.Context, value interface{}, pagination model.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64

	db.Model(value).Count(&totalRows)
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

	pagination.Limit = pageSize      // Quantidade de registros a retornar
	pagination.TotalRows = totalRows // Total de registros na base de dados
	pagination.Page = page           // Pagina recebida no param
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages // Total de paginas
	pagination.Sort = sort             // Ordenação dos registros

	c.Set("pagination", pagination)

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
