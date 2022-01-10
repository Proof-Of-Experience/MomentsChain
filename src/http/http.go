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
func GetAccount(c *gin.Context) {
    log.Infof("In GetAccount")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "Account API",
        "time": now,
    })

}
func GetAccountByKey(c *gin.Context) {
    log.Infof("In GetAccountByKey")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "AccountByKey API",
        "time": now,
    })

}

func GetAccountPrice(c *gin.Context) {
    log.Infof("In GetAccountPrice")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "AccountPrice API",
        "time": now,
    })

}

func GetAccounts(c *gin.Context) {
    log.Infof("In GetAccounts")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "Accounts API",
        "time": now,
    })

}

func AddPeer(c *gin.Context) {
    log.Infof("In AddPeer")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "AddPeer API",
        "time": now,
    })

}

func GetAllMiners(c *gin.Context) {
    log.Infof("In AllMiners")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "AllMiners API",
        "time": now,
    })

}
func GetBlock(c *gin.Context) {
    log.Infof("In GetBlock")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetBlock API",
        "time": now,
    })

}
func GetBlog(c *gin.Context) {
    log.Infof("In GetBlog")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetBlog API",
        "time": now,
    })

}
func GetConfig(c *gin.Context) {
    log.Infof("In GetConfig")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetConfig API",
        "time": now,
    })

}
func GetContent(c *gin.Context) {
    log.Infof("In GetContent")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetContent API",
        "time": now,
    })

}
func GetCount(c *gin.Context) {
    log.Infof("In GetCount")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetCount API",
        "time": now,
    })

}
func GetFeed(c *gin.Context) {
    log.Infof("In GetFeed")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetFeed API",
        "time": now,
    })

}
func GetFollowers(c *gin.Context) {
    log.Infof("In GetFollowers")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetFollowers API",
        "time": now,
    })

}
func GetFollows(c *gin.Context) {
    log.Infof("In GetFollows")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetFollows API",
        "time": now,
    })

}
func GetHistory(c *gin.Context) {
    log.Infof("In GetHistory")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetHistory API",
        "query": c.Query("view"),
        "time": now,
    })

}
func GetHot(c *gin.Context) {
    log.Infof("In GetHot")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetHot API",
        "time": now,
    })

}
func GetImage(c *gin.Context) {
    log.Infof("In GetImage")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetImage API",
        "time": now,
    })

}
func GetLeader(c *gin.Context) {
    log.Infof("In GetLeader")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetLeader API",
        "time": now,
    })

}
func MineBlock(c *gin.Context) {
    log.Infof("In MineBlock")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "MineBlock API",
        "time": now,
    })

}
func GetNew(c *gin.Context) {
    log.Infof("In GetNew")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetNew API",
        "time": now,
    })

}
func GetNewKeyPair(c *gin.Context) {
    log.Infof("In NewKeyPair")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "NewKeyPair API",
        "time": now,
    })

}
func GetPeers(c *gin.Context) {
    log.Infof("In GetPeers")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetPeers API",
        "time": now,
    })

}
func GetRank(c *gin.Context) {
    log.Infof("In GetRank")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetRank API",
        "time": now,
    })

}
func GetRecover(c *gin.Context) {
    log.Infof("In GetRecover")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetRecover API",
        "time": now,
    })

}
func GetRewardPool(c *gin.Context) {
    log.Infof("In GetRewardPool")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetRewardPool API",
        "time": now,
    })

}
func GetRewards(c *gin.Context) {
    log.Infof("In GetRewards")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetRewards API",
        "time": now,
    })

}
func GetSchedule(c *gin.Context) {
    log.Infof("In GetSchedule")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetSchedule API",
        "time": now,
    })

}
func GetSupply(c *gin.Context) {
    log.Infof("In GetSupply")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetSupply API",
        "time": now,
    })

}
func Transact(c *gin.Context) {
    log.Infof("In Transact")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "Transact API",
        "time": now,
    })

}

func TransactWaitConfirm(c *gin.Context) {
    log.Infof("In TransactWaitConfirm")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "TransactWaitConfirm API",
        "time": now,
    })

}

func GetTrending(c *gin.Context) {
    log.Infof("In GetTrending")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "Trending API",
        "time": now,
    })

}
func GetTx(c *gin.Context) {
    log.Infof("In GetTx")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "GetTx API",
        "time": now,
    })

}
func GetVotes(c *gin.Context) {
    log.Infof("In GetVotes")
    now := time.Now().Format("2006-01-02 15:04:05")

    c.JSON(200, gin.H{
        "service": "Votes API",
        "time": now,
    })

}



func SetupRouter(exit chan bool) *gin.Engine {
    r := gin.Default()
    r.GET("/", WhoAmI)
    r.GET("/account", GetAccount)
    r.GET("/accountByKey", GetAccountByKey)
    r.GET("/accountprice", GetAccountPrice)
    r.GET("/accounts", GetAccounts)
    r.GET("/addpeer", AddPeer)
    r.GET("/allminers", GetAllMiners)
    r.GET("/block", GetBlock)
    r.GET("/blog", GetBlog)
    r.GET("/config", GetConfig)
    r.GET("/content", GetContent)
    r.GET("/count", GetCount)
    r.GET("/feed", GetFeed)
    r.GET("/followers", GetFollowers)
    r.GET("/follows", GetFollows)
    r.GET("/history", GetHistory)
    r.GET("/hot", GetHot)
    r.GET("/image", GetImage)
    r.GET("/leader", GetLeader)
    r.GET("/mineblock", MineBlock)
    r.GET("/new", GetNew)
    r.GET("/newkeypair", GetNewKeyPair)
    r.GET("/peers", GetPeers)
    r.GET("/rank", GetRank)
    r.GET("/recover", GetRecover)
    r.GET("/rewardPool", GetRewardPool)
    r.GET("/rewards", GetRewards)
    r.GET("/schedule", GetSchedule)
    r.GET("/supply", GetSupply)
    r.GET("/transact", Transact)
    r.GET("/transactWaitConfirm", TransactWaitConfirm)
    r.GET("/trending", GetTrending)
    r.GET("/tx", GetTx)
    r.GET("/votes", GetVotes)

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
