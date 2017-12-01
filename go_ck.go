package main

import ( "fmt"
"net/http"
"io/ioutil"
"encoding/json"
)

type result struct {
   Currency []Currency `json:"result"`
}

type Currency struct {
    Cur string `json:"Currency"`

}

func (l Currency) String() string {
  return fmt.Sprintf(l.Cur)

}

//https://bittrex.com/api/v1.1/public/getmarkethistory?market=BTC-DOGE
func main(){
  resp, _ := http.Get("https://bittrex.com/api/v1.1/public/getcurrencies")
  body, _ := ioutil.ReadAll(resp.Body)
//  string_body := string(bytes)
//  fmt.Println(string_body)
  resp.Body.Close()

  s := result{}
  json.Unmarshal(body, &s)
  fmt.Println(s.Currency)
  var url_cur string = "https://bittrex.com/api/v1.1/public/getmarkethistory?market=BTC-"
  url_cur = url_cur + s.Currency[0].Cur

  url, _ := http.Get(url_cur)
  body, _ := ioutil.ReadAll(url.Body)
  json.Unmarshal(body, &bit)
  fmt.Println(bit.Quantity)
  url.Body.Close()
/*  for index := 0; index < len(s.Currency); index++ {
    var url_cur string = "https://bittrex.com/api/v1.1/public/getmarkethistory?market=BTC-"
    url_cur = url_cur + s.Currency[index].Cur
    url, _ := http.Get(url_cur)
    body, _ := ioutil.ReadAll(url.Body)
    json.Unmarshal(body, &bit)
    fmt.Println(bit.Quantity)
    url.Body.Close()
  }
*/
  //fmt.Println(s.Currency)

}
