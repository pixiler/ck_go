package main

import ( "fmt"
"net/http"
"io/ioutil"
//"encoding/xml"
)

type SitemapIndex struct {
  Locations []Locations `xml:"sitemap"`
}

type Locations struct {
    Loc string `xml:"loc"`
}

func (l Locations) String() string {
  return fmt.Sprintf(l.Loc)

}

func main(){
  resp, _ := http.Get("https://bittrex.com/home/markets")
  bytes, _ := ioutil.ReadAll(resp.Body)
  string_body := string(bytes)
  fmt.Println(string_body)
  resp.Body.Close()

  //var s SitemapIndex
  //xml.Unmarshal(bytes, &s)
  //fmt.Println(s.Locations)

}
