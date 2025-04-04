package helpers

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func FileServer(port int, ip net.IP, path string) {
	log.SetFlags(log.Ldate | log.Ltime)
	fileServer := http.FileServer(http.Dir(path))
	http.Handle("/", fileServer)
	fmt.Printf("Server starting from %s and will be available via http://%v:%d\n", path, ip.String(), port)
	err := http.ListenAndServe(fmt.Sprintf("%v:%d", ip, port), logRequest(http.DefaultServeMux))
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s %s\n", r.RemoteAddr, r.Method, r.URL, r.UserAgent())
		handler.ServeHTTP(w, r)
	})
}
