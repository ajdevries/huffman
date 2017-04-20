# huffman
> Alten Playground Golang

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
