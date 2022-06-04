package handlers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	FONTS               = []string{"standard", "shadow", "thinkertoy"}
	Templates, TemplErr = template.ParseGlob("ui/templates/*.html")
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		log.Println("Status: Not Found (404)")
		Templates.ExecuteTemplate(w, "error.html", 404)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		log.Printf("Status: 405. %v Method is Not Allowed", r.Method)
		Templates.ExecuteTemplate(w, "error.html", http.StatusMethodNotAllowed)
		return
	}

	Templates.ExecuteTemplate(w, "index.html", nil)
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		log.Printf("Status: 405. %v Method is Not Allowed", r.Method)
		Templates.ExecuteTemplate(w, "error.html", http.StatusMethodNotAllowed)
		return
	}
	if TemplErr != nil {
		log.Fatal(TemplErr)
	}
	texts, textOk := r.Form["text"]
	fonts, fontOk := r.Form["fontType"]
	if !textOk || !fontOk || !Contains(FONTS, fonts) || IsNotAscii(texts[0]) {
		w.WriteHeader(400)
		log.Println("Status: Bad Request (400)")
		Templates.ExecuteTemplate(w, "error.html", http.StatusBadRequest)
		return
	}
	text := strings.Join(texts, "") // takes value from textfield
	font := strings.Join(fonts, "") // takes value from radio-button

	fontType := "fontstyles/" + font + ".txt" // path to .txt
	filePath, err := ioutil.ReadFile(fontType)
	if err != nil || !DHashSum(filePath) {
		w.WriteHeader(500)
		log.Println("Status: Internal Server Error (500).")
		Templates.ExecuteTemplate(w, "error.html", http.StatusInternalServerError)
		return
	}
	Templates.ExecuteTemplate(w, "index.html", nil)
	log.Println("Status: OK (200)")
	symbol := GetMap(string(filePath))
	text1 := strings.Split(text, "\n")

	fmt.Fprintf(w, "<div id='maincontent'><pre>")
	for _, val := range text1 {
		fmt.Fprint(w, GetStr(val, symbol)) // prints ascii-art depending on fontstyle
	}
	fmt.Fprintf(w, "</pre></div>")
}
