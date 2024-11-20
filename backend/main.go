package main

import (
	"encoding/json"
	"fmt"
	"strings"

	// "io"
	"log"
	// "net"
	"net/http"
	// "net/url"
	// "strings"
)

func main() {
	http.HandleFunc("/dog", handleDogRequest)
	// http.HandleFunc("/", handleRequest)
	log.Println("Starting proxy server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleDogRequest(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	breed := req.URL.Query().Get("breed")
	if breed == "" {
		http.Error(w, "Breed not specified", http.StatusBadRequest)
		return
	}

	apiBreed := formatBreedName(breed)

	imageUrl, err := fetchDogImage(apiBreed)
	if err != nil {
		http.Error(w, "Error fetching dog image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"imageUrl": imageUrl,
	})

}
func formatBreedName(breed string) string {
	breedMap := map[string]string{
		"Weiner Dog":          "dachshund",
		"Australian Shepherd": "australian/shepherd",
		"Pitbull":             "pitbull",
		"Golden Retriever":    "retriever/golden",
	}

	if apiBreed, ok := breedMap[breed]; ok {
		return apiBreed
	}

	breed = strings.ToLower(breed)
	breed = strings.ReplaceAll(breed, " ", "-")
	return breed
}
func fetchDogImage(breed string) (string, error) {
	apiUrl := "https://dog.ceo/api/breed/" + breed + "/images/random"
	fmt.Println(apiUrl)
	resp, err := http.Get(apiUrl)
	fmt.Println(resp)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}
	var result struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	if result.Status != "success" {
		return "", err
	}
	return result.Message, nil

}

// func handleRequest(w http.ResponseWriter, req *http.Request) {
// 	if req.Method == http.MethodConnect {
// 		handleHTTPS(w, req)
// 	} else {
// 		handleHTTP(w, req)
// 	}
// }
// func handleHTTPS(w http.ResponseWriter, req *http.Request) {
// 	log.Printf("HTTPS Request: %s", req.Host)

// 	destConn, err := net.Dial("tcp", req.Host)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusServiceUnavailable)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	hijacker, ok := w.(http.Hijacker)
// 	if !ok {
// 		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
// 		return
// 	}
// 	clientConn, _, err := hijacker.Hijack()

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusServiceUnavailable)
// 		return
// 	}
// 	go transfer(destConn, clientConn)
// 	go transfer(clientConn, destConn)
// }

// func transfer(destination net.Conn, source net.Conn) {
// 	defer destination.Close()
// 	defer source.Close()
// 	io.Copy(destination, source)
// }

// func handleHTTP(w http.ResponseWriter, req *http.Request) {
// 	log.Printf("HTTP Request: %s %s", req.Method, req.URL.String())
// 	url, err := url.Parse(req.URL.String())
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 	}

// 	proxyReq, err := http.NewRequest(req.Method, url.String(), req.Body)
// 	if err != nil {
// 		http.Error(w, "Error creating Proxy request", http.StatusInternalServerError)
// 	}
// 	proxyReq.Header = req.Header
// 	proxyReq.Header.Add("X-Proxy-Server", "GoProxy")

// 	client := http.Client{}
// 	resp, err := client.Do(proxyReq)
// 	if err != nil {
// 		http.Error(w, "Error forwarding request", http.StatusInternalServerError)
// 	}
// 	defer resp.Body.Close()

// 	for key, values := range resp.Header {
// 		for _, value := range values {
// 			w.Header().Add(key, value)
// 		}
// 	}
// 	w.WriteHeader(resp.StatusCode)
// 	io.Copy(w, resp.Body)

// 	bodyBytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		http.Error(w, "Error reading response body", http.StatusInternalServerError)
// 		return
// 	}
// 	resp.Body.Close()

// 	bodyString := string(bodyBytes)
// 	bodyString = strings.ReplaceAll(bodyString, "Hello", "Hi")
// 	w.Write([]byte(bodyString))

// }
