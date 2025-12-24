package public

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/jeanhua/ZanaoAPI/config"
	"github.com/jeanhua/ZanaoAPI/utils"
)

// /ping
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// 主页帖子
// /home?school=scu&from=0
func HomePage(c *gin.Context) {
	schoolAlias := c.Query("school")
	from := c.Query("from")
	client := resty.New()
	token := config.GetCfonfig().GetString("token")
	client.SetHeaders(utils.GetHeaders(token, schoolAlias))
	var data map[string]any
	_, err := client.R().SetResult(&data).Post("https://api.x.zanao.com/thread/v2/list?from_time=" + from + "&with_comment=true&with_reply=true")
	if err != nil {
		c.String(200, "这不是你的问题，也不是我的问题")
		return
	}
	c.JSON(200, data)
}

// 热贴
// /hot?school=scu
func Hot(c *gin.Context) {
	schoolAlias := c.Query("school")
	client := resty.New()
	token := config.GetCfonfig().GetString("token")
	client.SetHeaders(utils.GetHeaders(token, schoolAlias))
	var data map[string]any
	_, err := client.R().SetResult(&data).Post("https://api.x.zanao.com/thread/hot?count=10&type=3")
	if err != nil {
		c.String(200, "这不是你的问题，也不是我的问题")
		return
	}
	c.JSON(200, data)
}
