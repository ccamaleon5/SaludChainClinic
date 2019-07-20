package business
import (
	"fmt"
)

func getNameSubject(contentSubject interface{}) string {
	content := contentSubject.(map[string]interface{})
	name := content["name"]

	return fmt.Sprintf("%v", name)
}

func getLastNameSubject(contentSubject interface{}) string {
	content := contentSubject.(map[string]interface{})
	name := content["lastName"]

	return fmt.Sprintf("%v", name)
}

func getEmailSubject(contentSubject interface{}) string {
	content := contentSubject.(map[string]interface{})
	name := content["name"]

	return fmt.Sprintf("%v", name)
}

func getBirthDateSubject(contentSubject interface{}) string {
	content := contentSubject.(map[string]interface{})
	name := content["birthdate"]

	return fmt.Sprintf("%v", name)
}