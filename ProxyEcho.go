package main

import(

  "./web"
  //"database/sql"
  //_ "./go-sqlite3"
  atlantis "./atlantis/types"
  //"os"
  //"strconv"
  "fmt"
  "log"
)

var(
   listenAddr string
   //db driver.Conn
)

func init() {
  cfg, err := atlantis.LoadAppConfig()
  if err != nil {
     log.Printf("error opening using default port")
     listenAddr = ":9999"
     return
  }
  listenAddr = fmt.Sprintf(":%d", cfg.HTTPPort)
  //db, err := sql.Open("sqlite3", "./local.db")
  //if err != nil {
   //  log.Fatal(err)
 // }
  //defer db.Close()


  
} 

func hello(val string) string{
   return "hello " + val
}
func healthz(ctx *web.Context, val string){
   ctx.ContentType("text/plain")
   ctx.ResponseWriter.Header().Add("Server-Status", "OK")
   ctx.ResponseWriter.Write([]byte("OK\n"))
}

func main(){
   web.Get("/(healthz)", healthz)
   web.Get("/process/(.*)", hello)
   web.Run("0.0.0.0" + listenAddr)
}
