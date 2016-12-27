# Docsgen

[![Go Report Card](https://goreportcard.com/badge/github.com/jacted/docsgen)](https://goreportcard.com/report/github.com/jacted/docsgen)
[![MIT licensed](https://img.shields.io/github/license/jacted/docsgen.svg?maxAge=2592000)](https://github.com/jacted/docsgen/blob/master/LICENSE)

A CLI tool for creating documentation websites.

## Folder structure

```
 - api_docs (Input folder)
   - Group name
     - Section name 1
       - content.md
       - example.md
     - Section name 2
       - content.md
       - example.md
   - Group name 2
     - Section name 3
       - content.md
       - example.md
```

## Commands

|`docsgen <command>`|Description|
|------------------|-----------|
|`help`|Shows help description.|
|`run`|Creates the HTML files.|
A CLI tool for creating documentation websites.

## CLI Flags

|`Flag`|Description|Options|
|------------------|-----------|---------|
|`theme`|Theme name|"default"|
|`out`|Output folder|
|`in`|Input folder|

## Bindata

This project uses bindata to bundle the themes

go-bindata -o themes.go themes/...

[https://github.com/jteeuwen/go-bindata](https://github.com/jteeuwen/go-bindata)

## Questions and issues

Use the Github issue tracker for any questions, issues or feature suggestions.

## License

MIT licensed.