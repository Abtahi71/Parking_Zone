package main

import (
	"gotickets/internal/config"
	"gotickets/internal/server"
)

func main() {
  cfg:=config.LoadEnv()
  db:=config.ConnectDb(cfg)
  server.Start(cfg,db)
  }
