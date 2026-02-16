package main
import ("encoding/json"; "fmt"; "os"; "net/http")
type health struct{ Status string `json:"status"` }

func healthHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","application/json")
  data, err := json.Marshal(health{Status:"ok"})
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(os.Stderr, "Error marshaling health response: %v\n", err)
    return
  }
  if _, err := w.Write(data); err != nil {
    fmt.Fprintf(os.Stderr, "Error writing response: %v\n", err)
    return
  }
}

func main(){
  mux := http.NewServeMux()
  mux.HandleFunc("/healthz", healthHandler)
  
  fmt.Fprintln(os.Stdout, "core-go listening on :8080")
  if err := http.ListenAndServe(":8080", mux); err != nil {
    fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
    os.Exit(1)
  }
}
