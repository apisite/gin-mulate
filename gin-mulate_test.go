package ginmulate

import (
	"html/template"
	"reflect"
	"testing"

	"github.com/apisite/mulate"
	"github.com/gin-gonic/gin"
	"gopkg.in/birkirb/loggers.v1"
)

func TestNew(t *testing.T) {
	type args struct {
		mlt *mulate.Template
		log loggers.Contextual
	}
	tests := []struct {
		name string
		args args
		want *Engine
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := New(tt.args.mlt, tt.args.log); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. New() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestEngine_Middleware(t *testing.T) {
	type fields struct {
		FuncHandler func(ctx *gin.Context, funcs template.FuncMap) template.FuncMap
		log         loggers.Contextual
		mlt         *mulate.Template
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		e := &Engine{
			FuncHandler: tt.fields.FuncHandler,
			log:         tt.fields.log,
			mlt:         tt.fields.mlt,
		}
		if got := e.Middleware(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Engine.Middleware() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestEngine_Route(t *testing.T) {
	type fields struct {
		FuncHandler func(ctx *gin.Context, funcs template.FuncMap) template.FuncMap
		log         loggers.Contextual
		mlt         *mulate.Template
	}
	type args struct {
		prefix string
		r      *gin.Engine
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		e := &Engine{
			FuncHandler: tt.fields.FuncHandler,
			log:         tt.fields.log,
			mlt:         tt.fields.mlt,
		}
		e.Route(tt.args.prefix, tt.args.r)
	}
}

func TestEngine_handleHTML(t *testing.T) {
	type fields struct {
		FuncHandler func(ctx *gin.Context, funcs template.FuncMap) template.FuncMap
		log         loggers.Contextual
		mlt         *mulate.Template
	}
	type args struct {
		uri string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   gin.HandlerFunc
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		e := &Engine{
			FuncHandler: tt.fields.FuncHandler,
			log:         tt.fields.log,
			mlt:         tt.fields.mlt,
		}
		if got := e.handleHTML(tt.args.uri); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Engine.handleHTML() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestEngine_HTML(t *testing.T) {
	type fields struct {
		FuncHandler func(ctx *gin.Context, funcs template.FuncMap) template.FuncMap
		log         loggers.Contextual
		mlt         *mulate.Template
	}
	type args struct {
		ctx *gin.Context
		uri string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		e := &Engine{
			FuncHandler: tt.fields.FuncHandler,
			log:         tt.fields.log,
			mlt:         tt.fields.mlt,
		}
		e.HTML(tt.args.ctx, tt.args.uri)
	}
}

func TestFuncHandler(t *testing.T) {
	type args struct {
		ctx   *gin.Context
		funcs template.FuncMap
	}
	tests := []struct {
		name string
		args args
		want template.FuncMap
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := FuncHandler(tt.args.ctx, tt.args.funcs); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FuncHandler() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
