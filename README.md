#### Note: Command should work. Documentation work is going on.

# go-null

Alternative way of handling struct and json in Go.


## Why?

Go initializes struct object fields with default values (i.e. string as "", int  as 0, float64 as 0.0).
In case, struct is used to define the json in the code, defaults values are exported even though nothing was assigned in the fields of the struct object.
This causes incorrect data exported.


## Available Options (Go way):

1. Converting field to pointer:

#### Cons:

- The code will contains lot of * and & if you use this JSON struct in your buisness logic.
- Need to take care of allocation of the values and good understanding of pointers



## How does go-null tool solves it?

go-null helps you by generating very thin wrappers around enum and custom types (i.e. null package).
The generated *null* package types has IsNull and SetSafe function to support null check and enum value check.
All you have to to is now to add the null types in your json struct declaration.
It also generates *jsonx* package which supports marshalling of null types (generated using the go-null command) and you can use *omitempty* to drop null values.



## Usage

```
go-null -package=<package path> -output=<generated code path>


#Add below line on top of the enums and custom type packages for which you want to generate null types.
#Note: choose any one file of the package and write once
//go:generate go-null -package=github.com/rimpo/go-null/examples/examples1/ -output=..

#null and jsonx package will be generated in path github.com/rimpo/go-null/examples/examples1/
```


