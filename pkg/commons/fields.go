package commons

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
)

func FilterFieldsWithValues(originalMap map[string]any) map[string]any {
	filteredMap := make(map[string]any)

	for k, v := range originalMap {
		if v == nil || IsEmptyValue(reflect.ValueOf(v)) {
			continue
		}

		filteredMap[k] = v
	}

	return filteredMap
}

func CloneFieldMap(fields map[string]any) map[string]any {
	return FilterFieldsByName(fields)
}

func FilterFieldsByName(fields map[string]any, field ...string) map[string]any {
	if fields == nil {
		return make(map[string]any)
	}

	filteredFields := make(map[string]any)

	for name, value := range fields {
		if Contains(field, name) {
			continue
		}

		filteredFields[name] = value
	}

	return filteredFields
}

func Contains(slice []string, target string) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}

	return false
}

func GetField(name string, fields map[string]any) any {
	if v, ok := fields[name]; ok {
		return v
	}

	return nil
}

func GetFieldAsStringOrElse(name string, fields map[string]any, orElse string) string {
	return GetAsStringOrElse(GetField(name, fields), orElse)
}

func GetFieldAsIntOrElse(name string, fields map[string]any, orElse int) int {
	return GetAsIntOrElse(GetField(name, fields), orElse)
}

func GetFieldAsBoolOrElse(name string, fields map[string]any, orElse bool) bool {
	return GetAsBoolOrElse(GetField(name, fields), orElse)
}

func GetFieldAsStringMapOrElse(name string, fields map[string]any, orElse map[string]string) map[string]string {
	return GetAsStringMapOrElse(GetField(name, fields), orElse)
}

func GetFieldAsTimeOrElse(name string, fields map[string]any, orElse time.Time) time.Time {
	return *GetAsTimePtrOrElse(GetField(name, fields), &orElse)
}

func GetFieldAsTimePtrOrElse(name string, fields map[string]any, orElse *time.Time) *time.Time {
	return GetAsTimePtrOrElse(GetField(name, fields), orElse)
}

func GetFieldAsRequestPtrOrElse(name string, fields map[string]any, orElse *http.Request) *http.Request {
	return GetAsRequestPtrOrElse(GetField(name, fields), orElse)
}

func GetFieldAsResponsePtrOrElse(name string, fields map[string]any, orElse *http.Response) *http.Response {
	return GetAsResponsePtrOrElse(GetField(name, fields), orElse)
}

func GetAsRequestPtrOrElse(value any, orElse *http.Request) *http.Request {
	switch v := value.(type) {
	case http.Request:
		return &v
	case *http.Request:
		return v
	}

	return orElse
}

func GetAsResponsePtrOrElse(value any, orElse *http.Response) *http.Response {
	switch v := value.(type) {
	case http.Response:
		return &v
	case *http.Response:
		return v
	}

	return orElse
}

func GetAsTimePtrOrElse(value any, orElse *time.Time) *time.Time {
	switch v := value.(type) {
	case time.Time:
		return &v
	case *time.Time:
		return v
	}

	return orElse
}

func GetAsStringMapOrElse(value any, orElse map[string]string) map[string]string {
	switch v := value.(type) {
	case map[string]string:
		return v
	case *map[string]string:
		return *v
	}

	return orElse
}

func GetAsBoolOrElse(value any, orElse bool) bool {
	switch v := value.(type) {
	case bool:
		return v
	case *bool:
		return *v
	}

	return orElse
}

func GetAsStringOrElse(value any, orElse string) string {
	switch v := value.(type) {
	case string:
		return v
	case *string:
		return *v
	case int, int64, uint, uint64, int8, int16, int32, uint8, uint16, uint32,
		*int, *int64, *uint, *uint64, *int8, *int16, *int32, *uint8, *uint16, *uint32,
		float64, float32, *float64, *float32:
		return fmt.Sprintf("%v", v)
	}

	return orElse
}

func GetAsFloat64OrElse(value any, orElse float64) float64 {
	switch v := value.(type) {
	case int:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint64:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case *int:
		return float64(*v)
	case *int64:
		return float64(*v)
	case *uint:
		return float64(*v)
	case *uint64:
		return float64(*v)
	case *int8:
		return float64(*v)
	case *int16:
		return float64(*v)
	case *int32:
		return float64(*v)
	case *uint8:
		return float64(*v)
	case *uint16:
		return float64(*v)
	case *uint32:
		return float64(*v)
	case float64:
		return v
	case float32:
		return float64(v)
	case *float64:
		return *v
	case *float32:
		return float64(*v)
	}

	return orElse
}

func GetAsIntOrElse(value any, orElse int) int {
	return int(GetAsFloat64OrElse(value, float64(orElse)))
}
