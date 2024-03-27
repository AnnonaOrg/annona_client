package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/AnnonaOrg/annona_client/internal/dbredis"

	_ "github.com/AnnonaOrg/annona_client/internal/log"
	"github.com/AnnonaOrg/annona_client/internal/repository"
)

func main() {
	if err := dbredis.Init(); err != nil {
		fmt.Println("初始化数据库(REDIS)失败: %v", err)
	}
	defer dbredis.Close()

	go listenTdlib()

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		repository.Tdlib.Stop()
		os.Exit(1)
	}()
	select {}
	fmt.Println("main Exit")
}
