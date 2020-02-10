package hasher

import (
	"crypto/sha1"
	"strconv"

	"github.com/lytics/base62"
)

func HashURL(url string, miniSalt uint64) string {
	data := []byte(url + strconv.FormatUint(miniSalt, 10))
	sha1sum := sha1.Sum(data)
	var hash = make([]byte, 100)
	base62.StdEncoding.Encode(hash, []byte(sha1sum[:]))
	// log.Printf("hash: %+v", hash)
	if len(hash) > 6 {
		hash = hash[:6]
	}

	return string(hash)
}
