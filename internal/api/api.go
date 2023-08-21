package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jexodusmercado/inventory-system-be/internal/conf"
	"gorm.io/gorm"
)

type API struct {
	db *gorm.DB
	config conf.Config
	router *gin.Engine
}

func NewAPI(database *gorm.DB, globalConfig conf.Config, router *gin.Engine) *API {
	// initialize api
	api := &API{db: database, config: globalConfig, router: router}

	return api
}


func (api *API) Serve() {
	port := api.config.Port

	api.router.Run(port)
}
