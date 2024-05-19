module github.com/ajablonsk1/homelab/notes/media

go 1.21.4

replace github.com/ajablonsk1/homelab/notes/note => ../note

require (
	github.com/ajablonsk1/homelab/notes/config v0.0.0-00010101000000-000000000000
	github.com/ajablonsk1/homelab/notes/note v0.0.0-00010101000000-000000000000
	golang.org/x/text v0.14.0
)

require gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect

require gopkg.in/yaml.v3 v3.0.1 // indirect

replace github.com/ajablonsk1/homelab/notes/config => ../config
