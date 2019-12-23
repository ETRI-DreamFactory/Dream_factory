package langs

import "fmt"

func GenerateValidationMessage(field string, rule string) (message string) {
	switch rule {
	case "required":
		return fmt.Sprintf("Field '%s' is '%s'.",field, rule)
	default:
		return fmt.Sprintf("Field '%s' is not valid.", field)
	}
}