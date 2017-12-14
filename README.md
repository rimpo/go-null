# go-null

Go initializes struct object fields with default values (i.e. string as "", int  as 0, float64 as 0.0).
So, while json marshalling the defaults values are exported even though nothing was assigned in the field.
This tool generates nullable types using the custom type supplied to the command. 
