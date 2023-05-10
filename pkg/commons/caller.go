package commons

import (
	"runtime"

	"github.com/logx-go/contract/pkg/logx"
)

// SetCallerInfo add the caller info to the fields map
func SetCallerInfo(skip int, override bool, fields map[string]any) map[string]any {
	for fieldName := range fields {
		if fieldName == logx.FieldNameCallerFunc || fieldName == logx.FieldNameCallerFile || fieldName == logx.FieldNameCallerLine {
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

	fields[logx.FieldNameCallerFunc] = funcName
	fields[logx.FieldNameCallerFile] = file
	fields[logx.FieldNameCallerLine] = line

	return fields
}
