// Copyright (c) 2020 Taras Sheremeta. All right reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file

// Package assert is a simple library that allows developers to quickly
// verify certain assumptions or state of a program. Once assertions
// enabled and assertion expression equals `true` assertion library will exit
// from your application with exit status 1 and write assertion error to your output.
//
// This library extends your application's command-line API with two flags:
// 		--ea, --enableassertions
// To enable assertions in your application please provide one of two boolean flags
// (which are equivalent). For more information of the flags format see documentation
// here https://golang.org/pkg/flag/:
//
// Set one of this flags to true to enable assertions for the application or false to
// disable assertion. If no one of these flags was not provided assertions will be disabled
// by default.
//
// You can use any format of flags usage according to Golang flags documentation
// Also don't forget to parse flags in your application with the command
// 		flag.Parse()
// in your application
package assert

import (
	"flag"
	"fmt"
	"log"
	"runtime"
)

var enabled = flag.Bool("enableassertions", false,
	"Set to true to enable assertions for the application or false to disable assertion")

func init() {
	flag.BoolVar(enabled, "ea", *enabled,
		"Set to true to enable assertions for the application or false to disable assertion")
}

// When validates expression and, if assertions are enabled and expression is true,
// print error message and exit from application with error status 1
func When(expression bool, message string, params ...interface{}) {
	if !*enabled || !expression {
		return
	}

	callerInfo := buildCallerInformation(true)
	log.Fatalf(buildErrorMessageWithCallerMetadata(callerInfo, message, params...))
}

// buildErrorMessageWithCallerMetadata builds error message using caller information
// and then join this with custom message (optionally with params) and return
// string built by format template.
func buildErrorMessageWithCallerMetadata(callerInformation, message string, params ...interface{}) string {
	assertionMessagePrefix := fmt.Sprintf("Assertion error %s: ", callerInformation)
	return fmt.Sprintf(assertionMessagePrefix+message, params...)
}

// buildCallerInformation extracts caller information (filename and line) from runtime
// and combine this information into pre-formatted string
func buildCallerInformation(skipMe bool) string {
	skipCallers := 1 // skip this call by default
	if skipMe {
		skipCallers++ // Skip caller by request
	}

	_, file, line, ok := runtime.Caller(skipCallers)
	if ok {
		return fmt.Sprintf("in %s#%d", file, line)
	}
	return ""
}
