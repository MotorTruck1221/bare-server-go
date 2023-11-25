package utils

import (
    "runtime"
    "fmt"
)

type Maintainer struct {
    email string
    website string
}

type Project struct {
    name string
    description string
    email string
    website string
    repository string
    version string
}

type Versions struct {
    versions []string
}

type Response struct {
    versions Versions
    language string
    memoryUsage uint64
    maintainer Maintainer
    project Project
}

func calcMem() uint64 {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    return m.Alloc / 1024 / 1024
}

func GetJson() string {
    jsonResp := `{
        "versions": ["v3"],
        "language": "Go",
        "memoryUsage": ` + fmt.Sprint(calcMem()) + `,
        "project": {
            "name": "bare-go",
            "description": "A Bare Server in GoLang",
            "email": "support@rubynetwork.tech",
            "website": "https://rubynetwork.tech",
            "repository": "https://github.com/ruby-network/bare-go",
            "version": "v0.0.1"
        }
    }`
    return jsonResp
}
