package main

import (
    "net/http"
)



func main() {

	router := NewRouter()

    http.ListenAndServe(":3000", router)

}

