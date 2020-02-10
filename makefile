SERVER="reptilia-server"
AGENT="reptilia-agent"
VERSION=`git symbolic-ref --short HEAD`
GIT_COMMIT=`git rev-parse HEAD`
BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%S%z"`

default:
	@echo ${VERSION} ${GIT_COMMIT} ${BUILD_DATE}
	@go build -o ${SERVER} -ldflags="-X github.com/longbei65/reptilia/version.AppName=${SERVER} -X github.com/longbei65/reptilia/version.GitCommit=${GIT_COMMIT} -X github.com/longbei65/reptilia/version.Version=${VERSION} -X github.com/longbei65/reptilia/version.BuildDate=${BUILD_DATE}" ./cmd/server/main.go 
	@go build -o ${AGENT} -ldflags="-X github.com/longbei65/reptilia/version.AppName=${AGENT} -X github.com/longbei65/reptilia/version.GitCommit=${GIT_COMMIT} -X github.com/longbei65/reptilia/version.Version=${VERSION} -X github.com/longbei65/reptilia/version.BuildDate=${BUILD_DATE}" ./cmd/agent/main.go 

info:
	@echo Version: ${VERSION} 
	@echo GIT_COMMIT: ${GIT_COMMIT} 
	@echo BUILD_DATE: ${BUILD_DATE}

server:
	@go build -o ${SERVER} -ldflags="-X github.com/longbei65/reptilia/version.AppName=${SERVER} -X github.com/longbei65/reptilia/version.GitCommit=${GIT_COMMIT} -X github.com/longbei65/reptilia/version.Version=${VERSION} -X github.com/longbei65/reptilia/version.BuildDate=${BUILD_DATE}" ./cmd/server/main.go 

agent:
	@go build -o ${AGENT} -ldflags="-X github.com/longbei65/reptilia/version.AppName=${AGENT} -X github.com/longbei65/reptilia/version.GitCommit=${GIT_COMMIT} -X github.com/longbei65/reptilia/version.Version=${VERSION} -X github.com/longbei65/reptilia/version.BuildDate=${BUILD_DATE}" ./cmd/agent/main.go 