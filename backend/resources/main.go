package resources

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func SetupRoutes() {
	r := gin.Default()
	username, userNameSetted := os.LookupEnv("user")
	password, passwordSetted := os.LookupEnv("password")
	compress := r.Group("/")
	compress.Use(gzip.Gzip(gzip.BestCompression))
	var stdApi *gin.RouterGroup = nil
	var defaultAuth *gin.RouterGroup = nil
	if userNameSetted && passwordSetted {
		stdApi = compress.Group("/api", gin.BasicAuth(gin.Accounts{
			username: password,
		},
		))
		defaultAuth = r.Group("/api", gin.BasicAuth(gin.Accounts{
			username: password,
		},
		))
	} else {
		stdApi = compress.Group("/api")
		defaultAuth = r.Group("/api")
	}

	InitLibraryApi(stdApi)
	InitBookApi(stdApi, defaultAuth)
	InitCollectionApi(stdApi)
	InitUploadApi(defaultAuth)

	r.Use(gzip.Gzip(gzip.BestCompression), func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=604800, immutable")
		static.Serve("/", static.LocalFile("./bundles", true))(c)
	})
	r.NoRoute(func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=604800, immutable")
		c.File("./bundles/index.html")
	})
	if err := r.Run(":8080"); err != nil {
		log.Println(err)
	}
}
