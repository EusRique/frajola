package model_test

import (
	"errors"
	"testing"

	"github.com/EusRique/frajola/domain/model"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
)

func TestNewDocument(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NewDocument Suite")
}

var _ = Describe("Document.isValid() validate struct", func() {
	Context("When document validation does not return an error", func() {
		It("Returns nil", func() {
			validateDocument := model.Document{
				ID:           "bb006978-ca82-4e83-81d4-797a9355e724",
				Document:     "07694278048",
				DocumentType: "CPF",
				IsBlockList:  false,
			}

			response := validateDocument.IsValid()
			Expect(response).To(BeNil())
		})
	})

	Context("When document validation return an error of struct", func() {
		It("Returns error", func() {
			validateDocument := model.Document{
				DocumentType: "CPF",
				IsBlockList:  false,
			}

			response := validateDocument.IsValid()
			Expect(response).To(MatchError(response.Error()))
		})
	})
})

var _ = Describe("Document.isValid() validate business roles", func() {
	err := errors.New("Invalid type document")
	errCNPJorCPFInvalid := errors.New("It is not a valid document")

	Context("When the document type is different from CPF or CNPJ", func() {
		It("Returns an error", func() {
			document := model.Document{
				DocumentType: "CNH",
			}

			response := document.IsValid()
			Expect(response).To(MatchError(err))
		})
	})

	Context("When the CNPJ is invalid", func() {
		It("Returns an error", func() {
			validateDocument := model.Document{
				DocumentType: "CNPJ",
				Document:     "000123456098761",
			}

			response := validateDocument.IsValid()
			Expect(response).To(MatchError(errCNPJorCPFInvalid))
		})

	})

	Context("When the CPF is invalid", func() {
		It("Return an error", func() {
			validateDocument := model.Document{
				DocumentType: "CPF",
				Document:     "00011122233",
			}

			response := validateDocument.IsValid()
			Expect(response).To(MatchError(errCNPJorCPFInvalid))
		})
	})

})

var _ = Describe("NewDocument", func() {
	var document *model.Document

	document = &model.Document{
		ID:           uuid.NewV4().String(),
		Document:     "076.942.780-48",
		DocumentType: "CPF",
		IsBlockList:  false,
	}

	_, err := model.NewDocument(document.Document, document.DocumentType, document.IsBlockList)
	Context("We wait for the correct values", func() {
		It("When we call document in regex", func() {
			Expect(document.Document).To(MatchRegexp("[^a-zA-Z0-9]+"))
		})
	})

	// Context("We wait for the correct values", func() {
	// 	It("When we call newDocument", func() {
	// 		Expect(document).To(Equal(newDocument))
	// 	})
	// })

	Context("We must return a newDocument", func() {
		When("The error return is empty", func() {
			It("Returns a nil error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Context("We must return a newDocument", func() {
		When("The error return not is empty", func() {
			It("Returns an error", func() {
				_, err := model.NewDocument("", "", document.IsBlockList)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
