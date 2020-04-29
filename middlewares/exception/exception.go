package exception

import "fmt"

// BadRequest handler
func BadRequest(message string, flag string) {
	fmt.Println(flag)
	errors := map[string]interface{}{
		"message": message, "flag": flag,
	}

	response := map[string]interface{}{
		"status": 400,
		"flag":   "BAD_REQUEST",
		"errors": errors,
	}

	panic(response)
}
