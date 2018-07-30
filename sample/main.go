package main

import (
	"fmt"
	"html/template"
	//	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/onrik/logrus/filename"
	//    "filename"

	"github.com/jessevdk/go-flags"

	mapper "github.com/birkirb/loggers-mapper-logrus"

	"github.com/apisite/gin-mulate"
	"github.com/apisite/mulate"

	"github.com/gin-gonic/gin"
)

// Config holds all config vars
type Config struct {
	Addr string `long:"http_addr" default:"localhost:8081"  description:"Http listen address"`

	Template mulate.Config `group:"Template Options"`
}

func main() {

	cfg := &Config{}
	p := flags.NewParser(cfg, flags.Default)

	_, err := p.Parse()
	if err != nil {
		if !strings.HasPrefix(err.Error(), "\nUsage") {
			fmt.Fprintf(os.Stderr, "error: %+v", err)
		}
		os.Exit(0)
	}

	l := logrus.New()

	if gin.IsDebugging() {
		l.SetLevel(logrus.DebugLevel)
		l.AddHook(filename.NewHook())
	}
	log := mapper.NewLogger(l)

	mlt, _ := mulate.New(cfg.Template, log)
	mlt.DisableCache(true)

	allFuncs := make(template.FuncMap, 0)
	err = mlt.LoadTemplates(allFuncs)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	/*
		r.Use(static.Serve("/", static.LocalFile("./static", false)))
		r.NoRoute(func(c *gin.Context) {
			c.File("static/index.html")
		})
	*/
	templates := ginmulate.New(mlt, log)
	templates.Route("", router)

	router.Run(cfg.Addr)
}