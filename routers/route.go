package routers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"com.ddabadi/estock/constants"
	v1 "com.ddabadi/estock/controllers"
	"com.ddabadi/estock/models"
	"com.ddabadi/estock/models/dto"
	"com.ddabadi/estock/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	version string

	nameService string

	debugMode string
)

func init() {
	version = "/v0.1.0"

	debugMode = utils.GetEnv("APPS_DEBUG", "debug")

}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		//AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge: 86400,
	}))

	fmt.Println(gin.IsDebugging())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// r.GET(accessPointTest, v1.TestController)

	var api *gin.RouterGroup

	// AuthController := new(v1.AuthController)
	// api = r.Group(version + "/auth")
	// api.POST("/login", AuthController.Login)
	// api.GET("/:id/reset-password", cekToken, AuthController.ResetPassword)

	CategoryController := new(v1.ProductCategoryController)
	api = r.Group(version + "/product-category")
	api.GET("", CategoryController.GetAll)

	return r

}

func cekToken(c *gin.Context) {

	res := models.Response{}
	tokenString := c.Request.Header.Get("Authorization")
	log.Println("tokenString -> ", tokenString)

	if strings.HasPrefix(tokenString, "Bearer ") == false {
		res.Rc = constants.ERR_CODE_53
		res.Msg = constants.ERR_CODE_53_MSG + " [01] "
		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if jwt.GetSigningMethod("HS256") != token.Method {
			res.Rc = constants.ERR_CODE_53
			res.Msg = constants.ERR_CODE_53_MSG + " [02] "
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			// return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(constants.TokenSecretKey), nil
	})

	if token != nil && err == nil {
		claims := token.Claims.(jwt.MapClaims)

		fmt.Println("claims : ", claims)

		fmt.Println("User name from TOKEN ", claims["user"])

		unixNano := time.Now().UnixNano()
		timeNowInInt := unixNano / 1000000

		tokenCreated := (claims["tokenCreated"])
		dto.CurrUserEmail = (claims["userEmail"]).(string)

		currUserId := (claims["userId"]).(string)
		dto.CurrUserID, _ = strconv.ParseInt(currUserId, 10, 64)

		currRestoId := (claims["restoId"]).(string)
		dto.CurrRestoID, _ = strconv.ParseInt(currRestoId, 10, 64)

		fmt.Println("now : ", timeNowInInt)
		fmt.Println("token created time : ", tokenCreated)
		fmt.Println("user by token : ", dto.CurrUserEmail)
		fmt.Println("user by token ID : ", dto.CurrUserID)

		tokenCreatedInString := tokenCreated.(string)
		tokenCreatedInInt, errTokenExpired := strconv.ParseInt(tokenCreatedInString, 10, 64)

		if errTokenExpired != nil {
			res.Rc = constants.ERR_CODE_53
			res.Msg = constants.ERR_CODE_53_MSG + " [03] "
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		if ((timeNowInInt - tokenCreatedInInt) / 1000) > constants.TokenExpiredInMinutes {
			res.Rc = constants.ERR_CODE_53
			res.Msg = constants.ERR_CODE_53_MSG + " [04] "
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}
		fmt.Println("Token already used for ", (timeNowInInt-tokenCreatedInInt)/1000, "sec, Max expired ", constants.TokenExpiredInMinutes, "sec ")
		// fmt.Println("token Valid ")

	} else {
		res.Rc = constants.ERR_CODE_53
		res.Msg = constants.ERR_CODE_53_MSG + " [05] "
		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
		return
	}
}
