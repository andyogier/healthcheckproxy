package healthcheck

import (
	"io"
	"log"
	"net/http"
)

func healthcheck(port string, appName string, query string) {
	http.HandleFunc("/health", customHandler(appName, query))
	http.ListenAndServe(port, nil)
}

func customHandler(appName string, query string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		response, err := http.Get("http://healthcheck-svc.healthcheck.svc.cluster.local/" + appName + "?" + query)
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
