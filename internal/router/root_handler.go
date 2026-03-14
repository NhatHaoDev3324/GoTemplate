package router

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {

	accept := c.GetHeader("Accept")
	ua := c.GetHeader("User-Agent")

	if strings.Contains(accept, "text/html") &&
		!strings.Contains(ua, "curl") &&
		!strings.Contains(ua, "wget") {
		c.HTML(200, "index.html", nil)
		return
	}

	c.Header("Content-Type", "text/plain")
	c.String(200, `
   _____       _______                  _       _       
  / ____|     |__   __|                | |     | |      
 | |  __  ___    | | ___ _ __ ___  _ __| | __ _| |_ ___ 
 | | |_ |/ _ \   | |/ _ \ '_ ' _ \| '__| |/ _' | __/ _ \
 | |__| | (_) |  | |  __/ | | | | | |  | | (_| | ||  __/
  \_____|\___/   |_|\___|_| |_| |_|_|  |_|\__,_|\__\___|

🚀 Go Clean Architecture API Server

GitHub:
https://github.com/NhatHaoDev3324/go-gin-gorm-postgres-template
`)
}
