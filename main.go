package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kandros/visits/internal/visit"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := visit.NewFromRequest(r)

		if err := v.Persist(); err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("ok"))

	})

	addr := ":" + os.Getenv("PORT")

	log.Fatal(http.ListenAndServe(addr, handler))
}
