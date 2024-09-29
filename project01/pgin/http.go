package pgin

import (
	"fmt"
	"net/http"
)
func HttpTT(){
	
     http.HandleFunc("/go_api", func(w http.ResponseWriter, s *http.Request){
		if s.Method == "GET"{
			fmt.Fprintf(w, "Hello World request: %s", s.URL.Path)
		}else if s.Method == "POST"{
			fmt.Fprintf(w, "error")
		}
	 })
    
	 http.ListenAndServe(":8080", nil)

}