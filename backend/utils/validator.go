package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func validateBody(body map[string]interface{}, schema map[string]interface{}) error {
	for key, ruleRaw := range schema {
		rulesStr, ok := ruleRaw.(string)
		if !ok {
			return fmt.Errorf("schema của field %s không đúng định dạng", key)
		}

		rules := strings.Split(rulesStr, "|")

		val, exists := body[key]

		// Required check
		if contains(rules, "required") && !exists {
			return fmt.Errorf("field %s là bắt buộc", key)
		}

		// Skip if not exists and not required
		if !exists {
			continue
		}

		// Type check
		expectedType := rules[0]
		actualKind := reflect.TypeOf(val).Kind().String()

		if expectedType != actualKind {
			// Hỗ trợ tự động parse int từ float64 nếu từ JSON
			if expectedType == "int" && actualKind == "float64" {
				// OK
			} else {
				return fmt.Errorf("field %s phải là kiểu %s, nhưng là %s", key, expectedType, actualKind)
			}
		}

		// Validate logic
		for _, rule := range rules[1:] {
			switch {
			case strings.HasPrefix(rule, "min="):
				minVal, _ := strconv.Atoi(strings.TrimPrefix(rule, "min="))
				if expectedType == "string" {
					if len(val.(string)) < minVal {
						return fmt.Errorf("field %s phải có ít nhất %d ký tự", key, minVal)
					}
				} else if expectedType == "int" || expectedType == "float" {
					if getNumber(val) < float64(minVal) {
						return fmt.Errorf("field %s phải >= %d", key, minVal)
					}
				}
			case strings.HasPrefix(rule, "max="):
				maxVal, _ := strconv.Atoi(strings.TrimPrefix(rule, "max="))
				if expectedType == "string" {
					if len(val.(string)) > maxVal {
						return fmt.Errorf("field %s không được vượt quá %d ký tự", key, maxVal)
					}
				} else if expectedType == "int" || expectedType == "float" {
					if getNumber(val) > float64(maxVal) {
						return fmt.Errorf("field %s phải <= %d", key, maxVal)
					}
				}
			case strings.HasPrefix(rule, "length="):
				expectLen, _ := strconv.Atoi(strings.TrimPrefix(rule, "length="))
				if len(val.(string)) != expectLen {
					return fmt.Errorf("field %s phải có đúng %d ký tự", key, expectLen)
				}
			case strings.HasPrefix(rule, "enum="):
				options := strings.Split(strings.TrimPrefix(rule, "enum="), ",")
				if !contains(options, fmt.Sprintf("%v", val)) {
					return fmt.Errorf("field %s phải nằm trong các giá trị: %v", key, options)
				}
			}
		}
	}

	return nil
}

func contains(list []string, target string) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

func getNumber(v interface{}) float64 {
	switch t := v.(type) {
	case float64:
		return t
	case int:
		return float64(t)
	case int64:
		return float64(t)
	default:
		return 0
	}
}
