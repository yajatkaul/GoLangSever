package main

import (
	"fmt"
	"io"
	"net/http"
)

//send response
func helloWorldPage(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path{
	case "/":
		fmt.Fprint(w, "Hello World")
	case "/ninja":
		fmt.Fprint(w, "Wallace")
	default:
		fmt.Fprint(w, "Hello")
	}
	fmt.Fprint(w, "Hello World")
}

//set headers
func htmlVsPlain(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

//method check
func handlePost(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
}

//reading the recieved post request
func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Read the request body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }

	//For clean up no matter how it returns better efficiency 
    defer r.Body.Close()

    // Print the request body
    fmt.Println("Request Body:", string(body))

    // Respond to the client
    w.Header().Set("Content-Type", "text/plain")
    _, err = w.Write([]byte("Hello, World!")) //This is the response sent
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
        return
	}
}



//start server
func main(){
	http.HandleFunc("/", htmlVsPlain)
	server := http.Server{
		Addr: ":5000",
		Handler: nil,
		ReadTimeout: 1000,
		WriteTimeout: 1000,
	}
	var muxNinjaMode http.ServeMux  //custom server mux
	server.Handler = &muxNinjaMode
	muxNinjaMode.HandleFunc("/",helloWorldPage)

	server.ListenAndServe()
}