package main

import (
	exact "go/constant"
)

// Value represents a declared constant.
type Value struct {
	typ  string //type name
	name string // The name of the constant.
	// The value is stored as a bit pattern alone. The boolean tells us
	// whether to interpret it as an int64 or a uint64; the only place
	// this matters is when sorting.
	// Much of the time the str field is all we need; it is printed
	// by Value.String.
	value   uint64 // Will be converted to int64 when needed.
	signed  bool   // Whether the constant is a signed type.
	str     string // The string representation given by the "go/exact" package.
	comment string
	kind    exact.Kind
}

func (v *Value) String() string {
	return v.str
}

type ValueSlice []Value

//Returns: strings containing code switch
// case enum.Active:
//	return true
// case enum.InActive:
// return true

func (vs *ValueSlice) generateSwitchCode(pkgName string) string {
	result := ""
	for i, _ := range *vs {
		result += "\tcase " + pkgName + "." + (*vs)[i].name + ":\n"
		result += "\t\treturn true\n"
	}
	return result
}

func (vs *ValueSlice) generateMapValue(pkgName string) string {
	result := ""
	for i, _ := range *vs {
		result += "\t" + pkgName + "." + (*vs)[i].name + ": true,\n"
	}
	return result
}

func (vs *ValueSlice) generateMapKeyValue(pkgName string) string {
	result := ""
	for i, _ := range *vs {
		result += "\t" + pkgName + "." + (*vs)[i].name + ": \"" + (*vs)[i].comment + "\",\n"
	}
	return result
}

// byValue lets us sort the constants into increasing order.
// We take care in the Less method to sort in signed or unsigned order,
// as appropriate.
type byValue []Value

func (b byValue) Len() int      { return len(b) }
func (b byValue) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byValue) Less(i, j int) bool {
	if b[i].signed {
		return int64(b[i].value) < int64(b[j].value)
	}
	return b[i].value < b[j].value
}
