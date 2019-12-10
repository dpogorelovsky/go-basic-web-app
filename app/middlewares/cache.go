package middlewares

import (
	"log"
	"net/http"
)

/*@todo
best hash function
do we need hash?
how to keep data?
if just uri as key - how to flush channel?
redis key prefix? should we rely on redis only?
*/

// Cacheable - will fetch response frpom if exists
func Cacheable(next http.Handler, cacheEngine string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println("r.RequestURI")
		log.Println(r.RequestURI)
		log.Println(cacheEngine)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
