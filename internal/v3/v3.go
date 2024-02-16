package v3

import (
	"io"
	"net/http"

	//"io/ioutil"
	"encoding/json"
	"fmt"
	"net/url"
)

func Handler(method string, userAgent string, w http.ResponseWriter, r *http.Request, headers http.Header) {
	xBareUrl := headers.Get("X-Bare-Url")
	xBareHeaders := headers.Get("X-Bare-headers")

	// Parse original URL
	originalURL, err := url.Parse(xBareUrl)
	if err != nil {
		fmt.Println("Invalid URL Ignoring", xBareUrl)
		return
	}

	if headers.Get("Upgrade") == "websocket" {
		fmt.Println("Websocket Detected Performing Handshake")
		HandleWebsocket(w, r)
	}
	// Create a new request
	req, err := http.NewRequest(method, originalURL.String(), r.Body)
	if err != nil {
		fmt.Println(originalURL.String())
		panic(err)
	}

	// Set headers from the JSON string
	var Dheaders map[string]interface{}
	err = json.Unmarshal([]byte(xBareHeaders), &Dheaders)
	if err != nil {
		fmt.Println("Error parsing headers:", err)
	}

	for k, v := range Dheaders {
		req.Header.Set(k, v.(string))
	}

	// Forward cookies from the original request
	for _, cookie := range r.Cookies() {
		req.AddCookie(cookie)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	BareResponseHeaders := make(http.Header)

	for key, values := range resp.Header {
		for _, value := range values {
			BareResponseHeaders[key] = []string{value}
		}
	}
	headersJSON, err := json.Marshal(BareResponseHeaders)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Oops something went wrong")
	}
	if err != nil {
		fmt.Println("Failed to parse")
	}
	if resp.Header.Get("X-Bare-Pass-Headers") != "" {
		fmt.Println("Passing Headers")
		fmt.Println(resp.Header.Get("X-Bare-Pass-Headers"))

	}

	w.Header().Set("X-Bare-Status", fmt.Sprintf("%d", resp.StatusCode))
	w.Header().Set("X-Bare-Status-Text", http.StatusText(resp.StatusCode))
	w.Header().Set("X-Bare-Headers", string(headersJSON))
	w.WriteHeader(resp.StatusCode)
	// Copy response body

	// fmt.Println(resp.Body)
	io.Copy(w, resp.Body)
}
