package rest

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/EusRique/frajola/application/factory"
	"github.com/EusRique/frajola/application/model"
	"github.com/EusRique/frajola/application/usecase"
	"github.com/EusRique/frajola/utils/csv"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//DocumentRestService struct defines the dependencies that will be used
type DocumentRestService struct {
	DocumentUseCase usecase.DocumentUseCase
}

//Users constructor
func NewUsers(DocumentUseCase usecase.DocumentUseCase) *DocumentRestService {
	return &DocumentRestService{
		DocumentUseCase: DocumentUseCase,
	}
}

func (d *DocumentRestService) CreateDocument(c *gin.Context) {
	var document model.Document

	if err := c.BindJSON(&document); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "JSON Invalido"})
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	documentUseCase := factory.DocumentUseCaseFactory(db)
	newDocument, err := documentUseCase.AddDocument(document.DocumentNumber, document.DocumentType, document.IsBlockList)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"results": newDocument})
}

func (d *DocumentRestService) ListAllDocuments(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)
	documentUseCase := factory.DocumentUseCaseFactory(db)
	documents, err := documentUseCase.SearchAllDocument(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if documents == nil {
		c.JSON(http.StatusOK, gin.H{"Result": documents})
		return
	}

	listDocuments := make([]interface{}, len(documents))
	for index, document := range documents {
		listDocuments[index] = document
	}

	paginate := c.MustGet("pagination")

	c.JSON(http.StatusCreated, gin.H{"pagination": paginate, "results": listDocuments})
}

func (d *DocumentRestService) DeleteDocument(c *gin.Context) {
	documentId := c.Param("document_id")
	db := c.MustGet("DB").(*gorm.DB)
	documentUseCase := factory.DocumentUseCaseFactory(db)
	err := documentUseCase.DeleteDocument(documentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": "Document deleted"})
}

func (d *DocumentRestService) UpdateDocument(c *gin.Context) {
	var document model.Document
	documentId := c.Param("document_id")
	if err := c.BindJSON(&document); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "JSON Invalido"})
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	documentUseCase := factory.DocumentUseCaseFactory(db)
	newDocument, err := documentUseCase.UpdateDocument(document.DocumentNumber, document.DocumentType, document.IsBlockList, documentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": newDocument})
}

func (d *DocumentRestService) ExportDocument(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)
	var documents [][]string
	documentUseCase := factory.DocumentUseCaseFactory(db)
	response, err := documentUseCase.SearchAllDocument(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	header := []string{
		"id",
		"numero_documento",
		"tipo_documento",
		"bloqueado",
	}

	documents = append(documents, header)

	for _, document := range response {
		body := []string{
			document.ID,
			document.Document,
			document.DocumentType,
			strconv.FormatBool(document.IsBlockList),
		}

		documents = append(documents, body)
	}

	nameFile := "tmp/" + time.Now().Format("20060102150405") + ".csv"
	err = csv.ExporToCSV(c, documents, nameFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = csv.DownloadFile(c, nameFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	os.Remove(nameFile)
}
