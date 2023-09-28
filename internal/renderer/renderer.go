package renderer

import (
	"embed"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/unrolled/render"
)

type (
	Views struct {
		r *render.Render
	}
)

func NewViews(_fs embed.FS) *Views {
	return &Views{
		r: render.New(render.Options{
			Directory:  "views",
			FileSystem: &render.EmbedFileSystem{FS: _fs},
			Extensions: []string{".html"},
		}),
	}
}

func (v *Views) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	opts := render.HTMLOptions{}
	layout, hasLayout := ctx.Get("layout").(string)
	if hasLayout {
		opts.Layout = layout
	}
	return v.r.HTML(w, 0, name, data, opts)
}
