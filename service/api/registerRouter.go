package api

import (
	"GoGinServerBestPractice/service/api/base"
	"GoGinServerBestPractice/service/api/user"
)

func RegisterRouter() {
	base.Router()
	user.Router()
}
