# reconf: Generate config file and run command

[![Test Status][test-badge]][test-url]
[![Release][release-badge]][release-url]

`reconf` generates files from templates and environment variables and
executes a command.

- [Build](#build)
- [Usage](#usage)
- [Test](#test)
- [License](#license)

[test-badge]: https://github.com/snsinfu/reconf/workflows/test/badge.svg
[test-url]: https://github.com/snsinfu/reconf/actions?query=workflow%3Atest
[release-badge]: https://img.shields.io/github/release/snsinfu/reconf.svg
[release-url]: https://github.com/snsinfu/reconf/releases


## Build

Requires a Go compiler. The following command builds `reconf` executable:

```console
$ go build -o reconf github.com/snsinfu/reconf
```


## Usage

`reconf` is a statically-built CLI utility. It accepts the following flags:

```
Usage: reconf [-f -w <file> ...] [<command>...]

  <command>   Command to execute. If command is not given, reconf will
              just generate files and exit.

Options:
  -w, --render <file>  Generate <file> (if it does not exist) by rendering
                       template file named "<file>.template".
  -f, --force          Force generating files, overwriting existing ones.
  -h, --help           Show this usage message and exit.
```

For example, the following command line generates a configuration file
`/srv/nginx.conf` and runs nginx with the generated configuration. The file
is generated from a template file `/srv/nginx.conf.template` using environment
variables available as `{{ .env.NAME }}`.

```console
$ reconf -w /srv/nginx.conf nginx -c /srv/nginx.conf
```

So, if the template contains the following `listen` stanza

```
listen {{.env.PORT}};
```

and if you provide `PORT` environment variable, `reconf` interpolates the
template part `{{.env.PORT}}` with the actual value:

```
$ PORT=8080 reconf -w /srv/nginx.conf nginx -c /srv/nginx.conf
nginx listens on port 8080
```

Passing `-f` (force) option makes sure the config file is always generated. So,
the following command line updates the existing configuration file, and now
the server listens on the different port number:

```console
$ PORT=9000 reconf -fw /srv/nginx.conf nginx -c /srv/nginx.conf
nginx listens on port 9000
```

Without the `-f` option, `reconf` skips to generate existing files.


## Test

Run unit tests:

```console
$ git clone https://github.com/snsinfu/reconf.git
$ cd reconf
$ go test .
```

Integration tests:

```console
$ tests/run
```


## License

MIT License.
