[![Go Report Card](https://goreportcard.com/badge/github.com/rimpo/go-null)](https://goreportcard.com/report/github.com/rimpo/go-null)

#### ***************************************
####            Work Under Progress
#### ***************************************

# go-null

go-null is a code generator for go. Alternative way of defining struct for representing json.

## Usage

In the package containing custom enums and custom type which is going get used in the json struct,
add below line in one of the source code file inside the package.
```
go-null -package=<package path> -output=<generated code path>

```

example1: [enums.go](https://github.com/rimpo/go-null/blob/master/examples/example1/enum/enums.go)
```
//go:generate go-null -package=github.com/rimpo/go-null/examples/examples1/ -output=..

#run
github.com/rimpo/go-null/examples/examples1/>go generate ./...
```
null and jsonx package will be auto-generated in [path](github.com/rimpo/go-null/examples/examples1/)

generated source code of null types:
[typememberstatus.go](https://github.com/rimpo/go-null/blob/master/examples/example1/null/typememberstatus.go)
[typephoneprivacy.go](https://github.com/rimpo/go-null/blob/master/examples/example1/null/typephoneprivacy.go)

null - contains all the wrapped types
json - package for marshalling null types in json

## Why?

Go initializes struct object fields with default values (i.e. string as "", int  as 0, float64 as 0.0).
In case, struct is used to define the json in the code, defaults values are exported even though nothing was assigned in the fields of the struct object.
This causes incorrect data exported.

Example (snippet):
```
package main

import (
  "fmt"
  "encoding/json"
)

type TypeOrderCode int 
const (
  OrderRequest TypeOrderCode = iota
  OrderResponse
  OrderTrade
)
type TypeOrderRequest struct {
  OrderCode TypeOrderCode `json:"order_code"`
  ScripCode string `json:"scrip_code"`
  Member string `json:"member"`
  Qty int `json:"qty"`
  Price int `json:"price"`
  MarketOrder bool `json:"market_order"`
}

func main(){
  var req TypeOrderRequest
  
  //Business Logic - start
  req.Qty = 100
  req.Member = "Rimpo"
  //Business Logic - end
  
 data, _ := json.Marshal(req)
 fmt.Println(string(data))
}
```
Output:
```
{order_code:0, scrip_code:"", member:"Rimpo", qty:0, price:0, market_order:false}
```
order_code, scrip_code, price & market_order is present in the output json (with default values) even though nothing was assigned to them.

Go json package supports *omitempty* in the tag but this will completely remove the actualy values i.e.
if TypeOrderRequest is declared in this way and
```
type TypeOrderRequest struct {
  OrderCode TypeOrderCode `json:"order_code"`
  ScripCode string `json:"scrip_code,omitempty"`
  Member string `json:"member,omitempty"`
  Qty int `json:"qty"`
  Price int `json:"price"`
  MarketOrder bool `json:"market_order,omitempty"`
}

//Business Logic - start
 req.Qty = 100
req.Member = "Rimpo"
req.MarketOrder = false
//Business Logic - end
```
Output:
```
{order_code:0, member:"Rimpo", qty:0, price:0}
```
*omitempty* causes default values not getting marshalled in json.Marshal.
MarketOrder was assigned false value still it didn't get into the output json.


## Available Option:

### Converting field to [pointer](https://stackoverflow.com/questions/18088294/how-to-not-marshal-an-empty-struct-into-json-with-go):
Problem with this approach 
- The code will contains lot of * and & if you use this JSON struct in your buisness logic.
- Need to take care of allocation of the values and good understanding of pointers



## How does go-null tool solves it?

go-null helps you by generating very thin wrappers around enum and custom types (i.e. null package).
The generated *null* package types has IsNull and SetSafe function to support null check and enum value check.
All you have to to is now to add the null types in your json struct declaration.
It also generates *jsonx* package which supports marshalling of null types (generated using the go-null command) and you can use *omitempty* to drop null values.




