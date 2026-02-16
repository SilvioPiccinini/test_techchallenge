
package main
import ("encoding/json"; "log"; "net/http")
type health struct{ Status string `json:"status"` }
func main(){
  http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type","application/json")
    if err := json.NewEncoder(w).Encode(health{Status:"ok"}); err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      log.Printf("Error encoding health response: %v", err)
      return
    }
  })
  log.Println("core-go listening on :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
