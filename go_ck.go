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

func main(){
  resp, _ := http.Get("https://bittrex.com/api/v1.1/public/getcurrencies")
  body, _ := ioutil.ReadAll(resp.Body)
  //string_body := string(bytes)
  //fmt.Println(string_body)
  resp.Body.Close()

  s := result{}
  json.Unmarshal(body, &s)
  fmt.Println(s.Currency)

}
