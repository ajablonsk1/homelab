module github.com/ajablonsk1/homelab/notes/common

go 1.21.4

require github.com/ajablonsk1/homelab/notes/config v0.0.0-00010101000000-000000000000

require (
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/ajablonsk1/homelab/notes/config => ../config
