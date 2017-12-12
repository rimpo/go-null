package main

const stringMainBody = `
type {{.TypeName}} struct {
	val {{.BuiltInTypeName}}
	valid bool
}

func New{{.TypeName}}(val {{.BuiltInTypeName}}) {{.TypeName}} {
	return {{.TypeName}}{val: val, valid:true }
}

func (t *{{.TypeName}}) Set(val {{.BuiltInTypeName}}) {
	t.val = val
	t.valid = true
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *{{.TypeName}}) SetSafe(val {{.BuiltInTypeName}}) {
	if IsValueSafe{{.TypeName}}(val) {
		log.Printf("ERROR: Unknown value:%v assigned to type:{{.TypeName}}!!.\n%v", val, debug.PrintStack())
	}
	t.val = val
	t.assigned = true
}

//Logs error message
func (t *{{.TypeName}}) Get() {{.BuiltInTypeName}} {
	if !t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:{{.TypeName}}!!.\n%v", debug.PrintStack())
	}
	return t.val
}

func (t *{{.TypeName}}) GetPtr() *{{.BuiltInTypeName}} {
	return &t.val
}

func (t *{{.TypeName}}) IsNull() bool {
	return !t.valid
}
func (t *{{.TypeName}}) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}

func IsValueSafe{{.TypeName}}(t {{.BuiltInTypeName}}) bool {
	{{.AllConstSwitch}}
}
`
