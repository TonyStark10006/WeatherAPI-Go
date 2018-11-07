package middlewares

import "net/http"

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//r2 := new(http.Request)
		//*r2 = *r
		//r2.Header.Set("csrf", "")
		if csrf := r.Header.Get("cookies"); csrf != "123" {
			return
		}
		h.ServeHTTP(w, r)
	})
}