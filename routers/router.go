package routers

//路由註冊

import (
	"firstweb/api"
	"firstweb/logrus"
	"firstweb/model"

	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {

	router := gin.New()
	//router := gin.Default()

	router.Use(logrus.Logrus())
	router.RedirectFixedPath = true

	

	//跨來源資源共用
	router.Use(cors.New(cors.Config{
		//允許的HTTP Method
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		//允許的Header 信息
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "signature"},
		//允許的domain
		AllowAllOrigins: true,
		//允許請求包含驗證憑證
		AllowCredentials: false,
		//可被存取的時間
		MaxAge: 12 * time.Hour,
	}))

	//set Logger
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/login", api.UserLogin)

	router.NoRoute(func(c *gin.Context) {

		c.JSON(400, gin.H{"error": "Bad Request"})
	})

	model.APIServer = http.Server{
		Addr:    ":7788",
		Handler: router,
	}

}
func RunRouter() { //把listen的功能用go routine的方式丟到背景去運行
	//接到os強制關閉服務

	if err := model.APIServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
