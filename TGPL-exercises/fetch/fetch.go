// Fetch prints the contents found at a URL.
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
        validateProtocolPrefix(&url)
        resp, err := http.Get(url)
        if err != nil {
            _, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        fmt.Printf("STATUS: %s\n", resp.Status)
        _, err = io.Copy(os.Stdout, resp.Body)
        if err != nil {
            _, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            _ = resp.Body.Close()
            os.Exit(1)
        }
        _ = resp.Body.Close()
    }
}

// Adds https prefix if a URL is missing it. Modifies http to https.
func validateProtocolPrefix(url *string) {
    if !strings.HasPrefix(*url, "https://") {
        if strings.HasPrefix(*url, "http://") {
            *url = "https://" + (*url)[7:]
        } else {
            *url = "https://" + *url
        }
    }
}
