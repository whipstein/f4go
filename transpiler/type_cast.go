package transpiler

import "fmt"

func castToGoType(fortranType string) (goType string, err error) {
	switch fortranType {
	case "integer(kind=4)", "integer(kind=8)":
		goType = "int"
	case "real(kind=4)", "real(kind=8)":
		goType = "float64"
	case "character(kind=1)":
		goType = "byte"
	default:
		fmt.Printf("Cannot CastToGoType: %v\n", fortranType)
		goType = fortranType
	}
	return
}
