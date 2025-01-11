package app

import "github.com/gantoho/osys/app/models"

func Start() {
	models.InitDB()
}
