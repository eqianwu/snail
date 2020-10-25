package routers

import "github.com/gin-gonic/gin"

func SetupRouter() (engine *gin.Engine) {
	engine = gin.Default()
	addHandler(engine)
	return engine
}

func addHandler(engine *gin.Engine) {
	registerGroup(engine)
}