package v3

import ( 
    "net/http"
    "io/ioutil"
    "fmt"
)

func Handler(method string, userAgent string, w http.ResponseWriter, r *http.Request, headers http.Header) {
    xBareUrl := headers.Get("X-Bare-Url")
    client := &http.Client{}
    request, err := http.NewRequest(method, xBareUrl, nil)
    if err != nil { panic(err) }
    request.Header.Set("User-Agent", userAgent)
    response, err := client.Do(request)
    if err != nil { panic(err) }
    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    if err != nil { panic(err) }
    w.Header().Add("Content-Encoding", response.Header.Get("Content-Encoding"))
    w.Header().Add("Content-Type", response.Header.Get("Content-Type"))
    w.Header().Add("Content-Length", response.Header.Get("Content-Length"))
    w.Header().Add("X-Bare-Status", "200")
    w.Header().Add("X-Bare-Status-text", "Ok")
    w.Write(body)
    fmt.Println(w.Header())
}
