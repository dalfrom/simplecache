package tests

import (
	"reflect"
	"testing"

	"github.com/dalfrom/simplecache/pkg/scl"
)

// TestSetScl calls scl.SetScl with a valid SET statement, checking
// for the data to be returned correctly, without a TTI value.
func TestSetScl(t *testing.T) {
	input := "SET users.id:1;"
	want := &scl.SetStmt{Collection: "users", Key: "id", Value: "1", Config: scl.Config{Tti: ""}}

	got, err := scl.SetScl(input)
	if err != nil {
		t.Errorf(`SetScl(%q) returned an error: %v`, input, err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`SetScl(%q) = %+v, want %+v`, input, got, want)
	}
}

// TestSetSclWithTTI calls scl.SetScl with a valid SET statement, checking
// for the data to be returned correctly, together with the TTI value.
func TestSetSclWithTTI(t *testing.T) {
	input := "SET users.id:1 TTI=15;"
	want := &scl.SetStmt{Collection: "users", Key: "id", Value: "1", Config: scl.Config{Tti: "15"}}

	got, err := scl.SetScl(input)
	if err != nil {
		t.Errorf(`SetScl(%q) returned an error: %v`, input, err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`SetScl(%q) = %+v, want %+v`, input, got, want)
	}
}

// TestGetScl calls scl.GetScl with a valid GET statement, checking
// for the data to be returned correctly.
func TestGetScl(t *testing.T) {
	input := "GET users.id;"
	want := &scl.GetStmt{Collection: "users", Key: "id"}

	got, err := scl.GetScl(input)
	if err != nil {
		t.Errorf(`GetScl(%q) returned an error: %v`, input, err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`GetScl(%q) = %+v, want %+v`, input, got, want)
	}
}

// TestGetEverythingScl calls scl.GetScl with a valid GET statement, checking
// for the data to be returned correctly.
func TestGetEverythingScl(t *testing.T) {
	input := "GET users.*;"
	want := &scl.GetStmt{Collection: "users", Key: "*"}

	got, err := scl.GetScl(input)
	if err != nil {
		t.Errorf(`GetScl(%q) returned an error: %v`, input, err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`GetScl(%q) = %+v, want %+v`, input, got, want)
	}
}

// TestDeleteScl calls scl.DeleteScl with a valid DELETE statement, checking
// for the data to be returned correctly.
func TestDeleteScl(t *testing.T) {
	input := "DELETE users.id;"
	want := &scl.DeleteStmt{Collection: "users", Key: "id"}

	got, err := scl.DeleteScl(input)
	if err != nil {
		t.Errorf(`DeleteScl(%q) returned an error: %v`, input, err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`DeleteScl(%q) = %+v, want %+v`, input, got, want)
	}
}

// TestTruncateScl calls scl.TruncateScl with a valid TRUNCATE statement, checking
// for the data to be returned correctly.
func TestTruncateScl(t *testing.T) {
	input := "TRUNCATE users;"
	want := &scl.TruncateStmt{Collection: "users"}

	got, err := scl.TruncateScl(input)
	if err != nil {
		t.Errorf(`TruncateScl(%q) returned an error: %v`, input, err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`TruncateScl(%q) = %+v, want %+v`, input, got, want)
	}
}

// TestGetScl calls scl.DropScl with a valid DROP statement, checking
// for the data to be returned correctly
func TestDropScl(t *testing.T) {
	input := "DROP users;"
	want := &scl.DropStmt{Collection: "users", Key: "*"}

	got, err := scl.DropScl(input)
	if err != nil {
		t.Errorf(`DropScl(%q) returned an error: %v`, input, err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`DropScl(%q) = %+v, want %+v`, input, got, want)
	}
}

// TestUpdateScl calls scl.UpdateScl with a valid UPDATE statement, checking
// for the data to be returned correctly, without a TTI value.
func TestUpdateScl(t *testing.T) {
	input := "UPDATE users.id:2;"
	want := &scl.UpdateStmt{Collection: "users", Key: "id", Value: "2"}

	got, err := scl.UpdateScl(input)
	if err != nil {
		t.Errorf(`UpdateScl(%q) returned an error: %v`, input, err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`UpdateScl(%q) = %+v, want %+v`, input, got, want)
	}
}

// TestUpdateSclWithTTI calls scl.UpdateScl with a valid UPDATE statement, checking
// for the data to be returned correctly, with a TTI value.
func TestUpdateSclWithTTI(t *testing.T) {
	input := "UPDATE users.id:2 TTI=5;"
	want := &scl.UpdateStmt{Collection: "users", Key: "id", Value: "2", Config: scl.Config{Tti: "5"}}

	got, err := scl.UpdateScl(input)
	if err != nil {
		t.Errorf(`UpdateScl(%q) returned an error: %v`, input, err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`UpdateScl(%q) = %+v, want %+v`, input, got, want)
	}
}
