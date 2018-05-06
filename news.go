package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/xml"
    "golang.org/x/net/html/charset"
    "strings"
)

type Rss struct {
    Channel InformationChannel `xml:"channel"`
}

type InformationChannel struct {
    Title string `xml:"title"`
    Link string `xml:"link"`
    Description string `xml:"description"`
    Items []Item `xml:"item"`
}

type Item struct {
    Title string `xml:"title"`
    Link string `xml:"link"`
    Description string `xml:"description"`
    Category string `category:"category"`
    PubDate string `pubDate:"pubDate"`
}

func (i *Item) BuildLink() string {
    return "<a href=" + i.Link + ">Read More</a>"
}

func index_handler(w http.ResponseWriter, r *http.Request) {
    // r.Header.Add("Content-Type", "text/html; charset=utf-8")
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    resp, _ := http.Get("http://feeds.jn.pt/JN-Ultimas")
    bytes, _ := ioutil.ReadAll(resp.Body)
    var rss Rss
    var body string
    body = string(bytes)
    reader := strings.NewReader(body)
    decoder := xml.NewDecoder(reader)
    decoder.CharsetReader = charset.NewReaderLabel
    must(decoder.Decode(&rss))
    for _, item := range rss.Channel.Items {
        fmt.Fprintf(w, "Item title: %s &rarr; %s<br />", item.Title, item.BuildLink())
    }
}

func main() {
    http.HandleFunc("/", index_handler)
    http.ListenAndServe(":8000", nil)
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}  