SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o go_cnode main.go router.go 