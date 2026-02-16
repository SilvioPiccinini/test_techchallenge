
package main
import ("encoding/json"; "log"; "net/http")
type health struct{ Status string `json:"status"` }
func main(){
  http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type","application/json")
    json.NewEncoder(w).Encode(health{Status:"ok"})
  })
  log.Println("core-go listening on :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
