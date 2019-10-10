package telzarsms

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Callback execute on getting Push SMS back from Telzar
func Callback(h HTTPHandler, path string, mux *http.ServeMux,
	onCallback func(http.ResponseWriter, *http.Request, *DLRPush, error)) {

	handleCallback := func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if !(contentType == "application/x-www-form-urlencoded" && r.ContentLength > 0) {
			// Not a valid request, but do what ever you want from it
			if onCallback != nil {
				onCallback(w, r, nil, fmt.Errorf(
					"Content-Type (%s) or ContentLength (%d) are invalid",
					contentType, r.ContentLength,
				))
			}
			return
		}
		_ = r.ParseForm()
		if strings.Contains(os.Getenv("SMSHTTPDEBUG"), "dump=true") {
			fmt.Printf(">>>> Request dump: %+v\n", r)
			for k, v := range r.Form {
				fmt.Printf("\t %s=%v\n", k, v)
			}
		}

		dlrInfo := URLToDLRPush(r.Form)

		if onCallback != nil {
			onCallback(w, r, dlrInfo, nil)
		}
	}

	h.OnGettingSMS(path, mux, handleCallback)
}
