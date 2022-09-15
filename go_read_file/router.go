package app

import (
	"log"

	"github.com/PGITAb/bc-service-api/internal/api"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter() {
	r := gin.Default()
	r.Use(Logger())
	r.GET("/ping", api.Ping)

	r.POST("/player/create", api.PlayerCreate)
	r.POST("/player/update", api.PlayerUpdate)
	r.POST("/player/deposit", api.PlayerDeposit)
	r.POST("/player/withdraw", api.PlayerWithdraw)
	r.POST("/player/balance", api.PlayerBalance)

	r.POST("/game/list", api.GameList)
	r.POST("/game/info", api.GameInfo)

	r.POST("/history", api.History)
	r.POST("/history/roundid", api.HistoryRoundid)
	r.POST("/history/transfer", api.HistoryTransfer)

	log.Fatal(r.Run(viper.GetString("server.port")))
}
