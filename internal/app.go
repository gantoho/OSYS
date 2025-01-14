package internal

import (
	"github.com/gantoho/osys/internal/models"
	"github.com/gantoho/osys/internal/routers"
)

func Start() {
	models.InitDB()
	routers.InitRouters()
}
