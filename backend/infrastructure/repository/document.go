package repository

import (
	"errors"
	"math"

	modelApplication "github.com/EusRique/frajola/application/model"
	"github.com/EusRique/frajola/domain/model"
	"github.com/EusRique/frajola/utils/scopes"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type DocumentRepositoryDb struct {
	Db *gorm.DB
}

func (d *DocumentRepositoryDb) AddDocument(document *model.Document) error {
	err := d.Db.Create(document).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *DocumentRepositoryDb) FindDocumentByNumber(documentNumber string) (*model.Document, error) {
	var document model.Document

	resultQuery := d.Db.First(&document, "document = ?", documentNumber).Error

	if errors.Is(resultQuery, gorm.ErrRecordNotFound) {
		return &document, nil
	}

	if resultQuery != nil {
		return &document, resultQuery
	}

	return &document, nil
}

func (d *DocumentRepositoryDb) SearchAllDocument(c *gin.Context) ([]model.Document, error) {
	var documents []model.Document
	var pagination modelApplication.Pagination
	var totalRows int64

	err := d.Db.Scopes(
		scopes.Paginate(c, documents, pagination, d.Db),
		scopes.Filters(c)).
		Find(&documents).Error

	pagination = c.MustGet("pagination").(modelApplication.Pagination)
	d.Db.Scopes(
		scopes.Filters(c)).
		Model(&model.Document{}).Count(&totalRows)

	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	c.Set("pagination", pagination)
	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("Document not found")
	}

	return documents, nil
}

func (d *DocumentRepositoryDb) DeleteDocumentById(documentId string) error {
	var document model.Document

	err := d.Db.Where("id = ?", documentId).Delete(&document).Error
	if err != nil {
		return errors.New("Database error, please try again")
	}

	return nil
}

func (d *DocumentRepositoryDb) FindDocumentById(id string) (*model.Document, error) {
	var document model.Document

	resultQuery := d.Db.First(&document, "id = ?", id).Error

	if errors.Is(resultQuery, gorm.ErrRecordNotFound) {
		return &document, nil
	}

	if resultQuery != nil {
		return &document, resultQuery
	}

	return &document, nil
}

func (d *DocumentRepositoryDb) UpdateDocument(document *model.Document) (*model.Document, error) {
	err := d.Db.Save(&document).Error
	if err != nil {
		return nil, err
	}

	return document, nil
}
