package v3

import ( 
    "net/http"
    //"io/ioutil"
    "encoding/json"
    "net/url"
    "fmt"
    "github.com/go-resty/resty/v2"
)

func Handler(method string, userAgent string, w http.ResponseWriter, r *http.Request, headers http.Header) {
    xBareUrl := headers.Get("X-Bare-Url")
    //detect if websocket
    if headers.Get("Upgrade") == "websocket" {
        fmt.Println("Websocket Detected Performing Handshake")
        HandleWebsocket(w, r)
    }
    _, err := url.ParseRequestURI(xBareUrl)
    if err != nil { 
        fmt.Println("Invalid URL Ignoring", xBareUrl)
        return
    }
    client := resty.New()
    //client.SetHeaders(headers)
    client.SetHeader("User-Agent", userAgent)
    fmt.Println("Method", method)
    response, err := client.R().
    //create request based on method and x-bare-url 
        SetResult(&resty.Response{}). // or SetResult(AuthSuccess{}).
        Execute(method, xBareUrl)
    
    if err != nil {
        fmt.Println(err)
    }

    jsonHeaders, err := json.Marshal(response.Header())

    //if x-bare-pass-headers is set add it to jsonHeaders
    if headers.Get("X-Bare-Pass-Headers") != "" {
        fmt.Println("Passing Headers")
        fmt.Println(headers.Get("X-Bare-Pass-Headers"))
        fmt.Println(response.Header())
        fmt.Println(jsonHeaders)
        jsonHeaders = append(jsonHeaders, []byte(headers.Get("X-Bare-Pass-Headers"))...)
    }

    w.Header().Set("X-Bare-Status", fmt.Sprintf("%d", response.StatusCode()))
    w.Header().Set("X-Bare-Status-Text", response.Status())
    w.Header().Set("X-Bare-Headers", string(jsonHeaders))
    w.Write(response.Body())
}
