package v3

import (
    "net/http"
    "fmt"
)

func SplitHeaders(headers http.Header) string {
    fmt.Println(headers)
    var bareHeaders string
    for k, v := range headers {
        if k[:2] == "X-" {
            bareHeaders += fmt.Sprintf("%s: %s\r\n", k, v[0])
            delete(headers, k)
        }
    }
    fmt.Println(bareHeaders)
    return bareHeaders
}
