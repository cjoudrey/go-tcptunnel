# go-tcptunnel

go-tcptunnel is a simple TCP port forwarder that is based on [tcptunnel](http://www.vakuumverpackt.de/tcptunnel/). This tool listens to a local TCP port and all the received data is sent to a remote host.

## Usage

```
$ go-tcptunnel --help
Usage: ./go-tcptunnel [options...]

Options:

  -bind-address="127.0.0.1": bind address
  -buffer-size=512: buffer size
  -local-port=6000: local port
  -log=false: log
  -remote-host="": remote host
  -remote-port=0: remote port
```
