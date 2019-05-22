package hasher

import (
	"crypto/sha1"
	"log"
	"strconv"
	"sync"

	"github.com/lytics/base62"
)

var (
	counter uint64
	lock    = &sync.Mutex{}
)

// todo counter to db

func HashURL(url string) string {
	lock.Lock()
	salt := counter
	counter++
	lock.Unlock()

	data := []byte(url + strconv.FormatUint(salt, 10))
	log.Printf("data: %+v", data)
	sha1sum := sha1.Sum(data)
	log.Printf("sha1sum: %+v", sha1sum)
	// start := []byte(sha1sum[0:6])
	var hash = make([]byte, 100)
	base62.StdEncoding.Encode(hash, []byte(sha1sum[:]))
	log.Printf("hash: %+v", hash)
	if len(hash) > 6 {
		hash = hash[:6]
	}

	return string(hash)
}
