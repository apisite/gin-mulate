package ginmulate

// https://stackoverflow.com/questions/42747183/how-to-render-templates-to-multiple-layouts-in-go

import (
	"github.com/gin-gonic/gin"

	"html/template"
	"net/http"

	"gopkg.in/birkirb/loggers.v1"

	"github.com/apisite/mulate" // TODO: change to interface
)

const EngineKey = "github.com/apisite/mulate"

type Engine struct {
	log loggers.Contextual
	mlt *mulate.Template
}

func New(mlt *mulate.Template, log loggers.Contextual) *Engine {
	rv := Engine{
		mlt: mlt,
		log: log,
	}
	return &rv
}

func (e *Engine) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(EngineKey, e)
	}
}

func (e *Engine) Route(prefix string, r *gin.Engine) {
	if prefix != "" {
		prefix = prefix + "/"
	}

	// we need this before page registering
	r.Use(e.Middleware())

	for _, p := range e.mlt.Pages() {
		r.GET(prefix+p, e.handleHTML(p)) // TODO: map[content-type]Pages
	}
}

func (e *Engine) handleHTML(uri string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if val, ok := ctx.Get(EngineKey); ok {
			if e, ok := val.(*Engine); ok {
				e.HTML(ctx, uri)
				return
			}
		}
		e.log.Error("Context without valid engine key", EngineKey)
	}
}

func (e *Engine) HTML(ctx *gin.Context, uri string) {
	allFuncs := make(template.FuncMap, 0)
	p, err := e.mlt.RenderPage(uri, allFuncs)
	if err != nil {
		if p.Status == http.StatusMovedPermanently || p.Status == http.StatusFound {
			ctx.Redirect(p.Status, p.Title)
			return
		}
		e.log.Debugf("page error: (%+v)", err)
		if p.Status == http.StatusOK {
			p.Status = http.StatusInternalServerError
			p.SetError(p.Status, "Internal", err.Error(), false)
		}
	}
	renderer := mulate.NewRenderer(e.mlt, p)
	ctx.Render(p.Status, renderer)
}