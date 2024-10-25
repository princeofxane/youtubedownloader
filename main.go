package main

import (
	"fmt"
	"log"
	"net/http"
	"youtube_downloader/pkg/config"
	"youtube_downloader/pkg/server"
)

func main() {
	cnf := config.ReadConfig()
	r := server.New(cnf)
	log.Printf("server has started at port: %s", cnf.ServerCfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cnf.ServerCfg.Port), r))
}
