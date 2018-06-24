package gojsonq_test

import (
	"fmt"
	"os"

	"github.com/D1CED/gojsonq"
)

func ExampleJSONQ_Find() {
	var json = `{"name":{"first":"Tom","last":"Hanks"},"age":61}`
	name := gojsonq.New().JSONString(json).Find("name.first")
	fmt.Println(name)
	// Output: Tom
}

func Example() {
	var json = `{"city":"dhaka","type":"weekly","temperatures":[30,39.9,35.4,33.5,31.6,33.2,30.7]}`
	avg := gojsonq.New().JSONString(json).From("temperatures").Avg()
	fmt.Println(avg)
	// Output: 33.471428571428575
}

func ExampleJSONQ_Avg() {
	var json = `{"city":"dhaka","type":"weekly","temperatures":[30,39.9,35.4,33.5,31.6,33.2,30.7]}`
	avg := gojsonq.New().JSONString(json).From("temperatures").Avg()
	fmt.Println(avg)
	// Output: 33.471428571428575
}

func ExampleJSONQ_File() {
	jq := gojsonq.New().File("example.json")
	if jq.Error() != nil {
		fmt.Println(jq.Errors())
		os.Exit(1)
	}

	res := jq.Find("vendor.items.[1].name")
	if jq.Error() != nil {
		fmt.Println(jq.Errors())
		os.Exit(1)
	}

	fmt.Println(res)
	// Output: MacBook Pro 15 inch retina
}

func ExampleJSONQ_Sum() {
	jq := gojsonq.New().File("./example.json")
	res := jq.From("vendor.items").Where("price", ">", 1200).OrWhere("id", "=", nil).Sum("price")
	fmt.Println(res)
	// Output: 3900
}

func ExampleJSONQ_Pluck() {
	jq := gojsonq.New().File("./example.json")
	res := jq.From("vendor.items").Where("price", ">", 1200).Pluck("price")
	fmt.Println(res)
	// Output: [1350 1700]
}
