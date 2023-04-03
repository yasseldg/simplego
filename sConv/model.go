package sConv

type ValueType struct {
	Text string
	Html string
}

var ValueTypes = map[string]ValueType{
	"bool":   {"Boolean", `type="checkbox" value="0"`},
	"float":  {"Float", `type="number" step="0.01" min="0.00" value="0.00"`},
	"int":    {"Integer", `type="number" min="0" value="0"`},
	"string": {"String", `type="text" value=""`},
}

type ExtValues struct {
	V_1  string
	V_2  string
	step string
}
