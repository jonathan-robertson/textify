# Textify [![Apache-2.0 License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](https://github.com/puddingfactory/textify/blob/master/LICENSE)

Turns popular document formats into plain text for searching, indexing, etc.

## Purpose

I began working on [VaultedPages](https://github.com/vaultedpages/vaultedpages) and needed (deeply wanted!) a golang package that could extract text from PDF files.

There are some interesting packages available (such as [UniDoc](https://github.com/unidoc/unidoc) and [docconv](https://github.com/sajari/docconv)), but I prefer avoiding any non-golang dependencies and like the freedom of the Apache-2.0 license.

## Install

`go get -u github.com/puddingfactory/textify`

## Acknowledgements

- Rus Cox for his incredibly helpful [pdf](https://godoc.org/rsc.io/pdf) package. I [forked his package](https://github.com/puddingfactory/pdf) to have it return spaces since he was excluding them from output.