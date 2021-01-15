package handler

import (
	"fmt"
	"net/http"
)

//default handler function
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Please try 'website-details' path to get the website details.")
}
