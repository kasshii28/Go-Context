package overview

import (
	"net/http"
	"log"
)

func main() {
	// ハンドラの定義
	h1 := func(w http.ResponseWriter, _*http.Request){
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _*http.Request){
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint",h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}