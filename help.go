package coord

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func help(ctx *gin.Context) {
	ctx.String(http.StatusOK, `/coord/transform?from=<type>&to=<type>&longitude=<longitude>&latitude=<latitude>
`)
}
