package main

import(
  "net/http"
  "io/ioutil"
  "./web"
  atlantis "./atlantis/types"
  "fmt"
  "log"
)

var(
   echoaddr string
   echoport string
   listenAddr string
)

func init() {

  echoaddr = "127.0.0.1:9998"

  cfg, err := atlantis.LoadAppConfig()
  if err != nil {
     log.Printf("error opening using default port")
     listenAddr = ":9999"
     return
  }
  listenAddr = fmt.Sprintf(":%d", cfg.HTTPPort)
  
  if proxyechoDep := cfg.Dependencies["kyle-echo"]; proxyechoDep != nil {
           if p := proxyechoDep["address"].(string); p != "" {
              echoaddr = p
           }
  }
  log.Printf("proxying to: " + echoaddr);  
   
}
 
func redirect(url string) string{
  client := &http.Client{}
  log.Printf("my echoaddr: " + echoaddr)
  req, err := http.NewRequest("GET", "http://" + echoaddr + "/" + url, nil)

  resp, err := client.Do(req) 
  
  log.Printf("Attempting to get: " + echoaddr + "/" + url)
  if err != nil { 
     log.Printf("Trouble getting url, %s\n", err) 
  }
  defer resp.Body.Close()

  bodybytes, err := ioutil.ReadAll(resp.Body)
  if err != nil { log.Printf("Trouble reading resp body, %s\n", err) }
 
  bodytxt := string(bodybytes)  
  return "Proxied output from echo: " + bodytxt

}
func healthz(ctx *web.Context, val string){
   ctx.ContentType("text/plain")
   ctx.ResponseWriter.Header().Add("Server-Status", "OK")
   ctx.ResponseWriter.Write([]byte("OK\n"))
}

func main(){
   web.Get("/(.*)", redirect)
   web.Run("0.0.0.0" + listenAddr)
  /* dest := echoaddr + ":" + echoport
   src := "0.0.0.0" + listenAddr
   echourl, err := url.Parse(dest)
   if err != nil { log.Printf("could not parse dest url %s\n", err)}
   h := httputil.NewSingleHostReverseProxy(echourl)
   s := &http.Server{
          Addr:	 	src,	
	  Handler:	h,
	  ReadTimeout:	10 * time.Second,
	  WriteTimeout:	10 * time.Second,
 	  MaxHeaderBytes: 1 << 20,
  }
  log.Printf("error: %s\n", s.ListenAndServe())*/		
}
