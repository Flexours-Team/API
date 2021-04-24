module github.com/Flexours-Team/API

go 1.16

require (
	github.com/Flexours-Team/models v1.0.0
	github.com/gin-gonic/gin v1.7.1
	github.com/go-openapi/strfmt v0.20.1
	github.com/olivere/elastic v6.2.35+incompatible
	golang.org/x/sys v0.0.0-20210423185535-09eb48e85fd7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.8

)

replace github.com/Flexours-Team/models => ./../models
