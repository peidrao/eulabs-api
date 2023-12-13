package utils

import "github.com/peidrao/eulabs-api/domain/models"

type Response struct {
	Status   int         `json:"status"`
	Messages string      `json:"messages"`
	Data     interface{} `json:"data"`
}

type ProductResponse struct {
	Product models.Product `json:"product"`
}
