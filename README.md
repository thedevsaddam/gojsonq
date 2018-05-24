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
(float64) 6340.87
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
* [Sum](#sumproperty)
* [Count](#count)
* [Max](#maxproperty)
* [Min](#minproperty)
* [Avg](#avgproperty)
* [First](#first)
* [Last](#last)
* [Nth](#nthindex)
* [GroupBy](#groupbyproperty)
* [Sort](#sortorder)
* [SortBy](#sortbyproperty-order)
* [Reset](#reset)
* [Only](#onlyproperties)
* [Pluck](#pluckproperty)
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

    * `startsWith` : Check if the value of given **key** in data starts with (has a prefix of) the given **val**. This would only works for _String_ type data and exact match.

    * `endsWith` : Check if the value of given **key** in data ends with (has a suffix of) the given **val**. This would only works for _String_ type data and exact match.

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

### `Sum(property)`

* `property` -- the property name of the data.

**example:**

Let's say you want to find the sum of the _'price'_ of the _'items'_. You can do it like this:

```go
jq := gojsonq.New().File("./data.json").From("items")
fmt.Printf("%#v\n", jq.Sum("price"))
```

If the data you are aggregating slice of `int/float`, you don't need to pass the 'property' parameter.
See example below:

```go
jq := gojsonq.New().File("./data.json").From("prices")
fmt.Printf("%#v\n", jq.Sum())
```

### `Count()`

It will return the number of elements in the collection/object.

**example:**

Let's say you want to find how many elements are in the _'items'_ property. You can do it like:

```go
jq := gojsonq.New().File("./data.json").From("items")
fmt.Printf("%#v\n", jq.Count())

// or count properties of an object
jq := gojsonq.New().File("./data.json").From("items.[0]")
fmt.Printf("%#v\n", jq.Count())
```

### `Max(property)`

* `property` -- the property name of the data

**example:**

Let's say you want to find the maximum of the _'price'_ of the _'items'_. You can do it like this:

```go
jq := gojsonq.New().File("./data.json").From("items")
fmt.Printf("%#v\n", jq.Max("price"))
```

If the data you are querying a slice of `int/float`, you don't need to pass the 'property' parameter.
See example below:

```go
jq := gojsonq.New().File("./data.json").From("prices")
fmt.Printf("%#v\n", jq.Max())
```

### `Min(property)`

* `property` -- the property name of the data

**example:**

Let's say you want to find the minimum of the _'price'_ of the _'items'_. You can do it like this:

```go
jq := gojsonq.New().File("./data.json").From("items")
fmt.Printf("%#v\n", jq.Min("price"))
```

If the data you are querying a slice of `int/float`, you don't need to pass the 'property' parameter.
See detail example:

```go
jq := gojsonq.New().File("./data.json").From("prices")
fmt.Printf("%#v\n", jq.Min())
```

### `Avg(property)`

* `property` -- the property name of the data

**example:**

Let's say you want to find the average of the _'price'_ of the _'items'_. You can do it like this:

```go
jq := gojsonq.New().File("./data.json").From("items")
fmt.Printf("%#v\n", jq.Avg("price"))
```

If the data you are querying a slice of `int/float`, you don't need to pass the 'property' parameter.
See detail example:

```go
jq := gojsonq.New().File("./data.json").From("prices")
fmt.Printf("%#v\n", jq.Avg())
```

### `First()`

It will return the first element of the collection.

**example:**

```go
jq := gojsonq.New().File("./data.json").From("items")
fmt.Printf("%#v\n", jq.First())
```

### `Last()`

It will return the last element of the collection.

**example:**

```go
jq := gojsonq.New().File("./data.json").From("prices")
fmt.Printf("%#v\n", jq.Last())
```

### `Nth(index)`

* `index` -- index of the element to be returned.

It will return the nth element of the collection. If the given index is a **positive** value, it will return the nth element from the beginning. If the given index is a **negative** value, it will return the nth element from the end.

**example:**

```go
jq := gojsonq.New().File("./data.json").From("items")
fmt.Printf("%#v\n", jq.Nth(2))
```

### `GroupBy(property)`

* `property` -- The property by which you want to group the collection.

**example:**

Let's say you want to group the _'items'_ data based on the _'price'_ property. You can do it like:


```go
jq := gojsonq.New().File("./data.json").From("items")
fmt.Printf("%#v\n", jq.GroupBy("price").Get())
```

### `Sort(order)`

* `order` -- If you skip the _'order'_ property the data will be by default ordered as **ascending**. You need to pass **'desc'** as the _'order'_ parameter to sort the data in **descending** order.

**Note:** This method should be used for `Slice`. If you want to sort an Array of Objects you should use the **SortBy()** method described later.

**example:**

Let's say you want to sort the _'prices/names'_ data. You can do it like:


```go
jq := gojsonq.New().File("./data.json").From("prices")
fmt.Printf("%#v\n", jq.Sort().Get())

// or sort array of strings in descending order
jq := gojsonq.New().File("./data.json").From("names")
fmt.Printf("%#v\n", jq.Sort("desc").Get())
```

### `SortBy(property, order)`

* `property` -- You need to pass the property name on which the sorting will be done.
* `order` -- If you skip the _'order'_ property the data will be by default ordered as **ascending**. You need to pass **'desc'** as the _'order'_ parameter to sort the data in **descending** order.

**Note:** This method should be used for Array of Objects. If you want to sort a plain Array you should use the **Sort()** method described earlier.

**example:**

Let's say you want to sort the _'price'_ data of _'items'_. You can do it like:

```go
jq := gojsonq.New().File("./data.json").From("items")
fmt.Printf("%#v\n", jq.SortBy("price").Get())

// or in descending order
jq := gojsonq.New().File("./data.json").From("items")
fmt.Printf("%#v\n", jq.SortBy("price", "desc").Get())
```

### `Reset()`

Reset the queries with the original data so that you can query again.
See the example below:

```go
jq := gojsonq.New().File("./data.json").From("items")

res1 := jq.Where("price", ">", 900).Sum("price")

// got our first result, now reset the instance and query again
res2 := jq.Reset().SortBy("price").Get()
fmt.Printf("Res1: %#v\nRes2: %#v\n", res1, res2)
```

### `Only(properties)`

* `property` -- The property by which you want to get in final results. To get a clear idea see the example below:

**Example**

```go
jq := gojsonq.New().File("./data.json").From("items").WhereNotNil("id")
fmt.Printf("%#v\n", jq.Only("id", "price").Get())
```

**Output**
```go
[]interface {}{
    map[string]interface {}{"id":1, "price":1350},
    map[string]interface {}{"id":2, "price":1700},
    map[string]interface {}{"id":3, "price":1200},
    map[string]interface {}{"id":4, "price":850},
}
```

### `Pluck(property)`

* `property` -- The property by which you want to get an array.

Only returns a plain array of values for the property, you can chain further method to `Pluck` like `Sort/Min/Max/Avg/Count` etc . To get a clear idea see the example below:

**Example**
```go
jq := gojsonq.New().File("./data.json").From("items")
fmt.Printf("%#v\n", jq.Pluck("price").Get())
```

**Output**
```go
[]interface {}{1350, 1700, 1200, 850, 850}
```

### `Macro(operator, QueryFunc)`

Query matcher can be written as macro and used multiple time for further queries. Lets' say we don't have weak match for `WhereStartsWith`, we can write one. See the example below:

```go
jq := gojsonq.New().File("./data.json").From("items")
jq.Macro("WM", func(x, y interface{}) (bool, error) { // WM is our weak match operator
    xs, okx := x.(string)
    ys, oky := y.(string)
    if !okx || !oky {
        return false, fmt.Errorf("weak match only support string")
    }

    return strings.HasPrefix(strings.ToLower(xs), strings.ToLower(ys)), nil
})
jq.Where("name", "WM", "mac")
fmt.Printf("%#v\n", jq.Get())
```

### `Copy()`

It will return a complete clone of the Object instance. Note `Copy` method is very useful when working concurrently. You can copy the instance for multiple `goroutine`

**Example**

```go
jq := gojsonq.New().File("./data.json")
for i := 0; i < 10; i++ {
    go func(j *gojsonq.JSONQ) {
        fmt.Printf("Sum: %#v\n", j.From("items").Sum("price"))
    }(jq.Copy())

    go func(j *gojsonq.JSONQ) {
        fmt.Printf("Min: %#v\n", j.From("items").Min("price"))
    }(jq.Copy())
}
time.Sleep(time.Second)
```

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
