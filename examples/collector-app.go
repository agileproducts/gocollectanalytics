/*
An example web app that collects using the gocollectanalytics package
and then just writes it to log.
*/

package main

import (
	"fmt"
	"github.com/agileproducts/gocollectanalytics"
	"github.com/agileproducts/gocollectanalytics/keenio"
	"github.com/codegangsta/negroni"
	"log"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()

	logStore := gocollectanalytics.NewLogstore()
	logCollector := gocollectanalytics.NewCollector(logStore)

	kc := keenio.KeenIOConfig{WriteKey: os.Getenv("KEENIO_WRITE_KEY"), ProjectID: os.Getenv("KEENIO_PROJECT_ID")}
	ks, err := keenio.NewKeenIOStore(kc)
	if err != nil {
		log.Fatal(err)
	}
	keenCollector := gocollectanalytics.NewCollector(ks)

	mux.HandleFunc("/", SayHello)
	mux.HandleFunc("/collecttolog", logCollector.CollectData)
	mux.HandleFunc("/collecttokeenio", keenCollector.CollectData)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")

}

// SayHello is a http.HandlerFunc to return Hello, world given any request
func SayHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello, world")
}
