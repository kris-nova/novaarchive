package logger

import (
	"bytes"
	"fmt"
	"testing"
)

const (
	TestFormat = "%s %s %s money, success, fame glamour"
)

var (
	testA = []interface{}{"1", "2", "3"}
)

func TestAlwaysOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogAlways
	buffer := new(bytes.Buffer)
	Writer = buffer
	expected := Line(PreAlways, TestFormat, testA...)
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestSuccessOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogSuccess
	buffer := new(bytes.Buffer)
	Writer = buffer
	expected := Line(PreSuccess, TestFormat, testA...)
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestDebugOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogDebug
	buffer := new(bytes.Buffer)
	Writer = buffer
	expected := Line(PreDebug, TestFormat, testA...)
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestInfoOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogInfo
	buffer := new(bytes.Buffer)
	Writer = buffer
	expected := Line(PreInfo, TestFormat, testA...)
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestWarningOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogWarning
	buffer := new(bytes.Buffer)
	Writer = buffer
	expected := Line(PreWarning, TestFormat, testA...)
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestCriticalOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogCritical
	buffer := new(bytes.Buffer)
	Writer = buffer
	expected := Line(PreCritical, TestFormat, testA...)
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestDeprecatedOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogDeprecated
	buffer := new(bytes.Buffer)
	Writer = buffer
	expected := Line(PreDeprecated, TestFormat, testA...)
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestEverything(t *testing.T) {
	Level = -1
	BitwiseLevel = LogEverything
	buffer := new(bytes.Buffer)
	Writer = buffer
	cases := []string{PreDeprecated, PreDebug, PreInfo, PreWarning, PreCritical, PreSuccess, PreAlways}
	expected := ""
	for _, c := range cases {
		expected = fmt.Sprintf("%s%s", expected, Line(c, TestFormat, testA...))
	}
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestAlwaysCriticalDebugOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogAlways | LogCritical | LogDebug
	buffer := new(bytes.Buffer)
	Writer = buffer
	cases := []string{PreDebug, PreCritical, PreAlways}
	expected := ""
	for _, c := range cases {
		expected = fmt.Sprintf("%s%s", expected, Line(c, TestFormat, testA...))
	}
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}
