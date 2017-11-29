package main

import ( "fmt"
"net/http"
"io/ioutil"
"encoding/json"
//"reflect"
//"strings"
//"strconv"
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
fmt.Println(len(s.Currency))
  for i:=0 ; i< len(s.Currency); i++{
    //fmt.Println(s.Currency[i])

strng := "https://bittrex.com/api/v1.1/public/getmarketsummary?market=BTC-"
st:= s.Currency[i].Cur
url_c:= strng+st

      murat, _ := http.Get(url_c)
      bytes, _ := ioutil.ReadAll(murat.Body)
      string_body := string(bytes)
      fmt.Println(string_body)
    resp.Body.Close()

      }
//  fmt.Println(s.Currency[i])


}
