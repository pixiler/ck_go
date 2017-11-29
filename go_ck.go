package main

import ( "fmt"
"net/http"
"io/ioutil"
"encoding/json"
"strings"
)

type result struct {
   Currency []Currency `json:"result"`
   Quantity []Quantity `json:"Quantity"`
}

type Currency struct {
    Cur string `json:"Currency"`
}

type Quantity struct {
    Qua string `json:"Quantity"`
}

func (l Currency) String() string {
  return fmt.Sprintf(l.Cur)

}
func (k Quantity) String() string {
  return fmt.Sprintf(k.Qua)
}

func main(){
  resp, _ := http.Get("https://bittrex.com/api/v1.1/public/getcurrencies")
  body, _ := ioutil.ReadAll(resp.Body)
  //string_body := string(bytes)
  //fmt.Println(string_body)
  resp.Body.Close()

  s := result{}
  bit := result{}
  json.Unmarshal(body, &s)

  for index := 0; index < len(s.Currency); index++ {
    var url_cur string  = "https://bittrex.com/api/v1.1/public/getmarkethistory?market=BTC-"
    //url_cur += s.Currency[index]
    //url_cur = strings.Contains(url_cur,s.Currency[index])
    string_body := string(s.Currency)
    url, _ := http.Get(strings.Contains(url_cur,string_body))
    body, _ := ioutil.ReadAll(url.Body)
    json.Unmarshal(body, &bit)
    fmt.Println(bit.Quantity)
  }

  //fmt.Println(s.Currency)

}
