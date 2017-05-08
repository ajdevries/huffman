# huffman [![CircleCI](https://circleci.com/gh/ajdevries/huffman.svg?style=svg)](https://circleci.com/gh/ajdevries/huffman)
> Alten Playground Golang



### Install go
- Download from https://golang.org/dl/

#### Pick an editor or IDE
- Atom (https://atom.io/) + go-plus package
- Visual Studio Code + vscode-go extension (https://marketplace.visualstudio.com/items?itemName=lukehoban.Go)
- Gogland (https://www.jetbrains.com/go/)

#### Go resources
- https://gobyexample.com/
- https://tour.golang.org/welcome/1
- Even more resources https://dave.cheney.net/resources-for-new-go-programmers

### Get source
```
cd $GOPATH/src
go get github.com/ajdevries/huffman
```

### Build server
```
cd $GOPATH/src/github.com/ajdevries/huffman/server
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
