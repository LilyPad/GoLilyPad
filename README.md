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

Docker
--------
## Building ##
```bash
docker build -t lilypad/golilypad .
```

## Server Connect ##
```bash
docker run -itd --name LPConnect lilypad/golilypad connect
```

Running with config loaded from disk:
```bash
docker run -itd --name LPConnect -v `pwd`/connect-config:/data lilypad/golilypad connect
```

## Server Proxy ##
```bash
docker run -itd --name LPProxy --link LPConnect:connect lilypad/golilypad proxy
```

You might need to restart the proxy if you didn't load in a config. 
```bash
docker restart LPPRoxy
```

Running with config loaded from disk:
```bash
docker run -itd --name LPProxy --link LPConnect:connect -v `pwd`/proxy-config:/data lilypad/golilypad proxy
```
this will create a directory called proxy-config where the proxy.yml will be either generated or used. 
