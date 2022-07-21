package model

import (
	"errors"
	"regexp"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/klassmann/cpfcnpj"
	uuid "github.com/satori/go.uuid"
)

type DocumentRepositoryInterface interface {
	AddDocument(document *Document) error
	FindDocumentByNumber(document string) (*Document, error)
	FindDocumentById(id string) (*Document, error)
	SearchAllDocument(c *gin.Context) ([]Document, error)
	UpdateDocument(document *Document) (*Document, error)
	DeleteDocumentById(id string) error
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Document struct {
	ID           string     `gorm:"type:uuid;primary_key" json:"id" valid:"uuid"`
	Document     string     `gorm:"type:varchar(20)" json:"document" valid:"notnull"`
	DocumentType string     `gorm:"type:varchar(5)" json:"document_type" valid:"notnull"`
	IsBlockList  bool       `gorm:"type:bool" json:"is_block_list" valid:"-"`
	CreatedAt    time.Time  `gorm:"type:timestamp;autoCreateTime;default:CURRENT_TIMESTAMP" json:"created_at" valid:"-"`
	UpdatedAt    *time.Time `gorm:"type:timestamp;autoUpdateTime" json:"updated_at" valid:"-"`
}

func (document *Document) IsValid() error {
	_, err := govalidator.ValidateStruct(document)

	if document.DocumentType != "CPF" && document.DocumentType != "CNPJ" {
		return errors.New("Invalid type document")
	}

	switch document.DocumentType {
	case "CNPJ":
		cnpj := cpfcnpj.NewCNPJ(document.Document)
		if !cnpj.IsValid() {
			return errors.New("It is not a valid document")
		}
	case "CPF":
		cpf := cpfcnpj.NewCPF(document.Document)
		if !cpf.IsValid() {
			return errors.New("It is not a valid document")
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func NewDocument(documentNumber string, documentType string, isBlockList bool) (*Document, error) {
	cnpjOrCpf := documentNumber
	regex, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return nil, err
	}
	cnpjOrCpfFormated := regex.ReplaceAllString(cnpjOrCpf, "")

	document := Document{
		Document:     cnpjOrCpfFormated,
		DocumentType: documentType,
		IsBlockList:  isBlockList,
	}
	document.ID = uuid.NewV4().String()

	err = document.IsValid()
	if err != nil {
		return nil, err
	}

	return &document, nil
}
