package v3

import ( 
    "net/http"
    "io/ioutil"
    "encoding/json"
    "net/url"
    "fmt"
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
    client := &http.Client{}
    request, err := http.NewRequest(method, xBareUrl, nil)
    if err != nil { panic(err) }
    request.Header.Set("User-Agent", userAgent)
    response, err := client.Do(request)
    if err != nil { panic(err) }
    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    if err != nil { panic(err) }
    bareHeaders, err := json.Marshal(response.Header)
    if err != nil { panic(err) }
    // detect an upgrade to websocket
    w.Header().Set("X-Bare-Status", fmt.Sprintf("%d", response.StatusCode))
    w.Header().Set("X-Bare-Status-Text", response.Status)
    w.Header().Set("X-Bare-Headers", string(bareHeaders))
    w.Write(body)
}
