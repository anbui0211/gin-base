package pruntime

import (
	"fmt"
	"log"
	"runtime"
)

// Caller ...
func Caller(name string, value interface{}) string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("-- Caller error -- ")
	}

	return fmt.Sprintf("🚀 Caller: %s\n🚀 From: %s#%d\n🚀 Value: %v", name, file, line, value)
}
