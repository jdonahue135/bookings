package forms

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	isValid := form.Valid()

	if !isValid {
		t.Error("Got invalid when should be valid")
	}
}

func TestForm_Required(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	form = New(postedData)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("form shows invalid when required fields exist")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	field := "id"
	res := form.Has(field)

	if res {
		t.Error("Has shows true when no field exists")
	}

	postedData = url.Values{}
	postedData.Add("id", "12345")
	form = New(postedData)
	res = form.Has(field)

	if !res {
		t.Error("Has shows false when field should exist")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error but did not get one")
	}

	postedData = url.Values{}
	postedData.Add("id", "some value")
	form = New(postedData)
	form.MinLength("id", 100)
	if form.Valid() {
		t.Error("form shows min length of 100 met when data is shorter")
	}

	postedData = url.Values{}
	postedData.Add("id", "some value")
	form = New(postedData)
	form.MinLength("id", 10)
	if !form.Valid() {
		t.Error("form shows min length of 10 not met when data is 10")
	}
	isError = form.Errors.Get("id")
	if isError != "" {
		t.Error("should not have an error but got one")
	}
}

func TestForm_isEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedData = url.Values{}
	postedData.Add("email", "me@here.com")
	form = New(postedData)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("form shows not valid email for valid email address")
	}

	postedData = url.Values{}
	postedData.Add("email", "me")
	form = New(postedData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("form shows valid email for non valid email address")
	}

}
