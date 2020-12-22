# onetimecode

onetimecode is a Go package for generating one-time codes that are typically used in SMS verifications.

The default mode is "numbers", i.e. a purely numeric code. You can change that with the flag `--mode`.

The default length of the returned codes is always 6 chars. You can change that with the flag `--length`.

## Usage

```
go run cmd/main.go --mode=numbers --length=6

go run cmd/main.go --mode=alphanum --length=6

go run cmd/main.go --mode=alphanumuc --length=6

go rund cmd/main.go --mode=random --min=1 --max=10
```

- `numbers` returns only numbers, i.e. 123456
- `alphanum` returns an alphanumeric code, i.e. ie3cbk
- `alphanumuc` returns an alphanumeric **uppercase** code, i.e. YZNF44
- `random` returns a random number between the min and the max numbers (default: 1-10)

## Installation

```
go get -u github.com/jakoubek/onetimecode
```