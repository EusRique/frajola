package model

import "time"

type Document struct {
	DocumentNumber string `json:"document_number" binding:"required"`
	DocumentType   string `json:"document_type" binding:"required"`
	IsBlockList    bool   `json:"is_block_list"`
}

type ListAllDocuments struct {
	DocumentNumber string     `json:"document_number"`
	DocumentType   string     `json:"document_type"`
	IsBlockList    bool       `json:"is_block_list"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
