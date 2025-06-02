package main

import (
	"github.com/NarthurN/TODO-API-web/internal/config"
	"github.com/NarthurN/TODO-API-web/internal/server"
	"github.com/NarthurN/TODO-API-web/pkg/loger"
)

func main() {
	loger.Init()
	config.Init()

	server := server.New()

	if err := server.Run(); err != nil {
		loger.L.Error("Ошибка при запуске сервера", "err", err)
	}
}
