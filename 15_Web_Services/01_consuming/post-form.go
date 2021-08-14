package main
import (
   "crypto/tls"
   "fmt"
   "io/ioutil"
   "net/http"
   "net/http/cookiejar"
   "net/url"
   "strings"
)
func init() {
   cfg := &tls.Config{
      InsecureSkipVerify: true,
   }
   http.DefaultClient.Transport = &http.Transport{
      TLSClientConfig: cfg,
   }
}
func httpPostForm() {
   parm := url.Values{}
   parm.Add("client_id", "ayeshaj")
   parm.Add("response_type","code")
   parm.Add("scope","public_profile")
   parm.Add("redirect_uri","http://ayeshaj:8080/playground")
   cookieJar, _ := cookiejar.New(nil)
   tr := &http.Transport{
      TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
   }
   client := &http.Client{Transport: tr,
      Jar: cookieJar,
      CheckRedirect: func(req *http.Request, via []*http.Request) error {
         return http.ErrUseLastResponse
      }}
   fmt.Println(strings.NewReader(parm.Encode()))
   req, err := http.NewRequest("POST", "https://ayeshaj:9000/register", strings.NewReader(parm.Encode()))
   req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36")
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   resp, err := client.Do(req)
   if err != nil {
      fmt.Printf("%s", err)
   }
   fmt.Println(resp.Header)
   defer resp.Body.Close()
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      fmt.Printf("%s", err)
   }
   fmt.Println(string(body))
}
func main() {
   httpPostForm();
}