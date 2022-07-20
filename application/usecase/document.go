package usecase

import (
	"errors"
	"regexp"

	"github.com/EusRique/frajola/domain/model"
	"github.com/gin-gonic/gin"
)

type DocumentUseCase struct {
	DocumentRepository model.DocumentRepositoryInterface
}

func (d *DocumentUseCase) AddDocument(documentNumber string, documentType string, IsBlockList bool) (*model.Document, error) {
	newDocument, err := model.NewDocument(documentNumber, documentType, IsBlockList)
	if err != nil {
		return nil, err
	}

	document, err := d.DocumentRepository.FindDocumentByNumber(newDocument.Document)
	if err != nil {
		return nil, err
	}

	if document.ID == "" {
		d.DocumentRepository.AddDocument(newDocument)
		if newDocument.ID == "" {
			return nil, err
		}
		return newDocument, nil
	}

	return nil, errors.New("Document already registered in our database")
}

func (d *DocumentUseCase) SearchAllDocument(c *gin.Context) ([]model.Document, error) {
	listDocument, err := d.DocumentRepository.SearchAllDocument(c)
	if err != nil {
		return nil, err
	}

	return listDocument, nil
}

func (d *DocumentUseCase) DeleteDocument(id string) error {
	document, err := d.DocumentRepository.FindDocumentById(id)
	if err != nil {
		return err
	}

	if document.ID == "" {
		return errors.New("Document not found")
	}

	err = d.DocumentRepository.DeleteDocumentById(id)
	if err != nil {
		return err
	}

	return nil
}

func (d *DocumentUseCase) UpdateDocument(documentNumber string, documentType string, IsBlockList bool, documentId string) (*model.Document, error) {
	document, err := d.DocumentRepository.FindDocumentById(documentId)
	if err != nil {
		return nil, err
	}

	if document.ID == "" {
		return nil, errors.New("Document not found")
	}

	cnpjOrCpf := documentNumber
	regex, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return nil, err
	}
	cnpjOrCpfFormated := regex.ReplaceAllString(cnpjOrCpf, "")

	var newDocument *model.Document
	newDocument = &model.Document{
		ID:           documentId,
		Document:     cnpjOrCpfFormated,
		DocumentType: documentType,
		IsBlockList:  IsBlockList,
	}

	err = newDocument.IsValid()
	if err != nil {
		return nil, err
	}

	_, err = d.DocumentRepository.UpdateDocument(newDocument)
	if err != nil {
		return nil, err
	}

	return newDocument, nil
}
