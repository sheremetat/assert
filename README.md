# Assert - Golang assertion library
![](https://github.com/sheremetat/assert/workflows/Verification/badge.svg)

This simple library allows developers to quickly verify certain assumptions or state of a Golang program. Once assertions enabled and assertion expression equals `true` assertion library will exit from your application with exit status 1 and write assertion error to your output.

This library extends your application's command-line API with two flags:

    --ea, --enableassertions

To enable assertions in your application please provide one of two boolean flags(which are equivalent). For more information of the flags format see documentation  [here](https://golang.org/pkg/flag/). Set one of this flags to true to enable assertions for the application or false to disable assertion. If no one of these flags were not provided assertions will be disabled by default.

You can use any format of flags usage according to Golang flags documentation. Also don't forget to parse flags in your application with the command

    flag.Parse()

in the main function of your application.

## Installation

To install Assert library, use `go get`:

    go get github.com/sheremetat/assert

## Supported go versions

We support the latest major Go 1.14 version

## Enable Assertion

You can enable assertion with provides CLI flags, any of following flags are valid. Also, you can use `--` notation.

    -ea
    -ea=true
    -ea=1
    -enableassertions
    -enableassertions=true
    -enableassertions=1
    
Similar disabling assertions (they are disabled by default)

    -ea=false
    -ea=0
    -enableassertions=false
    -enableassertions=0
    
Then in your code:

```go
func main(){
    // ... configure you flags
    flag.Parse()
    // ... your code

}
```

## Using Assertions

The library has only one public method `When` for assertions. Examples:

```go
import "github.com/sheremetat/assert"

func myFunc(key string) {
	assert.When(len(key) == 0, "Empty key should never passed here")
	switch key {
	case "A": 
		// ... do something for A
	case "B":
		// ... do something for B
	default:
		assert.When(true, "Unsupported KEY `%s`", key)
	}
}
```

Example log output:

    Assertion error in /path/to/project/main/main.go#10: Unsupported KEY `key`
    exit status 1

## Best Practices

The most important thing to remember about assertions is that they can be disabled, so never assume they'll be executed.

Therefore, keep the followings things in mind when using assertions:

- Always check for nil values where appropriate
- Avoid using assertions to check inputs into a public method and instead use an proper input variables checking
- Don't call methods in assertion conditions and instead assign the result of the method to a local variable and use that variable with assert
- Assertions are great for places in the code that will never be executed, such as the default case of a switch statement or after a loop that never finishes

## Acknowledgments

This library inspired by [Java Assert Keyword](https://docs.oracle.com/javase/8/docs/technotes/guides/language/assert.html)

## License

Copyright (c) 2020 Taras Sheremeta

This project licensed under the terms of the MIT license.