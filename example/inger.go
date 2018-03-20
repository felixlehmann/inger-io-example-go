package main

// Author
//     Felix Lehmann
// Examples
//     go run inger.go
//     go run inger.go v1 google adwords v201710 deprecated

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
)

func main() {
    base_url := "http://inger.io/"
    inger_version := "v1"
    vendor := "facebook"
    service := "graph"
    version := "v2.10"
    method := "deprecated"

    if (len(os.Args) == 6) {
        inger_version = os.Args[1]
        vendor = os.Args[2]
        service = os.Args[3]
        version = os.Args[4]
        method = os.Args[5]
    }

    full_url := fmt.Sprintf("%s/%s/%s/%s/%s/%s", base_url, url.QueryEscape(inger_version), url.QueryEscape(vendor), url.QueryEscape(service), url.QueryEscape(version), url.QueryEscape(method))

    response, err := http.Get(full_url)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Printf("%s", string(data))
    }
}
