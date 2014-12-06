export APP_MODE="production"

cd $GOPATH/src/github.com/adred/wiki-player
go clean -i
go build

./wiki-player
