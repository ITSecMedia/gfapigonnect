package main

import (
	"fmt"
	"log"
	"net/http"
	"ITSecMedia/gfapigonnect"
)

func queryGravityFormsAPI(w http.ResponseWriter, r *http.Request) {

	var gf API

	gf.BaseURL = "http://<wordpressblog_domain>/gravityformsapi/"
	gf.KeyPublic = "<public_key>"
	gf.KeyPrivate = "<private_key>"

	gfID := "<gf_form_id_int_value_check_wordpress>" // Form ID
	gfType := "results"           // Result Type - Read docs as there are more than one

	j := gf.Call(gfID, gfType)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, j)

}

func main() {

	http.HandleFunc("/", queryGravityFormsAPI)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
