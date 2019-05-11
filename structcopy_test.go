package gostructcopy

import (
	"testing"
)

func BenchmarkStructCopy(bench *testing.B) {
	a := struct {
		Id     int
		Name   string
		Weight int
		a      int
	}{100, "Dog", 200, 9}
	b := struct {
		Id   int
		Name string
		Desc string
		b    int
	}{}
	for i := 0; i < bench.N; i++ {
		StructCopy(&a, &b)
	}
}

func TestStructCopy(t *testing.T) {
	a := struct {
		Id     int
		Name   string
		Weight int
		a      int
	}{100, "Dog", 200, 9}
	b := struct {
		Id   int
		Name string
		Desc string
		b    int
	}{}
	err := StructCopy(&a, &b)
	if err != nil {
		t.Fatal(err)
	} else if a.Id == b.Id && a.Name == b.Name {
		t.Log("Success")
	} else {
		t.Fatal("Copy Fail")
	}
}
