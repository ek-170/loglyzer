package controllers

import (
	"fmt"
	"myfishing/backend/config"
	"myfishing/backend/consts"
	"myfishing/backend/models"
	"net/http"
	"regexp"
)

func ParseIdFromURL(r *http.Request, validPath *regexp.Regexp, targetPath int) string {

	q := validPath.FindStringSubmatch(r.URL.Path)
	fmt.Println(q)
	if len(q) == 0 {
		return ""
	}
	return q[targetPath]
}

func StartMainServer() error {
	var dr = models.NewDiaryRepository()
	var dh = NewDiaryHandler(dr)

	// var tr = models.NewTestRepository()
	// var th = NewTestHandler(tr)

	http.Handle("/", http.FileServer(http.Dir(consts.Index)))

	// http.HandleFunc("/test/", th.HandleTest)
	http.HandleFunc("/diaries/", dh.HandleDiary)

	// http.HandleFunc("/search", )
	// http.HandleFunc("/register", )
	// http.HandleFunc("/login", )
	// http.HandleFunc("/config", )
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
