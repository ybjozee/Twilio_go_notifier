package handler

import (
	"app/notifier"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	clearContext()
	context.Data = notifier.GetSupportedMedia()
	render(w, "index")
}

func notify(w http.ResponseWriter, r *http.Request) {
	clearContext()

	r.ParseForm()
	medium := r.FormValue("medium")
	recipient := r.FormValue("recipient")
	message := r.FormValue("message")

	if err := validatePhoneNumber(recipient); err != nil {
		renderError(err, w, "index")
		return
	}

	notifier, err := notifier.GetNotifier(medium)

	if err != nil {
		renderError(err, w, "index")
		return
	}

	response, err := notifier.Notify(recipient, message)

	if err != nil {
		renderError(err, w, "index")
		return
	}

	context.Message = fmt.Sprintf("Message delivered successfully with SID: %s", response)

	render(w, "index")
}

func GetRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/notify", notify)
	return mux
}
