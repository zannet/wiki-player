export APP_MODE="test"

cd $GOPATH/src/github.com/adred/wiki-player
go clean -i
go build

./wiki-player
