package main

import (
    "fmt"
    "net/http"
    "net/textproto"
)

func Handler(w http.ResponseWriter, r *http.Request) {

    // recommend
    v1 := r.Header.Get("OPFROMFIRST")

    // need to convert string cases manually
    key := textproto.CanonicalMIMEHeaderKey("HEADER_KEY_2")
    // v2 type is []string
    v2, _ := r.Header[key]

    w.Write([]byte(fmt.Sprintf("v1: %s, v2: %v", v1, v2)))
}
