package fet

import "reflect"

func getKeyFromField(field reflect.StructField, name string) string {

	if key := field.Tag.Get("bson"); key != "" {
		return key
	}

	return name
}
