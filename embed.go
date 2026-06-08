package embed

import "embed"

//go:embed all:templates all:static
var Files embed.FS
