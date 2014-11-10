export APP_MODE="mock"
export APP_BRANCH="dev"

cd $GOPATH/src/github.com/adred/wiki-player
go clean -i
go build

./wiki-player
