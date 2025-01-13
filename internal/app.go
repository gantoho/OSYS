package internal

import (
	"github.com/gantoho/osys/internal/model"
	"github.com/gantoho/osys/internal/router"
)

func Start() {
	model.InitDB()
	router.InitRouter()
}
