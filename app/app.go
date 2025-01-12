package app

import (
	"github.com/gantoho/osys/app/models"
	"github.com/gantoho/osys/app/router"
)

func Start() {
	models.InitDB()
	router.InitRouter()
}
