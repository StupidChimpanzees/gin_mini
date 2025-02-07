package utils

func getType(data any) string {
	switch data.(type) {
	case bool:
		return "bool"
	case uint8:
		return "uint8"
	case uint16:
		return "uint16"
	case uint32:
		return "unit32"
	case uint64:
		return "unit64"
	case int8:
		return "int8"
	case int16:
		return "int16"
	case int32:
		return "int32"
	case int64:
		return "int64"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case complex64:
		return "complex64"
	case complex128:
		return "complex128"
	case string:
		return "string"
	case int:
		return "int"
	case uint:
		return "uint"
	case uintptr:
		return "uintptr"
	case any:
		return "any"
	case nil:
		return "nil"
	default:
		return "nil"
	}
}
