package exception

import "net/http"

// Empty -> response for empty result
func Empty(msg string, message string, flag string) {
	errors := map[string]interface{}{
		"message": message, "flag": flag,
	}

	response := map[string]interface{}{
		"status":  http.StatusOK,
		"message": msg,
		"errors":  errors,
	}

	panic(response)
}

// BadRequest -> response for bad request
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

// InternalServerError -> response for internal server error
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
