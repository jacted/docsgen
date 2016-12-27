# Docsgen

[![Go Report Card](https://goreportcard.com/badge/github.com/jacted/docsgen)](https://goreportcard.com/report/github.com/jacted/docsgen)
[![MIT licensed](https://img.shields.io/github/license/jacted/docsgen.svg?maxAge=2592000)](https://github.com/jacted/docsgen/blob/master/LICENSE)

A CLI tool for creating documentation websites.

## Installation

1. Download [latest release](https://github.com/jacted/docsgen/releases)
2. Put it in a folder and make sure its executeable
3. Add the folder to system PATH
4. You're set, you can now run it using `docsgen` in terminal

## Folder structure

```
 - api_docs (Input folder)
   - 1_group_name
     - 1_section_name
       - content.md
       - example.md
     - 2_section_name
       - content.md
       - example.md
   - 2_group_name
     - 1_section_name
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
|`name`|Docs name|
|`theme`|Theme name|"default"|
|`out`|Output folder|
|`in`|Input folder|

## Bindata

This project uses bindata to bundle the themes.

Every time there is a change to the theme folder the following command needs to be run:

`go-bindata -o themes.go themes/...`

[https://github.com/jteeuwen/go-bindata](https://github.com/jteeuwen/go-bindata)

## Questions and issues

Use the Github issue tracker for any questions, issues or feature suggestions.

## License

MIT licensed.