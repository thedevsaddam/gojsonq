package gojsonq

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

// ==================== Test Data===================
// ==================== DO NOT EDIT===================
var jsonStr = `
{
   "name":"computers",
   "description":"List of computer products",
   "vendor":{
      "name":"Star Trek",
      "email":"info@example.com",
      "website":"www.example.com",
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
            "id":5,
            "name":"HP core i5",
            "price":850,
            "key": 2300
         },
         {
            "id":6,
            "name":"HP core i7",
            "price":950
         },
         {
            "id":null,
            "name":"HP core i3 SSD",
            "price":850
         }
      ],
      "prices":[
         2400,
         2100,
         1200,
         400.87,
         89.90,
         150.10
     ],
     "names":[
        "John Doe",
        "Jane Doe",
        "Tom",
        "Jerry",
        "Nicolas",
        "Abby"
     ]
   }
}
`

//================= Test Helpers===========================

func createTestFile(t *testing.T, filename string) func() {
	// create data.json file from the jsonStr above
	if err := ioutil.WriteFile(filename, []byte(jsonStr), 0644); err != nil {
		t.Errorf("failed to create %s test file %v", filename, err)
	}

	return func() {
		if err := os.Remove(filename); err != nil {
			t.Errorf("failed to remove %s test file %v", filename, err)
		}
	}
}

func assertJSON(t *testing.T, v interface{}, expJSON string) {
	bb, err := json.Marshal(v)
	if err != nil {
		t.Errorf("failed to marshal: %v", err)
	}
	eb := []byte(expJSON)
	if !bytes.Equal(bb, eb) {
		t.Errorf("Expected: %v\nGot: %v", expJSON, string(bb))
	}
}
