package main

import (
    "bytes"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

var url = "http://p.iotek.org"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func exists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }
    if os.IsNotExist(err) {
        return false
    }
    return false
}

func prepare(fp string) *http.Request {
    f, err := ioutil.ReadFile(fp)
    check(err)

    var str = []byte(string(f))
    req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(str))
    req.Header.Set("Content-Type", "text/html; charset=utf-8")

    return req
}

func upload(req *http.Request) string {
    client := &http.Client{}

    resp, err := client.Do(req)
    check(err)

    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    return string(body)
}

func main() {
    flag.Parse()

    if len(flag.Args()) == 0 {
        fmt.Println("usage: pup [files]")
        os.Exit(1)
    }

    for _, p := range flag.Args() {
        if exists(p) {
            r := prepare(p)
            fmt.Print(upload(r))
        } else {
            log.Fatal("file doesn't exist: ", p)
        }
    }
}
