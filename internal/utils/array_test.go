package utils

import (
	"testing"
)

func TestAGet1(t *testing.T) {
	array := map[string]interface{}{
		"response": map[string]interface{}{
			"success": "Success Message",
			"error":   "Error Message",
		},
	}
	result := AGet(array, "response.success")
	if result != "Success Message" {
		t.Fatal()
	}

}

func TestAGet2(t *testing.T) {
	data := map[string]interface{}{
		"response": map[string]interface{}{
			"success": true,
			"message": "Data retrieved successfully",
			"data": []interface{}{
				map[string]interface{}{
					"id":   1,
					"name": "John",
				},
				map[string]interface{}{
					"id":   2,
					"name": "Jane",
				},
			},
		},
	}

	result := AGet(data, "response.success")

	if result != true {
		t.Fatal()
	}
}
