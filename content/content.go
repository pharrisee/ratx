package content

import "embed"

//go:embed all:static
var StaticFS embed.FS

//go:embed all:views
var ViewsFS embed.FS
