go env -w GOPROXY="https://goproxy.cn,direct"
go env -w GO111MODULE=on 

GO111MODULE=on go get -d golang.org/x/tools/gopls@latest
GO111MODULE=on go get golang.org/x/tools/gopls@latest

GO111MODULE=on go get  github.com/gin-gonic/gin@latest //it worked


git config --global --add url."git@github.com:".insteadOf "https://github.com/"

go clean -modcache
env GIT_TERMINAL_PROMPT=1 go get github.com/gin-gonic/gin@latest

export GIT_TERMINAL_PROMPT=1

go get -u github.com/gin-gonic/gin

mkdir cmd
mkdir config
mkdir internal
mkdir pkg


touch cmd/main.go
touch cmd/server.go

touch config/config.go

goimports -w .