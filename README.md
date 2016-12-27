GoLilyPad
=============

An implementation of LilyPad in the language Go.

You can visit us on our website @ http://www.lilypadmc.org

Compilation
-------------

You can currently compile either Server Proxy or Server Connect.

Pull the project and get the dependencies:
```bash
$ go get github.com/LilyPad/GoLilyPad
$ go get github.com/satori/go.uuid
$ go get gopkg.in/yaml.v2
$ go get github.com/klauspost/compress/zlib
```

### Server Connect ###

```bash
$ cd $GOPATH/src/github.com/LilyPad/GoLilyPad/server/connect/main
```

### Server Proxy ###

```bash
$ cd $GOPATH/src/github.com/LilyPad/GoLilyPad/server/proxy/main
```

### Lastly ###

```bash
$ go build
$ ./main
```
