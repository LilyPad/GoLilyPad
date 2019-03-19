GoLilyPad
=============

An implementation of LilyPad in the language Go.

You can visit us on our website @ http://www.lilypadmc.org

Compilation
-------------

You can currently compile either Server Proxy or Server Connect.

Install Go 1.11+ and pull the project:
```bash
$ export GO111MODULE=on
$ git clone git@github.com:LilyPad/GoLilyPad.git
```

### Server Connect ###

```bash
$ cd GoLilyPad/server/connect/main
```

### Server Proxy ###

```bash
$ cd GoLilyPad/server/proxy/main
```

### Lastly ###

```bash
$ go build
$ ./main
```
