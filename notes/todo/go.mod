module github.com/ajablonsk1/homelab/notes/todo

go 1.21.4

replace github.com/ajablonski/homelab/notes/todo/note => ../note

replace github.com/ajablonsk1/homelab/notes/todo/note => ../note

replace github.com/ajablonsk1/homelab/notes/note => ../note

require github.com/ajablonsk1/homelab/notes/note v0.0.0-00010101000000-000000000000

require github.com/ajablonsk1/homelab/notes/config v0.0.0-00010101000000-000000000000

require (
	github.com/kr/pretty v0.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/ajablonsk1/homelab/notes/config => ../config
