module github.com/ajablonsk1/homelab/notes

go 1.21.4

replace github.com/ajablonsk1/homelab/notes/note => ../note

require github.com/ajablonsk1/homelab/notes/note v0.0.0-00010101000000-000000000000

require (
	github.com/ajablonsk1/homelab/notes/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/flosch/pongo2/v6 v6.0.0
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/ajablonsk1/homelab/notes/config => ../config
