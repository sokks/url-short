package handlers

import (
	"fmt"
	"log"
	"net/http"

	"gophersises/02_url_shortener/hasher"
	"gophersises/02_url_shortener/storage"

	"github.com/gorilla/mux"
)

var maxRetries = 2

// IndexHandler handles url redirection
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	log.Printf("CHEKING HASH: %s", hash)

	fullURL, err := storage.Get(hash)
	if err == storage.ErrKeyNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fullURL, http.StatusFound)
}

// NewHandler for short url creation
func NewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write(formTmpl)
		return
	}

	longurl := r.FormValue("url")
	fmt.Fprintln(w, "you enter: ", longurl)

	hash := hasher.HashURL(longurl)

	var err error
	for i := 0; i < maxRetries; i++ {
		err = storage.Put(hash, longurl)
		if err == nil {
			break
		}
		if err == storage.ErrKeyAlreadyExists {
			log.Printf("[COLLISION] url:%s hash:%s", longurl, hash)
			hash = hasher.HashURL(longurl)
		} else {
			log.Printf("[SET error] %s", err)
		}
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "hash: ", hash)
}

var formTmpl = []byte(`
<html>
<head profile="http://www.w3.org/2005/10/profile">
<link rel="icon" 
      type="image/png" 
      href="/static/img/favicon.png" />
</head>
	<body>
	<form action="/new" method="post" name="urlform">
		url: <input type="text" name="url">
		<input type="submit" value="Shorten!">
	</form>
	</body>
</html>
`)

// todo big form
