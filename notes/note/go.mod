module github.com/ajablonsk1/homelab/notes/common

go 1.21.4

require (
	github.com/ajablonsk1/homelab/notes/config v0.0.0-00010101000000-000000000000
	github.com/flosch/pongo2/v6 v6.0.0
)

require gopkg.in/yaml.v3 v3.0.1 // indirect

replace github.com/ajablonsk1/homelab/notes/config => ../config
