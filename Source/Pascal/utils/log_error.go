package utils

import (
	"log"
	"net/http"
)

func tryLogError(err error, logError func(err error)) bool {
	if err != nil {
		logError(err)
		return true
	}
	return false
}

func TryLogIfError(err error) bool {
	return tryLogError(err, func(err error) {
		log.Println("Error: ", err)
	})
}
func TryLogIfErrorFatal(err error) bool {
	return tryLogError(err, func(err error) {
		log.Fatalln("Error: ", err)
	})
}
func TryLogIfErrorHttp(err error, w http.ResponseWriter) bool {
	if TryLogIfError(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}
