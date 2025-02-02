package main 

import (
  "log"
  "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }
  w.Write([]byte("Hello from SnippetBox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Display a specific snippet"))
}

func snipperCreate(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    w.Header().Set("Allow", http.MethodPost)
//    w.WriteHeader(405)
//    w.Write([]byte("Method not allowed"))
//    http.Error is a helper function that combineds w.Write and w.WriteHeader
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }
  w.Write([]byte("Create a specific snippet"))
}


func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet/view", snippetView)
  mux.HandleFunc("/snippet/create", snipperCreate)

  log.Print("Starting server on :4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
