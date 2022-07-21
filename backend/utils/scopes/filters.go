package scopes

import (
	"strconv"

	"github.com/EusRique/frajola/domain/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Filters(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var filter model.Document

		documentNumber := c.Query("document_number")
		if documentNumber != "" {
			filter.Document = documentNumber
		}

		documentType := c.Query("document_type")
		if documentType != "" {
			filter.DocumentType = documentType
		}

		isBlockList := c.Query("is_block_list")
		isBlockListBool, _ := strconv.ParseBool(isBlockList)
		if isBlockList != "" {
			filter.IsBlockList = isBlockListBool
		}

		return db.Where(&filter)
	}
}