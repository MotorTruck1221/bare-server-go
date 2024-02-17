package utils

import (
    "runtime"
    "fmt"
    "encoding/json"
)

type Maintainer struct {
    Email string `json:"email"`
    website string
}

type Project struct {
    Name string `json:"name"`
    Description string `json:"description"`
    Email string `json:"email"`
    Website string `json:"website"`
    Repository string `json:"repository"`
    Version string `json:"version"`
}


type Response struct {
    Versions []string `json:"versions"`
    Language string `json:"language"`
    MemoryUsage uint64 `json:"memoryUsage"`
    Maintainer Maintainer `json:"maintainer"`
    Project Project `json:"project"`
}

func calcMem() uint64 {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    return m.Alloc / 1024 / 1024
}

func GetJson() string {
    maintainer := Maintainer{"motortruck1221@protonmail.com", "https://motortruck1221.is-a.dev"}
    project := Project{"Bare Server Go", "A TOMPHTTP complaint server written in Go", "", "", "https://github.com/tomphttp/bare-server-go", "v0.0.2"}
    response := Response{[]string{"v3"}, "Go", calcMem(), maintainer, project}
    //pretty print json 
    json, err := json.MarshalIndent(response, "", "  ")
    if err != nil { fmt.Println("Error in making json: ", err) }
    return string(json)
}
