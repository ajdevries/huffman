# huffman
> Alten Playground Golang

### Install go
- Download from https://golang.org/dl/

#### Pick an IDE
- Install Atom https://atom.io/ + go-plus
- Install Gogland https://www.jetbrains.com/go/
- Visual Studio Code (https://marketplace.visualstudio.com/items?itemName=lukehoban.Go)

### Build server

```
cd $GOPATH/src
mkdir -p github.com/ajdevries
cd github.com/ajdevries
git clone git@github.com:ajdevries/huffman.git
cd huffman/server
go build
```

### Test

```
go test ./... -race
```

### Start server

```
cd $GOPATH/src/github.com/ajdevries/huffman/server
go build && ./server
```
