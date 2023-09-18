package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func parseIntParameter(stringParameter string, paramterName string, w http.ResponseWriter) (int, error) {
	nilIntValue := -1

	if stringParameter != "" {
		intValue, err := strconv.Atoi(stringParameter)
		fmt.Println("intValue:", intValue)
		if err != nil {
			message := fmt.Sprintf("Invalid '%s parameter", paramterName)
			respondWithError(w, http.StatusBadRequest, message)
			return nilIntValue, err
		}
		return intValue, nil
	} else {
		return nilIntValue, nil
	}
}
