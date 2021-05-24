package main

import (
	"flag"
	"fmt"
	"github.com/yarysh/iban-service/app/iban/validator"
	"net/http"
	"os"
	"strings"
)

func NewIBANService() *Server {
	server := NewServer()
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ibanParam, ok := r.URL.Query()["iban"]
		if !ok || len(ibanParam) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "iban param required"}`))
			return
		}
		iban := ibanParam[0]
		iban = strings.ReplaceAll(iban, " ", "")
		iban = strings.ToUpper(iban)
		valid := validator.Validate(iban)
		w.Write([]byte(fmt.Sprintf(`{"iban": "%s", "valid": %t}`, iban, valid)))
	})
	return server
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", ":8080", "server addr")
	flag.Parse()

	fmt.Printf("Starting server at %s...\n", addr)
	err := http.ListenAndServe(addr, NewIBANService())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
