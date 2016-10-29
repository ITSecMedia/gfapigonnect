# gfapigonnect — Go package to query Gravity Forms Web API

A simple **Go-Lang package** to query the **Gravity Forms Web API**

## API documentation

The documentation for the Gravity Forms API can be found here:

[https://www.gravityhelp.com/documentation/article/gravity-forms-api/](https://www.gravityhelp.com/documentation/article/gravity-forms-api/)

### Example

All you have to do is to supply 3 variables:

- URL to your blog with your gravityforms installation
- Key Public
- Key Private

```
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/ITSecMedia/gfapigonnect"
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
```

### Use case

We use this package against [http://www.bullion-investor.com/blog/](http://www.bullion-investor.com/blog/) to visualize survey results in highcharts on our [gold price comparison website](http://www.bullion-investor.com/).