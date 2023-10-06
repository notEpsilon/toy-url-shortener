# toy-url-shortener
Simple zero-dependency clean architecture URL shortener based on query params and an in-memory storage

## Usage

First of all, you need [Go installed](https://go.dev/dl/), then follow below instructions:

```shell
git clone https://github.com/notEpsilon/toy-url-shortener.git
```

Then, if you have `Make` installed run this:
```shell
make run
```

Otherwise, run this:
```shell
go run main.go
```

A HTTP server should start listening on `127.0.0.1:7000`.

## Requests

Issue a shortening request first to `/shorten`, for example:
```shell
http://127.0.0.1:7000/shorten?url=https://github.com
```

You will receive a `slug`, copy the value of the `slug` and issue another request to get redirected. Example:
```shell
http://127.0.0.1:7000/r?s=<the value of the slug>
```

You will get redirected to the original URL.
