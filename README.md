# go-summarizer

Summarize web page by ChatGPT

## Setup

```shell
export OPENAI_API_KEY="YOUR_API_KEY"
```

or if you use [direnv](https://direnv.net/):

```shell
cp .envrc.example .envrc
# edit .envrc
```

## Quick start

```shell
make run URL="http://example.com"
```

## Usage

```shell
go run . --url "http://example.com" --lang "English"
```

## Install

```shell
# Install as `go-summarizer`
make install
go-summarizer --url "http://example.com" --lang "English"
```
