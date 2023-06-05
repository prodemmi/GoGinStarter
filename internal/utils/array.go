package utils

import (
	"strconv"
	"strings"
)

func AGet(data interface{}, key string) interface{} {
	keys := strings.Split(key, ".")

	for _, k := range keys {
		switch v := data.(type) {
		case map[string]interface{}:
			data = v[k]
		case []interface{}:
			i, err := strconv.Atoi(k)
			if err != nil || i < 0 || i >= len(v) {
				return nil
			}
			data = v[i]
		default:
			return nil
		}
	}

	return data
}

func AKeyExists(i interface{}, key string) bool {
	m, ok := i.(map[string]interface{})
	if !ok {
		return false
	}
	_, ok = m[key]

	return ok
}
