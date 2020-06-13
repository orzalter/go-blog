package main

import (
	"fmt"
	"net/http"

	"github.com/orzalter/go-blog/pkg/setting"
	"github.com/orzalter/go-blog/routers"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
