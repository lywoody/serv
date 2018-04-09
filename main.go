package main

import "net/http"

func httpserv() {

	h := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("baa!!"))
	}
	http.HandleFunc("/", h)
	http.ListenAndServe(":8080", nil)
}


func main(){
	httpserv()
}
