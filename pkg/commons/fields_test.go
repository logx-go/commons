package commons

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestFilterEmptyValues(t *testing.T) {
	originalMap := map[string]any{
		"key1": "value1",
		"key2": nil,
		"key3": "",
		"key4": 0,
		"key5": []string{},
		"key6": map[string]any{},
		"key7": true,
	}

	filteredMap := FilterFieldsWithValues(originalMap)

	expectedMap := map[string]any{
		"key1": "value1",
		"key7": true,
	}

	if !reflect.DeepEqual(filteredMap, expectedMap) {
		t.Errorf("Filtered map does not match expected result. Got: %v, Want: %v", filteredMap, expectedMap)
	}
}

func TestFilterFieldsByName(t *testing.T) {
	testCases := []struct {
		Name           string
		Fields         map[string]interface{}
		Filter         []string
		ExpectedResult map[string]interface{}
	}{
		{
			Name: "FilterFieldsByName with fields",
			Fields: map[string]interface{}{
				"name":     "John",
				"age":      30,
				"location": "New York",
				"email":    "john@example.com",
			},
			Filter: []string{"age", "email"},
			ExpectedResult: map[string]interface{}{
				"name":     "John",
				"location": "New York",
			},
		},
		{
			Name:           "FilterFieldsByName with nil fields",
			Fields:         nil,
			Filter:         []string{"age", "email"},
			ExpectedResult: map[string]interface{}{},
		},
		{
			Name:           "FilterFieldsByName with empty fields",
			Fields:         map[string]interface{}{},
			Filter:         []string{"age", "email"},
			ExpectedResult: map[string]interface{}{},
		},
	}

	for _, tc := range testCases {
		tc := tc // Capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			filteredFields := FilterFieldsByName(tc.Fields, tc.Filter...)

			if !reflect.DeepEqual(filteredFields, tc.ExpectedResult) {
				t.Errorf("FilterFieldsByName() = %v, want %v", filteredFields, tc.ExpectedResult)
			}
		})
	}
}

func TestGetAsResponsePtrOrElse(t *testing.T) {
	resp := &http.Response{StatusCode: 200}
	testCases := []struct {
		name   string
		value  interface{}
		orElse *http.Response
		want   *http.Response
	}{
		{
			name:   "http.Response",
			value:  *resp,
			orElse: nil,
			want:   resp,
		},
		{
			name:   "*http.Response",
			value:  resp,
			orElse: nil,
			want:   resp,
		},
		{
			name:   "unsupported type",
			value:  42,
			orElse: resp,
			want:   resp,
		},
		{
			name:   "unsupported type with orElse nil",
			value:  "not a response",
			orElse: nil,
			want:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetAsResponsePtrOrElse(tc.value, tc.orElse) //nolint:bodyclose
			if (got == nil && tc.want != nil) || (got != nil && tc.want == nil) || (got != nil && tc.want != nil && !reflect.DeepEqual(got, tc.want)) {
				t.Errorf("GetAsResponsePtrOrElse(%v, %v) = %v; want %v", tc.value, tc.orElse, got, tc.want)
			}
		})
	}
}

func TestGetAsRequestPtrOrElse(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	testCases := []struct {
		name   string
		value  interface{}
		orElse *http.Request
		want   *http.Request
	}{
		{
			name:   "http.Request",
			value:  *req,
			orElse: nil,
			want:   req,
		},
		{
			name:   "*http.Request",
			value:  req,
			orElse: nil,
			want:   req,
		},
		{
			name:   "unsupported type",
			value:  42,
			orElse: req,
			want:   req,
		},
		{
			name:   "unsupported type with orElse nil",
			value:  "not a request",
			orElse: nil,
			want:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetAsRequestPtrOrElse(tc.value, tc.orElse)
			if (got == nil && tc.want != nil) || (got != nil && tc.want == nil) || (got != nil && tc.want != nil && !reflect.DeepEqual(got, tc.want)) {
				t.Errorf("GetAsRequestPtrOrElse(%v, %v) = %v; want %v", tc.value, tc.orElse, got, tc.want)
			}
		})
	}
}

func TestGetAsTimePtrOrElse(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name   string
		value  interface{}
		orElse *time.Time
		want   *time.Time
	}{
		{
			name:   "time.Time",
			value:  now,
			orElse: nil,
			want:   &now,
		},
		{
			name:   "*time.Time",
			value:  &now,
			orElse: nil,
			want:   &now,
		},
		{
			name:   "unsupported type",
			value:  42,
			orElse: &now,
			want:   &now,
		},
		{
			name:   "unsupported type with orElse nil",
			value:  "not a time",
			orElse: nil,
			want:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetAsTimePtrOrElse(tc.value, tc.orElse)
			if (got == nil && tc.want != nil) || (got != nil && tc.want == nil) || (got != nil && tc.want != nil && !got.Equal(*tc.want)) {
				t.Errorf("GetAsTimePtrOrElse(%v, %v) = %v; want %v", tc.value, tc.orElse, got, tc.want)
			}
		})
	}
}

func TestGetAsStringMapOrElse(t *testing.T) {
	testCases := []struct {
		name   string
		value  interface{}
		orElse map[string]string
		want   map[string]string
	}{
		{
			name:   "map[string]string",
			value:  map[string]string{"key": "value"},
			orElse: nil,
			want:   map[string]string{"key": "value"},
		},
		{
			name:   "*map[string]string",
			value:  newStringMap(map[string]string{"key": "value"}),
			orElse: nil,
			want:   map[string]string{"key": "value"},
		},
		{
			name:   "unsupported type",
			value:  42,
			orElse: map[string]string{"default": "value"},
			want:   map[string]string{"default": "value"},
		},
		{
			name:   "unsupported type with orElse nil",
			value:  "not a map",
			orElse: nil,
			want:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetAsStringMapOrElse(tc.value, tc.orElse)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetAsStringMapOrElse(%v, %v) = %v; want %v", tc.value, tc.orElse, got, tc.want)
			}
		})
	}
}

func newStringMap(m map[string]string) *map[string]string {
	return &m
}

func TestGetAsBoolOrElse(t *testing.T) {
	testCases := []struct {
		name   string
		value  interface{}
		orElse bool
		want   bool
	}{
		{"bool true", true, false, true},
		{"bool false", false, true, false},
		{"*bool true", newBool(true), false, true},
		{"*bool false", newBool(false), true, false},
		{"unsupported type", 42, true, true},
		{"unsupported type with orElse false", "not a bool", false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetAsBoolOrElse(tc.value, tc.orElse)
			if got != tc.want {
				t.Errorf("GetAsBoolOrElse(%v, %v) = %v; want %v", tc.value, tc.orElse, got, tc.want)
			}
		})
	}
}

func newBool(b bool) *bool {
	return &b
}

func TestGetAsStringOrElse(t *testing.T) {
	testCases := []struct {
		name   string
		value  interface{}
		orElse string
		want   string
	}{
		{"string", "hello", "", "hello"},
		{"*string", newString("hello"), "", "hello"},
		{"int", 42, "", "42"},
		{"int64", int64(42), "", "42"},
		{"uint", uint(42), "", "42"},
		{"uint64", uint64(42), "", "42"},
		{"int8", int8(42), "", "42"},
		{"int16", int16(42), "", "42"},
		{"int32", int32(42), "", "42"},
		{"uint8", uint8(42), "", "42"},
		{"uint16", uint16(42), "", "42"},
		{"uint32", uint32(42), "", "42"},
		{"float64", float64(42.0), "", "42"},
		{"float32", float32(42.0), "", "42"},
		{"unsupported type", struct{}{}, "", ""},
		{"unsupported type with orElse", struct{}{}, "default", "default"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetAsStringOrElse(tc.value, tc.orElse)
			if got != tc.want {
				t.Errorf("GetAsStringOrElse(%v, %q) = %q; want %q", tc.value, tc.orElse, got, tc.want)
			}
		})
	}
}

func newString(s string) *string {
	return &s
}

func TestGetAsIntOrElse(t *testing.T) {
	testCases := []struct {
		name   string
		value  interface{}
		orElse int
		want   int
	}{
		{"int", 42, 0, 42},
		{"int64", int64(42), 0, 42},
		{"uint", uint(42), 0, 42},
		{"uint64", uint64(42), 0, 42},
		{"int8", int8(42), 0, 42},
		{"int16", int16(42), 0, 42},
		{"int32", int32(42), 0, 42},
		{"uint8", uint8(42), 0, 42},
		{"uint16", uint16(42), 0, 42},
		{"uint32", uint32(42), 0, 42},
		{"float64", float64(42.0), 0, 42},
		{"float32", float32(42.0), 0, 42},
		{"unsupported type", "not a number", 0, 0},
		{"unsupported type with orElse", "not a number", -1, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetAsIntOrElse(tc.value, tc.orElse)
			if got != tc.want {
				t.Errorf("GetAsIntOrElse(%v, %d) = %d; want %d", tc.value, tc.orElse, got, tc.want)
			}
		})
	}
}
