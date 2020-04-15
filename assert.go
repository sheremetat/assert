// Copyright (c) 2020 Taras Sheremeta. All right reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file

// Package assert is a simple library that allows developers to quickly
// verify certain assumptions or state of a program.
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

import "flag"

var enabled = flag.Bool("enableassertions", false,
	"Set to true to enable assertions for the application or false to disable assertion")

func init() {
	flag.BoolVar(enabled, "ea", *enabled,
		"Set to true to enable assertions for the application or false to disable assertion")
}
