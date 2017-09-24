# Textify [![Apache-2.0 License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0) [![GoDoc](https://godoc.org/github.com/jonathan-robertson/textify?status.svg)](https://godoc.org/github.com/jonathan-robertson/textify)

Turns popular document formats into plain text for searching, indexing, etc.

## Purpose

I began working on [LockedArchive](https://github.com/jonathan-robertson/lockedarchive) and really wanted a golang package that could extract text from PDF files.

There are some interesting packages available (such as [UniDoc](https://github.com/unidoc/unidoc) and [docconv](https://github.com/sajari/docconv)), but I prefer avoiding any non-golang dependencies and like the freedom of the Apache-2.0 license.

## Install

`go get -u github.com/jonathan-robertson/textify`

## Acknowledgements

- Rus Cox for his incredibly helpful [pdf](https://godoc.org/rsc.io/pdf) package. I [forked his package](https://github.com/jonathan-robertson/pdf) to have it return spaces since he was excluding them from output.