export APP_MODE="real"
export APP_BRANCH="dev"

cd $GOPATH/src/github.com/adred/wiki-player
go clean -i
go build

./wiki-player
