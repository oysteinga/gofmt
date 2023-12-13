package str

import (
	"fmt"
	"reflect"
)

func Sprint(a ...any) string {
	var newString string = fmt.Sprintln(fmt.Sprintln(a...))
	return newString[:len(newString)-2]
}

func SprintValueType(any any) string {

	var typeName string
	value := reflect.ValueOf(any)

	if value.Kind() == reflect.Pointer {
		typeName = "*" + reflect.TypeOf(any).Elem().Name()
		value = value.Elem()
	} else {
		typeName = reflect.TypeOf(any).Name()
	}

	if value.Kind() != reflect.Struct {
		typeName += ":"
	}

	return fmt.Sprintf(typeName+"%+v", value)
}

func PrettyCmplx(c complex128) string {
	if real(c) != 0 && imag(c) != 0 {
		return fmt.Sprint(c)
	} else if real(c) != 0 {
		return fmt.Sprint(real(c))
	} else if imag(c) != 0 {
		return fmt.Sprint(imag(c), "i")
	}

	return fmt.Sprint(0)
}

func Sprintln(a ...any) string {
	return fmt.Sprintln(a...)
}

func Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}
