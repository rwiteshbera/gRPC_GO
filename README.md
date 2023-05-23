Install protocol compiler plugins
```bash
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

Generate gRPC code
```bash
$ make gen
```

Delete generated gRPC code
```bash
$ make clean
```