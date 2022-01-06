package http

import (
    log "github.com/sirupsen/logrus"
    "github.com/gin-gonic/gin"

    "time"
)

func WhoAmI(c *gin.Context) {
    log.Infof("In WhoAmI")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "http",
        "time": now,
    })

}

func SetupRouter(exit chan bool) *gin.Engine {
    r := gin.Default()
    r.GET("/", WhoAmI)

    return r
}

func StartHTTP() {
    log.Debugf("Starting http server")

    exit := make(chan bool)
    var r *gin.Engine
    r = SetupRouter(exit)
    port := ":9001"
    r.Run(port)
}
