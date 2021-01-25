# reconf: Generate config file and run command

[![Test Status][test-badge]][test-url]

`reconf` generates files from templates and executes a command using
environment variables.

```console
$ ls
server server.conf.template
$ reconf -w server.conf ./server
server.conf is created and the server starts
```

[test-badge]: https://github.com/snsinfu/reconf/workflows/test/badge.svg
[test-url]: https://github.com/snsinfu/reconf/actions?query=workflow%3Atest


## Build

```console
$ go build -o reconf github.com/snsinfu/reconf
```


## Test

```console
$ git clone https://github.com/snsinfu/reconf.git
$ cd reconf
$ go test .
```


## License

MIT License.
