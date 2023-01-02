package word

import (
	"log"
	"testing"
)

func TestToLower(t *testing.T) {
	s := ToLower("ABC")
	if s != "abc" {
		log.Fatal("TestToLower Error!")
	}
}

func TestToUpper(t *testing.T) {
	s := ToUpper("abc")
	if s != "ABC" {
		log.Fatal("TestToUpper Error!")
	}
}

func TestUnderscoreToLowerCamelCase(t *testing.T) {
	s := UnderscoreToLowerCamelCase("test_lower_camel_case")
	if s != "testLowerCamelCase" {
		log.Fatal("TestUnderscoreToLowerCamelCase Error!")
	}
}

func TestUnderscoreToUpperCamelCase(t *testing.T) {
	s := UnderscoreToUpperCamelCase("test_underscore_to_upper_camel_case")
	if s != "TestUnderscoreToUpperCamelCase" {
		log.Fatal("TestUnderscoreToUpperCamelCase Error!")
	}
}

func TestCamelCaseToUnderscore(t *testing.T) {
	s1 := CamelCaseToUnderscore("TestUnderscoreToUpperCamelCase")
	if s1 != "test_underscore_to_upper_camel_case" {
		log.Fatal("TestCamelCaseToUnderscore s1 Error!")
	}
	s2 := CamelCaseToUnderscore("testLowerCamelCase")
	if s2 != "test_lower_camel_case" {
		log.Fatal("TestCamelCaseToUnderscore s2 Error!")
	}
}