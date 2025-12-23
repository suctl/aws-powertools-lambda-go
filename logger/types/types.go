package types

import (
	"io"
)

type LogConfig struct {
	Writer         io.Writer
	IncludeContext bool
	Properties     map[string]string
}

type LambdaContext struct {
	FunctionName       string
	FunctionMemorySize string
	FunctionARN        string
	FunctionRequestId  string
}
