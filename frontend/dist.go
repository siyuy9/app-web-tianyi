package frontend

import (
	"embed"
)

//go:embed dist/static
var Static embed.FS

//go:embed dist/index.html
var Index embed.FS

//go:embed dist/favicon.ico
var Favicon embed.FS
