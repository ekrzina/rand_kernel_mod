package apiserve

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type ApiHandler struct {
	Port string
}

/* To be replaced with kernel module */
func getRandomNumber(w http.ResponseWriter, r *http.Request) {
	randomNumber := rand.Intn(10000)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(randomNumber)
}

/* Serves REST API on specified port */
func (t *ApiHandler) HandleRequests() {
	fmt.Printf("Server started on port %s...", t.Port)
	http.Handle("/randnumber", http.HandlerFunc(getRandomNumber))
	err := http.ListenAndServe(":"+t.Port, nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
