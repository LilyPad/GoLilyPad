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
$ go get launchpad.net/goyaml
$ go get code.google.com/p/go-uuid
```

### Error? ###
If you get an error similar to:

```
can't load package: package .: no buildable Go source files in /home/pnap18587/go/pkg/linux_amd64/github.com/LilyPad/GoLilyPad/server/connect/main
```
Then try:
```bash
$ go get -d github.com/LilyPad/GoLilyPad
```

### Server Connect ###

```bash
$ cd $GOPATH/pkg/github.com/LilyPad/GoLilyPad/server/connect/main
```

### Server Proxy ###

```bash
$ cd $GOPATH/pkg/github.com/LilyPad/GoLilyPad/server/proxy/main
```

### Lastly ###

```bash
$ go build
$ ./main
```
