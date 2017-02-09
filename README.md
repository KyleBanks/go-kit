# go-kit

[![Build Status](https://travis-ci.org/KyleBanks/go-kit.svg?branch=master)](https://travis-ci.org/KyleBanks/go-kit) &nbsp;
[![GoDoc](https://godoc.org/github.com/KyleBanks/go-kit?status.svg)](https://godoc.org/github.com/KyleBanks/go-kit) &nbsp;
[![Go Report Card](https://goreportcard.com/badge/github.com/KyleBanks/go-kit)](https://goreportcard.com/report/github.com/KyleBanks/go-kit)

This repository contains generic Go packages that are reused throughout various Go projects. 

Most packages are designed to be used standalone, however a few such as `auth` have additional dependencies on other packages in the `go-kit`. 

## Packages

- [auth](./auth) provides generic authentication functionality intended for prototyping applications, and not for real production use.
- [cache](./cache) is a wrapper for Redis, with a mock to use for unit testing without a Redis server.
- [contains](./contains) provides functions for checking if a slice contains a specific value.
- [convert](./convert) provides conversion of various data-types, such as a string slice to int slice, and vice-versa.
- [env](./env) is a simple environment variable wrapper to return an application environment (Dev/Test/Prod) via an environment variable.
- [git](./git) provides the ability to install a pre-commit git hook within Go.
- [gonamo](./gonamo) provides a simple wrapper around the DynamoDB SDK.
- [job](./job) provides the ability to execute tasks on a timed interval.
- [log](./log) is a simple `fmt` wrapper for logging.
- [milliseconds](./milliseconds) gives the ability to get the current time in milliseconds, and the ability to get a specific time in milliseconds.
- [orm](./orm) is a wrapper for [GORM](https://github.com/jinzhu/gorm) with a couple helpful utilities.
- [push](./push) wraps other open-source libraries to provide a simple means of sending push notifications to both iOS and Android devices.
- [router](./router) is a very barebones router implementation that relies heavily on the standard http package and provides a couple little additions.
- [timer](./timer) provides the ability to time abritrary events, like the duration of a method call.
- [today](./today) provides methods related to today's date, such as getting today's date with the time right before midnight.
- [unique](./unique) provides methods to create unique subsets of slices.

## Testing

```
./sanity.sh
```

## License

```
The MIT License (MIT)

Copyright (c) 2017 Kyle Banks

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```