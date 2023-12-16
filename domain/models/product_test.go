package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestProduct_JSONSerialization(t *testing.T) {
	p := Product{ID: 1, Price: 10.0, Name: "Produto"}
	data, err := json.Marshal(p)
	if err != nil {
		t.Errorf("Erro ao serializar o produto para JSON: %v", err)
	}
	expectedData := []byte(`{"id":1,"price":10,"name":"Produto"}`)
	if !reflect.DeepEqual(data, expectedData) {
		t.Errorf("Dados JSON esperados não encontrados: \n%s\nEsperado: \n%s", string(data), string(expectedData))
	}
}

func TestProduct_JSONDeserialization(t *testing.T) {
	jsonData := []byte(`{"id":1,"price":10,"name":"Produto"}`)
	var p Product
	err := json.Unmarshal(jsonData, &p)
	if err != nil {
		t.Errorf("Erro ao desserializar JSON para um produto: %v", err)
	}
	expectedProduct := Product{ID: 1, Price: 10.0, Name: "Produto"}
	if p != expectedProduct {
		t.Errorf("Produto desserializado não corresponde ao esperado: \n%v\nEsperado: \n%v", p, expectedProduct)
	}
}
