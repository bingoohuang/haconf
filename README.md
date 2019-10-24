# mci(mysqlclusterinit)

a tool to modify HAConfig conf file and restart HAConfig.

## Build

`env GOOS=linux GOARCH=amd64 go install ./...`

## Usage

```bash
➜  haconf haconf -h
Usage of haconf:
      --BackAddrs string             BackAddrs
      --Debug                        Debug
      --HAProxyCfg string            HAProxyCfg
      --HAProxyRestartShell string   HAProxyRestartShell
      --ListenName string            ListenName
      --Port int                     Port
      --check                        check HAConf
  -c, --config string                config file path (default "./config.toml")
  -v, --version                      show version
pflag: help requested
```

## Demo

```bash
➜  haconf --ListenName yc --Port 1990 --BackAddrs 127.0.0.1:8888,127.0.0.2:8888 --Debug --HAProxyCfg testdata//test.cfg
INFO[0000] HAProxy:
listen yc
  bind 127.0.0.1:1990
  mode tcp
  option tcpka
  server s-1 127.0.0.1:8888 check inter 1s
  server s-2 127.0.0.2:8888 check inter 1s
➜  more testdata/test.cfg
#ycStart
#ycEnd
➜  haconf --ListenName yc --Port 1990 --BackAddrs 127.0.0.1:8888,127.0.0.2:8888  --HAProxyCfg testdata//test.cfg
INFO[0000] prepare to overwriteHAProxyCnf
listen yc
  bind 127.0.0.1:1990
  mode tcp
  option tcpka
  server s-1 127.0.0.1:8888 check inter 1s
  server s-2 127.0.0.2:8888 check inter 1s
INFO[0000] overwriteHAProxyCnf completed
ERRO[0000] error error exiting code 127, stdout:[], stderr:[]
```
