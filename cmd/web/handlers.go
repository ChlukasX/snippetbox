package main

import (
"fmt"
"html/template"
"log"
"net/http"
"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

  files := []string{
    "ui/html/pages/base.tmpl.html",
    "ui/html/pages/home.tmpl.html",
    "ui/html/pages/nav.tmpl.html",
  }

  ts, err := template.ParseFiles(files...)
  if err != nil {
    log.Print(err.Error())
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }

  err = ts.ExecuteTemplate(w, "base", nil)
  if err != nil {
    log.Print(err.Error())
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }
}

func snippetView(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    http.NotFound(w, r)
    return
  }

  fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
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

