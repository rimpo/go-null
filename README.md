# go-null

Alternative way of handling struct and json in Go.

## Why?

Go initializes struct object fields with default values (i.e. string as "", int  as 0, float64 as 0.0).
In case, struct is used to define the json in the code, defaults values are exported even though nothing was assigned in the fields of the struct object.
This causes incorrect data exported.

## Available Options:

## How does go-null tool solves it?
go-null reads all the type in packages where the *go:generate go-null* is added and generate 2 package *null* and *jsonx*.
*null* will contains all the null type (which is thin wrapper like *sql.NullString*)created using the type read from package.
*jsonx* uses *null* package and knows how to Marshal the null types.

## 
