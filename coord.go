package coord

import (
	"net/http"
	"strconv"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"

	// 感觉生成的有问题
	"github.com/qichengzx/coordtransform"
)

func Transform(c *gin.Context) {

	var coord Coordinates
	if err := c.ShouldBind(&coord); err != nil {
		ezap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ezap.Debugf("coordinates transform task: %+v", coord)
	lat, lon := doTask(&coord)

	if coord.H {
		latStr := strconv.FormatFloat(lat, 'f', -1, 64)
		lonStr := strconv.FormatFloat(lat, 'f', -1, 64)
		c.String(http.StatusOK, latStr+", "+lonStr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"latitude": lat, "longitude": lon})
}

func doTask(coord *Coordinates) (float64, float64) {
	switch coord.From + coord.To {
	case "bd09" + "gcj02":
		return coordtransform.BD09toGCJ02(coord.Latitude, coord.Longitude)
	case "gcj02" + "bd09":
		return coordtransform.GCJ02toBD09(coord.Latitude, coord.Longitude)
	case "wgs84" + "gcj02":
		return coordtransform.WGS84toGCJ02(coord.Latitude, coord.Longitude)
	case "gcj02" + "wgs84":
		return coordtransform.GCJ02toWGS84(coord.Latitude, coord.Longitude)
	case "bd09" + "wgs84":
		return coordtransform.BD09toWGS84(coord.Latitude, coord.Longitude)
	case "wgs84" + "bd09":
		return coordtransform.WGS84toBD09(coord.Latitude, coord.Longitude)
	default:
		return 0, 0
	}
}
