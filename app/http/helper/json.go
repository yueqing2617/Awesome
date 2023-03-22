package helper

import (
	"Awesome/app/models"
	"encoding/json"
)

// ProceduresToJson 处理存储过程
func ProceduresToJson(procedures *[]models.Procedure) string {
	var result string
	if len(*procedures) > 0 {
		b, err := json.Marshal(procedures)
		if err != nil {
			return err.Error()
		}
		result = string(b)
	}
	return result
}

// JsonToProcedures 处理存储过程
func JsonToProcedures(jsonStr string) []models.Procedure {
	var result []models.Procedure
	if jsonStr != "" {
		_ = json.Unmarshal([]byte(jsonStr), &result)
	}
	return result
}

// ContainsToJSON  contains to json
func ContainsToJSON(contains *[]models.ClothOrderContains) string {
	var result string
	if len(*contains) > 0 {
		b, err := json.Marshal(contains)
		if err != nil {
			return err.Error()
		}
		result = string(b)
	}
	return result
}

// JSONToContains  json to contains
func JSONToContains(jsonStr string) []models.ClothOrderContains {
	var result []models.ClothOrderContains
	if jsonStr != "" {
		_ = json.Unmarshal([]byte(jsonStr), &result)
	}
	return result
}
