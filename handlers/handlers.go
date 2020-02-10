package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/sokks/url-short/hasher"
	"github.com/sokks/url-short/storage"
)

var (
	maxRetries = 2
	baseURL    = ""
)

func Init(newBaseURL string, maxRehash int) {
	baseURL = newBaseURL
	maxRetries = maxRehash
}

// IndexHandler handles url redirection
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	log.Printf("[REQUEST] get hash=%s", hash)

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

type NewResponse struct {
	Long  string `json:"long"`
	Short string `json:"short"`
}

// NewHandler for short url creation
func NewHandler(w http.ResponseWriter, r *http.Request) {
	longurl := r.FormValue("url")
	log.Printf("[REQUEST] /new url=%s", longurl)
	longurl = strings.Trim(longurl, " ")
	if longurl == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !strings.HasPrefix(longurl, "http://") && !strings.HasPrefix(longurl, "https://") {
		longurl = "http://" + longurl
	}

	hash := hasher.HashURL(longurl, 0)

	var err error
	for i := 0; i < maxRetries; i++ {
		err = storage.Put(hash, longurl)
		if err == nil {
			break
		}
		if err == storage.ErrKeyAlreadyExists {
			var existingURL string
			existingURL, err = storage.Get(hash)
			if err == nil && existingURL == longurl {
				break
			}

			log.Printf("[COLLISION] url=%s hash=%s", longurl, hash)
			hash = hasher.HashURL(longurl, uint64(i+1))
		} else {
			log.Printf("[SET error] %s", err)
		}
	}
	if err != nil {
		log.Printf("[ERROR] cannot save url hash after %d tries url=%s hash=%s err=%s", maxRetries, longurl, hash, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := NewResponse{
		Long:  longurl,
		Short: baseURL + hash,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
