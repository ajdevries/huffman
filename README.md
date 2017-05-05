# huffman
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
