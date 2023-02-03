# Usage

## As a command

```sh
$ go run go.ipse.one/ugen/cmd
D5XlMsS2e5Xb
$ go run go.ipse.one/ugen/cmd -h
Usage:
  ugen [flags]

Flags:
  -a, --any             Use any character from the random generator.
      --base64          Output base64 encoded strings. Sets --any.
  -c, --count int       The number of unique strings to generate. (default 1)
  -h, --help            help for ugen
      --hex             Output hex encoded strings. Sets --any.
  -l, --length int      Length of the generated string. (default 12)
      --lower           Output characters in lower case.
      --prefix string   Write prefix before each generated string.
      --suffix string   Write suffix after each generated string.
      --ulid            Generate a ULID.
      --upper           Output characters in upper case.
      --uuid            Generate a random (v4) UUID.
```

## As a module

```sh
go get -u go.ipse.one/ugen
```
