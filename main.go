package main

import (
	"fmt"
	"github.com/wenchangshou2/gin-blog/pkg/setting"
	"github.com/wenchangshou2/gin-blog/routers"
	"net/http"
)

func main(){
	router:=routers.InitRouter()
	s:=&http.Server{
		Addr:fmt.Sprintf(":%d",setting.HTTPPort),
		Handler:router,
		ReadTimeout:setting.ReadTimeout,
		WriteTimeout:setting.WriteTimeout,
		MaxHeaderBytes:1<<20,
	}
	s.ListenAndServe()
}