gojsonq
===============
[![Build Status](https://travis-ci.org/thedevsaddam/gojsonq.svg?branch=master)](https://travis-ci.org/thedevsaddam/gojsonq)
[![Project status](https://img.shields.io/badge/version-v1.1-green.svg)](https://github.com/thedevsaddam/gojsonq/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/thedevsaddam/gojsonq)](https://goreportcard.com/report/github.com/thedevsaddam/gojsonq)
[![Coverage Status](https://coveralls.io/repos/github/thedevsaddam/gojsonq/badge.svg?branch=master)](https://coveralls.io/github/thedevsaddam/gojsonq?branch=master)
[![GoDoc](https://godoc.org/github.com/thedevsaddam/gojsonq?status.svg)](https://godoc.org/github.com/thedevsaddam/gojsonq)
[![License: CC0-1.0](https://img.shields.io/badge/License-CC0%201.0-lightgrey.svg)](https://github.com/thedevsaddam/gojsonq/blob/dev/LICENSE.md)

A simple Go package to Query over JSON Data


### Installation

Install the package using
```go
$ go get github.com/thedevsaddam/gojsonq
```

### Usage

To use the package import it in your `*.go` code
```go
import "github.com/thedevsaddam/gojsonq"
```

Let's see a quick example:

```go
package main

import "github.com/thedevsaddam/gojsonq"

const json = `{"name":{"first":"Tom","last":"Hanks"},"age":61}`

func main() {
	name := gojsonq.New().JSONString(json).Find("name.first")
	println(name.(string)) // Tom
}

```

Another example:

```go
package main

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

const json = `{"city":"dhaka","type":"weekly","temperatures":[30,39.9,35.4,33.5,31.6,33.2,30.7]}`

func main() {
	avg := gojsonq.New().JSONString(json).From("temperatures").Avg()
	fmt.Println(avg) // 33.471428571428575
}

```

You can Query your data using the various query methods such as **[Find](https://github.com/thedevsaddam/gojsonq/wiki/Queries#findpath)**, **[First](https://github.com/thedevsaddam/gojsonq/wiki/Queries#first)**, **[Nth](https://github.com/thedevsaddam/gojsonq/wiki/Queries#nthindex)**, **[Pluck](https://github.com/thedevsaddam/gojsonq/wiki/Queries#pluckproperty)**,  **[Where](https://github.com/thedevsaddam/gojsonq/wiki/Queries#wherekey-op-val)**, **[OrWhere](https://github.com/thedevsaddam/gojsonq/wiki/Queries#orwherekey-op-val)**, **[WhereIn](https://github.com/thedevsaddam/gojsonq/wiki/Queries#whereinkey-val)**, **[WhereStartsWith](https://github.com/thedevsaddam/gojsonq/wiki/Queries#wherestartswithkey-val)**, **[WhereEndsWith](https://github.com/thedevsaddam/gojsonq/wiki/Queries#whereendswithkey-val)**, **[WhereContains](https://github.com/thedevsaddam/gojsonq/wiki/Queries#wherecontainskey-val)**, **[Sort](https://github.com/thedevsaddam/gojsonq/wiki/Queries#sortorder)**,  **[GroupBy](https://github.com/thedevsaddam/gojsonq/wiki/Queries#groupbyproperty)**,  **[SortBy](https://github.com/thedevsaddam/gojsonq/wiki/Queries#sortbyproperty-order)** and so on. Also you can aggregate your data after query using **[Avg](https://github.com/thedevsaddam/gojsonq/wiki/Queries#avgproperty)**,  **[Count](https://github.com/thedevsaddam/gojsonq/wiki/Queries#count)**, **[Max](https://github.com/thedevsaddam/gojsonq/wiki/Queries#maxproperty)**, **[Min](https://github.com/thedevsaddam/gojsonq/wiki/Queries#minproperty)**, **[Sum](https://github.com/thedevsaddam/gojsonq/wiki/Queries#sumproperty)** etc.

## Find more query API in [Wiki page](https://github.com/thedevsaddam/gojsonq/wiki/Queries)

## Bugs and Issues

If you encounter any bugs or issues, feel free to [open an issue at
github](https://github.com/thedevsaddam/gojsonq/issues).

Also, you can shoot me an email to
<mailto:thedevsaddam@gmail.com> for hugs or bugs.

## Credit

Special thanks to [Nahid Bin Azhar](https://github.com/nahid) for the inspiration and guidance for the package. Thanks to [Ahmed Shamim Hasan Shaon](https://github.com/me-shaon) for his support from the very beginning.

## Contributors
* [Lenin Hasda](https://github.com/leninhasda)
* [Sadlil Rhythom](https://github.com/sadlil)

## Contribution
If you are interested to make the package better please send pull requests or create an issue so that others can fix.
[Read the contribution guide here](CONTRIBUTING.md)

## See all [contributors](https://github.com/thedevsaddam/gojsonq/graphs/contributors)
