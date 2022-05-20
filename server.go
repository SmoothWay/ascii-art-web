package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const PORT = "8080"

var FONTS = []string{"standard", "shadow", "thinkertoy"}

func main() {
	fileServer := http.FileServer(http.Dir("./templates")) // openning templates folder and extracts html file
	http.Handle("/", fileServer)                           // handling main page to show index.html
	http.HandleFunc("/ascii-art", formHandler)             // Execute formhandler when submit button pressed
	fmt.Printf("Starting server at port %v\n", PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil { // start the server
		log.Fatal(err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		log.Printf("Status: 405. %v Method is Not Allowed", r.Method)
		return
	}

	texts, textOk := r.Form["text"]
	fonts, fontOk := r.Form["fontType"]
	if !textOk || !fontOk || !contains(FONTS, r.Form["fontType"]) || len(texts) != 1 || len(fonts) != 1 {
		w.WriteHeader(422)
		log.Println("Status: Unprocessable Entity (422)")
		return
	}
	// if _, ok := r.Form["text"]; !ok {
	// 	log.Println("Status: Unprocessable Entity (422)")
	// 	w.WriteHeader(422)
	// 	return
	// }
	// if _, ok := r.Form["fontType"]; !ok {
	// 	log.Println("Status: Unprocessable Entity (422)")
	// 	w.WriteHeader(422)
	// 	return
	// }
	// // }
	// if !contains(FONTS, r.Form["fontType"]) {
	// 	w.WriteHeader(422)
	// 	log.Println("Status: Unprocessable Entity (422)")
	// 	return
	// }
	// if err != nil {
	// 	w.WriteHeader(422)
	// 	log.Printf("Status: 422 Unprocessable Entity.")
	// 	return
	// }
	// log.Printf("body: %v", len(body))
	text := texts[0] // takes value from textfield
	font := fonts[0] // takes value from radio-button
	if font == "" {
		w.WriteHeader(400)
		log.Println("Status: Bad Request (400)")
		fmt.Fprintf(w, "400 Bad Request.")
		return
	}
	fontType := "fontstyles/" + font + ".txt" // path to .txt
	arg := strings.ReplaceAll(text, "\n", "\\n")
	if isNotAscii(arg) {
		w.WriteHeader(400)
		log.Println("Status: Bad Request (400). Entered not ASCII symbol")
		fmt.Fprintf(w, "400 Bad Request.\nEntered not supported symbol")
		return
	}
	filePath, err := ioutil.ReadFile(fontType)
	if err != nil {
		w.WriteHeader(500)
		log.Println("Status: Internal Server Error (500). ") // if file doesn't exist
		fmt.Fprintf(w, "500 Internal Server Error.\n")
		return
	}

	// TODO HashSum
	if len(filePath) != 6623 && font == "standard" || len(filePath) != 7463 && font == "shadow" || len(filePath) != 4703 && font == "thinkertoy" || err != nil {
		w.WriteHeader(500)
		log.Println("Status: Internal Server Error (500). ") // if file differs from original one
		fmt.Fprintf(w, "500 Internal Server Error.\n")
		return
	}
	if arg == "" {
		return
	}
	if isEmpty(arg) {
		count := len(arg) / 2
		for i := 0; i < count; i++ {
			fmt.Fprintf(w, "\n")
		}
		return
	} else {
		// w.WriteHeader(200)
		log.Println("Status: OK (200)")
		symbol := getMap(string(filePath))
		t := strings.Split(arg, "\\n")
		for _, val := range t {
			fmt.Fprintf(w, getStr(val, symbol)) // prints ascii-art depending on fontstyle
		}

	}
}
