package coord

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoute(g *gin.Engine) {
	if !viper.GetBool(ID) && !viper.GetBool("allservices") {
		return
	}
	// list route
	g.GET("/help/"+ID, help)

	r := g.Group(RoutePrefix)
	r.GET("/tansform", Transform)
	r.POST("/tansform", Transform)
}
