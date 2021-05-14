module github.com/isaurabhkaushik/hp/service

go 1.14

replace github.com/isaurabhkaushik/hp/api => ./../api

require (
	github.com/gin-gonic/gin v1.7.1
	github.com/go-playground/validator/v10 v10.6.0 // indirect
	github.com/google/uuid v1.1.2
	github.com/isaurabhkaushik/hp/api v0.0.0-00010101000000-000000000000
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/spf13/cast v1.3.1
	github.com/ugorji/go v1.2.5 // indirect
	golang.org/x/crypto v0.0.0-20210506145944-38f3c27a63bf // indirect
	golang.org/x/sys v0.0.0-20210507161434-a76c4d0a0096 // indirect
	golang.org/x/text v0.3.6 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
