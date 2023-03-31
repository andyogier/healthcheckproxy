package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	var path = "/health"
	var appName = "appNameFromDeploymentYaml"
	var query = "percentageHealthyPods=80" // Can use percentageHealthyPods or min
	var namespace = "myAppsNamespace"
	var port = ":80"

	http.HandleFunc(path, CustomHandler(appName, query, namespace))

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func CustomHandler(appName string, query string, namespace string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		response, err := http.Get("http://healthcheck-svc.healthcheck.svc.cluster.local/chkapp/" + appName + "?" + query + "&namespace=" + namespace)
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()

		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(response.StatusCode)
		w.Write(bodyBytes)
	}
}
