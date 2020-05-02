package exception

import "net/http"

// BadRequest handler
func BadRequest(message string, flag string) {
	errors := map[string]interface{}{
		"message": message, "flag": flag,
	}

	response := map[string]interface{}{
		"status": http.StatusBadRequest,
		"flag":   "BAD_REQUEST",
		"errors": errors,
	}

	panic(response)
}

// InternalServerError handler
func InternalServerError(message string, flag string) {
	errors := map[string]interface{}{
		"message": message, "flag": flag,
	}

	response := map[string]interface{}{
		"status": http.StatusInternalServerError,
		"flag":   "INTERNAL_SERVER_ERROR",
		"errors": errors,
	}

	panic(response)
}
