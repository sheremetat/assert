package assert

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestItWithEnabledAssertionsAndTrueExpression(t *testing.T) {
	b := true
	enabled = &b

	if os.Getenv("FLAG") == "1" {
		When(true, "Expression %d!=%d is invalid", 1, 2)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestItWithEnabledAssertionsAndTrueExpression")
	cmd.Env = append(os.Environ(), "FLAG=1")
	err := cmd.Run()

	e, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatal("Expected exit with error")
	}

	expectedErrorString := "exit status 1"
	if expectedErrorString != e.Error() {
		t.Fatalf("Expected exit with error with message: %s", e)
	}
}

func TestItWithEnabledAssertionsAndFalseExpression(t *testing.T) {
	b := true
	enabled = &b

	if os.Getenv("FLAG") == "1" {
		When(false, "test message")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestItWithEnabledAssertionsAndFalseExpression")
	cmd.Env = append(os.Environ(), "FLAG=1")
	err := cmd.Run()

	if err != nil {
		t.Fatalf("Expected exit without error but was %s", err.Error())
	}

	err, ok := err.(*exec.ExitError)
	if ok {
		t.Fatalf("Expected exit without error but was %s", err.Error())
	}
}

func TestItWithDisabledAssertions(t *testing.T) {
	b := false
	enabled = &b

	if os.Getenv("FLAG") == "1" {
		When(true, "test message")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestItWithDisabledAssertions")
	cmd.Env = append(os.Environ(), "FLAG=1")
	err := cmd.Run()

	if err != nil {
		t.Fatalf("Expected exit without error but was %s", err.Error())
	}

	err, ok := err.(*exec.ExitError)
	if ok {
		t.Fatalf("Expected exit without error but was %s", err.Error())
	}
}

func Test_buildErrorMessage(t *testing.T) {
	callerInformation := "testing.go#992"

	type args struct {
		message string
		params  []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Must contains prefix",
			args: args{
				"test message",
				nil,
			},
			want: "Assertion error ",
		},
		{
			name: "Build message with no params",
			args: args{
				"test message",
				nil,
			},
			want: ": test message",
		},
		{
			name: "Build message with params",
			args: args{
				"test message %s %d",
				[]interface{}{"param1", 2},
			},
			want: ": test message param1 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildErrorMessageWithCallerMetadata(callerInformation, tt.args.message, tt.args.params...)
			if !strings.Contains(got, tt.want) {
				t.Errorf("return error must contains custom error message `%s`, but was %s", tt.want, got)
			}
			if !strings.Contains(got, callerInformation) {
				t.Errorf("return error must contains caller file and line `%s`, but was %s", callerInformation, got)
			}
		})
	}
}

func Test_buildCallerInformation(t *testing.T) {
	type args struct {
		skipMe bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "With skip this file",
			args: args{
				skipMe: true,
			},
			want: "testing.go",
		},
		{
			name: "Without skip this file",
			args: args{
				skipMe: false,
			},
			want: "assert_test.go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildCallerInformation(tt.args.skipMe); !strings.Contains(got, tt.want) {
				t.Errorf("buildCallerInformation() must contains file name `%s`, but was %s", tt.want, got)
			}
		})
	}
}
