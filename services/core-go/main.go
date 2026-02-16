
package main
import ("encoding/json"; "log"; "net/http")
type health struct{ Status string `json:"status"` }
func main(){
  http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type","application/json")
    data, err := json.Marshal(health{Status:"ok"})
    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      log.Printf("Error marshaling health response: %v", err)
      return
    }
    if _, err := w.Write(data); err != nil {
      log.Printf("Error writing response: %v", err)
      return
    }
  })
  log.Println("core-go listening on :8080")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Printf("Server error: %v", err)
    return
  }
}
