package ginmulate

// https://stackoverflow.com/questions/42747183/how-to-render-templates-to-multiple-layouts-in-go

import (
	"github.com/gin-gonic/gin"

	"html/template"
	"net/http"

	"gopkg.in/birkirb/loggers.v1"

	"github.com/apisite/mulate" // TODO: change to interface
)

// EngineKey holds gin context key name for engine storage
const EngineKey = "github.com/apisite/mulate"

type Config struct {
	mulate.Config
}

// Engine holds template engine attributes
type Template struct {
	*mulate.Template
	FuncHandler func(ctx *gin.Context, funcs template.FuncMap)
	log         loggers.Contextual
	config      Config
}

func New(cfg Config, log loggers.Contextual) *Template {
	return &Template{Template: mulate.New(cfg.Config), config: cfg}
}

// Middleware stores Engine in gin context
func (tmpl *Template) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(EngineKey, tmpl)
	}
}

// Route registers template routes into gin
func (tmpl *Template) Route(prefix string, r *gin.Engine) {
	if prefix != "" {
		prefix = prefix + "/"
	}

	// we need this before page registering
	r.Use(tmpl.Middleware())

	for _, p := range tmpl.Pages() {
		r.GET(prefix+p, tmpl.handleHTML(p)) // TODO: map[content-type]Pages
	}
}

// handleHTML returns gin page handler
func (tmpl *Template) handleHTML(uri string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if val, ok := ctx.Get(EngineKey); ok {
			if t, ok := val.(*Template); ok {
				t.HTML(ctx, uri)
				return
			}
		}
		tmpl.log.Error("Context without valid engine key", EngineKey)
	}
}

// HTML renders page for given uri
func (tmpl *Template) HTML(ctx *gin.Context, uri string) {
	funcs := make(template.FuncMap, 0)
	// Get funcMap copy
	for k, v := range tmpl.Funcs {
		funcs[k] = v
	}
	FuncHandler(ctx, funcs)
	if tmpl.FuncHandler != nil {
		(tmpl.FuncHandler)(ctx, funcs)
	}
	p, err := tmpl.RenderPage(uri, funcs, ctx.Request)
	if err != nil {
		if p.Status == http.StatusMovedPermanently || p.Status == http.StatusFound {
			ctx.Redirect(p.Status, p.Title)
			return
		}
		tmpl.log.Debugf("page error: (%+v)", err)
		if p.Status == http.StatusOK {
			p.Status = http.StatusInternalServerError
			p.Raise(p.Status, "Internal", err.Error(), false)
		}
	}
	renderer := mulate.NewRenderer(tmpl.Template, p)
	ctx.Header("Content-Type", p.ContentType)

	ctx.Render(p.Status, renderer)
}

// FuncHandler is a sample of passing functions to templates
func FuncHandler(ctx *gin.Context, funcs template.FuncMap) {
	funcs["param"] = func(key string) string { return ctx.Param(key) }
}
