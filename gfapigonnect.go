package gfapigonnect

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func errorAPI(e error) {
	if e != nil {
		log.Fatalln(e)
		fmt.Fprintf(os.Stderr, "Error: %s\n", e)
	}
}

// API ...
type API struct {
	BaseURL    string
	KeyPublic  string
	KeyPrivate string
}

// calculateSignature ...
// val = value to sign
// key = private key
func (a *API) calculateSignature(val string, key string) string {

	sig := hmac.New(sha1.New, []byte(key))
	sig.Write([]byte(val))
	mac := sig.Sum(nil)
	str := base64.StdEncoding.EncodeToString(mac)
	esc := url.QueryEscape(str)

	return esc

}

// createQueryURL ...
func (a *API) createQueryURL(formID string, formType string) string {

	route := "forms/" + formID + "/" + formType
	time := time.Now().Add(time.Duration(1) * time.Hour).Unix()
	expiry := strconv.FormatInt(time, 10)
	val := fmt.Sprintf("%s:%s:%s:%s", a.KeyPublic, "GET", route, expiry)
	sig := a.calculateSignature(val, a.KeyPrivate)

	// https://www.gravityhelp.com/documentation/article/web-api-examples/#retrieve-entries-using-field-filters-one-condition
	query := a.BaseURL + route + "/" + "?api_key=" + a.KeyPublic + "&signature=" + sig + "&expires=" + expiry + "&paging[page_size]=1000" // + "&search=" + filter

	return query
}

// Call ...
// Query GravityForms API
func (a *API) Call(formID string, formType string) string {

	query := a.createQueryURL(formID, formType)
	fmt.Println(query)

	resp, err := http.Get(query)
	errorAPI(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	errorAPI(err)

	return string(body)

}
