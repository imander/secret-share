package app

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/multitemplate"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

var AppRoot string

func init() {
	root := os.Getenv("APP_ROOT")
	root = strings.TrimRight(root, "/")
	if root == "" || strings.HasPrefix(root, "/") {
		AppRoot = root
	} else {
		AppRoot = "/" + root
	}
}

func New() *http.Server {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logger())
	r.Use(limits.RequestSizeLimiter(100 * 1024))
	r.HTMLRender = loadTemplates("./templates")
	g := r.Group(AppRoot)
	g.GET("/", homeHandler)
	g.POST("/s", SecretPost)
	g.GET("/s/:id", SecretGet)
	g.Static("/assets", "./assets")
	r.NoRoute(notFound)

	s := &http.Server{
		Handler: r,
		TLSConfig: &tls.Config{
			MinVersion:       tls.VersionTLS12,
			CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		},
	}
	return s
}

func logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	fm := template.FuncMap{"appRoot": func() string { return AppRoot }}
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFilesFuncs(filepath.Base(include), fm, files...)
	}
	return r
}
