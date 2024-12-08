package middleware

import (
	"log"
	"net/http"
)

func AuthMiddleware() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("auth mid")

			// body, err := io.ReadAll(r.Body)
			// if err != nil {
			// 	log.Println("read body error")
			// 	log.Println(err)
			// 	return
			// }
			// log.Println(string(body))
			// r.Body = io.NopCloser(bytes.NewBuffer(body))

			// var test models.Test
			// err = json.Unmarshal(body, &test)
			// if err != nil {
			// 	os.OpenFile("", os.O_APPEND, 0777)

			// 	var serr *json.SyntaxError
			// 	// var serr *os.SyscallError
			// 	if errors.As(err, &serr) {
			// 		log.Println("GOT IT!!!")
			// 	} else {
			// 		log.Println("NOPE!!!")
			// 	}
			// 	log.Println("parse json error")
			// 	log.Println(err)
			// 	return
			// }
			// log.Printf("Value of key: %s", test.Key)

			h.ServeHTTP(w, r)
		})
	}
}
