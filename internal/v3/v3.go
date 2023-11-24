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
    w.Header().Set("X-Bare-Status", "200")
    w.Header().Set("X-Bare-Status-Text", "Ok")
    w.Header().Set("X-Bare-Headers", string(bareHeaders))
    w.Write(body)
}
