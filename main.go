package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"

	_ "com.ddabadi/estock/database"
	"com.ddabadi/estock/routers"
	"com.ddabadi/estock/utils"
)

var (
	port string
)

func main() {

	// gin.DisableConsoleColor()

	// f, _ := os.Create("e-stock.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	maxProc, _ := strconv.Atoi(utils.GetEnv("MAXPROCS", "1"))
	port = utils.GetEnv("PORT", "8001")
	runtime.GOMAXPROCS(maxProc)

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", port)

	log.Println("[info] start http server listening %s", endPoint)

	routersInit.Run(":" + port)

	// r := gin.Default()

	// r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))

	// r.Use(gin.Recovery())

	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run()
}
