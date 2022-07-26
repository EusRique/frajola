package scopes

import (
	"regexp"
	"strconv"

	"github.com/EusRique/frajola/domain/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Filters(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var filter model.Document

		documentNumber := c.Query("document_number")
		if documentNumber != "" && documentNumber != "null" {
			cnpjOrCpf := documentNumber
			regex, _ := regexp.Compile("[^a-zA-Z0-9]+")
			cnpjOrCpfFormated := regex.ReplaceAllString(cnpjOrCpf, "")
			filter.Document = cnpjOrCpfFormated
		}

		documentType := c.Query("document_type")
		if documentType != "" && documentType != "null" {
			filter.DocumentType = documentType
		}

		isBlockList := c.Query("is_block_list")
		isBlockListBool, _ := strconv.ParseBool(isBlockList)
		if isBlockList != "" && isBlockList != "null" {
			filter.IsBlockList = isBlockListBool
		}

		return db.Where(&filter)
	}
}
