//Need to decide on proper package arrangement later
//For now edit import to proper mux package
package main

import (
	"fmt"
	"net/http"
)

func main(){
	testMux := New()
	testHandler1 := testHandler{}
	testMux.Handle("/test", nil, testHandler1)
	http.ListenAndServe(":8080", testMux)
}

type testHandler struct{
}

func (t testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Test</h1>")
	return
}

