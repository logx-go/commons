package commons

import (
	"runtime"
)

// SetCallerInfo add the caller info to the fields map
func SetCallerInfo(skip int, override bool, fields map[string]any, fieldNameCallerFunc, fieldNameCallerFile, fieldNameCallerLine string) map[string]any {
	for fieldName := range fields {
		if fieldName == fieldNameCallerFunc || fieldName == fieldNameCallerFile || fieldName == fieldNameCallerLine {
			if !override {
				return fields
			}

			continue
		}
	}

	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return fields
	}

	funcName := runtime.FuncForPC(pc).Name()

	fields[fieldNameCallerFunc] = funcName
	fields[fieldNameCallerFile] = file
	fields[fieldNameCallerLine] = line

	return fields
}
