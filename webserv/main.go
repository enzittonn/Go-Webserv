package main


import (
  "net/http"
  "fmt"
  "time"
  "html/template"
)


type Welcome struct {
  Name string
  Time string
}


func main() {
  welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

  template := template.Must(template.ParseFiles("template/welcome-template.html"))

  http.Handle("/static/",
    http.StripPrefix("/static/",
        http.FileServer(http.Dir("static"))))


  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if name := r.FormValue("name"); name != "" {
      welcome.Name = name;
    }

    if err := template.ExecuteTemplate(w, "welcome-template.html", welcome);
    err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })


  fmt.Println("Listening")
  fmt.Println(http.ListenAndServe(":8080", nil));

}
