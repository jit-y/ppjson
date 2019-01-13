ppjson
===

A pretty printer for JSON written in golang.

## Usage

```sh
go get github.com/jit-y/ppjson/cmd/ppjson

echo '[1,"2",null, {"foo": "bar"}]' | ppjson
[
  1,
  "2",
  null,
  {
    "foo": "bar"
  }
]
```

## License

MIT
