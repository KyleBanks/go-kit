# go-kit

[![Build Status](https://travis-ci.org/KyleBanks/go-kit.svg?branch=master)](https://travis-ci.org/KyleBanks/go-kit)

This repository contains generic Go packages that are reused throughout various Go projects.

## Testing

```
go test $(go list ./... | grep -v /vendor/)
```
