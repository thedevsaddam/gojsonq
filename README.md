gojsonq
===============
[![Build Status](https://travis-ci.org/thedevsaddam/gojsonq.svg?branch=master)](https://travis-ci.org/thedevsaddam/gojsonq)
[![Project status](https://img.shields.io/badge/version-beta-yellow.svg)](https://github.com/thedevsaddam/gojsonq/releases)
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


You can Query your data using the various query methods such as **Find**, **Where**, **OrWhere**, **WhereIn**, **WhereStartsWith**, **WhereEndsWith**, **WhereContains** and so on. Also you can aggregate your data after query using **Sum**, **Count**, **GroupBy**, **Max**, **Min** etc.

Let's see a quick example:

<details><summary>Sample data (data.json)</summary>
<pre>
{
   "name":"computers",
   "description":"List of computer products",
   "prices":[2400, 2100, 1200, 400.87, 89.90, 150.10],
   "names":["John Doe", "Jane Doe", "Tom", "Jerry", "Nicolas", "Abby"],
   "items":[
      {
         "id":1,
         "name":"MacBook Pro 13 inch retina",
         "price":1350
      },
      {
         "id":2,
         "name":"MacBook Pro 15 inch retina",
         "price":1700
      },
      {
         "id":3,
         "name":"Sony VAIO",
         "price":1200
      },
      {
         "id":4,
         "name":"Fujitsu",
         "price":850
      },
      {
         "id":null,
         "name":"HP core i3 SSD",
         "price":850
      }
   ]
}
</pre>
</details>

**Example:**
```go
jq := gojsonq.New().
    File("./data.json").
    From("items").
    Where("price", ">", 1200)
fmt.Printf("%#v\n", jq.Get())
```
**Output:**
```go
[]interface {}{
    map[string]interface {}{"id":1, "name":"MacBook Pro 13 inch retina", "price":1350},
    map[string]interface {}{"id":2, "name":"MacBook Pro 15 inch retina", "price":1700},
}
```

Let's say we want to get the Summation of price of the Queried result. We can do it easily by calling the Sum() method instead of Get():

**Example**
```go
jq := gojsonq.New().
    File("./data.json").
    From("prices")
fmt.Printf("%#v\n", jq.Sum())
```

**Output**
```json
(float64) 3050
```


Let's explore the full API to see what else magic this library can do for you.
Shall we?

## API

Following API examples are shown based on the sample JSON data given above. To get a better idea of the examples see that JSON data first.

**List of API:**

* [File](#filepath)
* [JSONString](#jsonstringjson)
* [Reader](#readerioreader)
* [Get](#get)
* [Find](#findpath)
* [From](#frompath)
* [Where](#wherekey-op-val)
* [OrWhere](#orwherekey-op-val)
* [WhereIn](#whereinkey-val)
* [WhereNotIn](#wherenotinkey-val)
* [WhereNil](#wherenilkey)
* [WhereNotNil](#wherenotnilkey)
* [WhereEqual](#whereequalkey-val)
* [WhereNotEqual](#wherenotequalkey-val)
* [WhereStartsWith](#wherestartswithkey-val)
* [WhereEndsWith](#whereendswithkey-val)
* [WhereContains](#wherecontainskey-val)
* [WhereStrictContains](#wherestrictcontainskey-val)
* [Sum](#sum)
* [Count](#count)
* [Max](#max)
* [Min](#min)
* [Avg](#avg)
* [First](#first)
* [Last](#last)
* [Nth](#nth)
* [GroupBy](#groupby)
* [Sort](#sort)
* [SortBy](#sortby)
* [Reset](#reset)
* [Only](#only)
* [Pluck](#pluck)
* [Macro](#macrooperator-queryfunc)
* [Copy](#copy)

### `File(path)`

This method takes a JSON file path as argument for further queries.

```go
res := gojsonq.New().File("./data.json").From("items").Get()
fmt.Printf("%#v\n", res)
```

### `JSONString(json)`

This method takes a valid JSON string as argument for further queries.

```go
res := gojsonq.New().JSONString("[19, 90.9, 7, 67.5]").Sum()
fmt.Printf("%#v\n", res)
```
### `Reader(io.Reader)`

This method takes an `io.Reader` as argument to read JSON data for further queries.

```go
strReader := strings.NewReader("[19, 90.9, 7, 67.5]")
res := gojsonq.New().Reader(strReader).Avg()
fmt.Printf("%#v\n", res)
```

### `Get()`

This method will execute queries and will return the resulted data. You need to call it finally after using some query methods. [See usage in the above example](#filepath)

### `Find(path)`

* `path` -- the path hierarchy of the data you want to find.

You don't need to call `Get()` method after this. Because this method will fetch and return the data by itself.

**caveat:** You can't chain further query methods after it. If you need that, you should use `From()` method.

**example:**

Let's say you want to get the value of _'items'_ property of your JSON Data. You can do it like this:

```go
items := gojsonq.New().File("./data.json").Find("vendor.items");
fmt.Printf("%#v\n", items)
```

If you want to traverse to more deep in hierarchy, you can do it like:

```go
item := gojsonq.New().File("./data.json").Find("vendor.items.[0]");
fmt.Printf("%#v\n", item)
```

### `From(path)`

* `path` (optional) -- the path hierarchy of the data you want to start query from.

By default, query would be started from the root of the JSON Data you've given. If you want to first move to a nested path hierarchy of the data from where you want to start your query, you would use this method. Skipping the `path` parameter or giving **'.'** as parameter will also start query from the root Data.

Difference between this method and `Find()` is that `Find()` method will return the data from the given path hierarchy. On the other hand, this method will return the Object instance, so that you can further chain query methods after it.

**Example:**

Let's say you want to start query over the values of _'items'_ property of your JSON Data. You can do it like this:

```go
jq := gojsonq.New().File("./data.json").From("items").Where("price", ">", 1200)
fmt.Printf("%#v\n", jq.Get())
```

If you want to traverse to more deep in hierarchy, you can do it like:

```go
jq := gojsonq.New().File("./data.json").From("vendor.items").Where("price", ">", 1200)
fmt.Printf("%#v\n", jq.Get())
```

### `where(key, op, val)`

* `key` -- the property name of the data.

* `val` -- value to be matched with. It can be a _int_, _string_, _bool_ or _slice_ - depending on the `op`.

* `op` -- operand to be used for matching. The following operands are available to use:

    * `=` : For equality matching

    * `eq` : Same as `=`

    * `!=` : For 4not equality matching

    * `neq` : Same as `!=`

    * `<>` : Same as `!=`

    * `>` : Check if value of given **key** in data is Greater than **val**
    * `gt` : Same as `>`

    * `<` : Check if value of given **key** in data is Less than **val**

    * `lt` : Same as `<`

    * `>=` : Check if value of given **key** in data is Greater than or Equal of **val**

    * `gte` : Same as `>=`

    * `<=` : Check if value of given **key** in data is Less than or Equal of **val**

    * `lte` : Same as `<=`

    * `in` : Check if the value of given **key** in data is exists in given **val**. **val** should be a _Slice_ of `int/float64/string`.

    * `notIn` : Check if the value of given **key** in data is not exists in given **val**. **val** should be a _Slice_ of `int/float64/string`.

    * `startsWith` : Check if the value of given **key** in data starts with (has a prefix of) the given **val**. This would only works for _String_ type data.

    * `endsWith` : Check if the value of given **key** in data ends with (has a suffix of) the given **val**. This would only works for _String_ type data.

    * `contains` : Check if the value of given **key** in data has a substring of given **val**. This would only works for _String_ type data and loose match.

    * `strictContains` : Check if the value of given **key** in data has a substring of given **val**. This would only works for _String_ type data and exact match.

**example:**

Let's say you want to find the _'items'_ who has _price_ greater than `1200`. You can do it like this:

```go
jq := gojsonq.New().File("./data.json").From("items").Where("price", ">", 1200)
fmt.Printf("%#v\n", jq.Get())
```

You can add multiple _where_ conditions. It'll give the result by AND-ing between these multiple where conditions.

```go
jq := gojsonq.New().File("./data.json").From("items").Where("price", ">", 500).Where("name","=", "Fujitsu")
fmt.Printf("%#v\n", jq.Get())
```
### `OrWhere(key, op, val)`

Parameters of `OrWhere()` are the same as `Where()`. The only difference between `Where()` and `OrWhere()` is: condition given by the `OrWhere()` method will OR-ed the result with other conditions.

For example, if you want to find the _'items'_ with _id_ of `1` or `2`, you can do it like this:

```go
jq := gojsonq.New().File("./data.json").From("items").Where("id", "=", 1).OrWhere("id", "=", 2)
fmt.Printf("%#v\n", jq.Get())
```

### `WhereIn(key, val)`

* `key` -- the property name of the data
* `val` -- it should be a **Slice** of `int/float64/string`

This method will behave like `where(key, 'in', val)` method call.

### `WhereNotIn(key, val)`

* `key` -- the property name of the data
* `val` -- it should be a **Slice** of `int/float64/string`


This method will behave like `Where(key, 'notIn', val)` method call.

### `WhereNil(key)`

* `key` -- the property name of the data

This method will behave like `Where(key, "=", nil)` method call.

### `WhereNotNil(key)`

* `key` -- the property name of the data

This method will behave like `Where(key, '!=', nil)` method call.

### `WhereEqual(key, val)`

* `key` -- the property name of the data
* `val` -- it should be a `int/float64/string`

This method will behave like `Where(key, '=', val)` method call.

### `WhereNotEqual(key, val)`

* `key` -- the property name of the data
* `val` -- it should be a `int/float64/string`

This method will behave like `Where(key, '!=', val)` method call.

### `WhereStartsWith(key, val)`

* `key` -- the property name of the data
* `val` -- it should be a String

This method will behave like `Where(key, 'startsWith', val)` method call.

### `WhereEndsWith(key, val)`

* `key` -- the property name of the data
* `val` -- it should be a String

This method will behave like `where(key, 'endsWith', val)` method call.

### `WhereContains(key, val)`

* `key` -- the property name of the data
* `val` -- it should be a String

This method will behave like `Where(key, 'contains', val)` method call.

### `WhereStrictContains(key, val)`

* `key` -- the property name of the data
* `val` -- it should be a String

This method will behave like `Where(key, 'strictContains', val)` method call.


## TODO

- [ ] Add missing methods
- [ ] Write full documentation with example

## Bugs and Issues

If you encounter any bugs or issues, feel free to [open an issue at
github](https://github.com/thedevsaddam/gojsonq/issues).

Also, you can shoot me an email to
<mailto:thedevsaddam@gmail.com> for hugs or bugs.

## Credit

Special thanks to [Nahid Bin Azhar](https://github.com/nahid) for the inspiration and guidance for the package.

## Contribution
If you are interested to make the package better please send pull requests or create an issue so that others can fix.
[Read the contribution guide here](CONTRIBUTING.md)

## See all [contributors](https://github.com/thedevsaddam/gojsonq/graphs/contributors)
