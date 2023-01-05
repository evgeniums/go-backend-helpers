package test_common

import (
	"reflect"
	"testing"

	"github.com/evgeniums/go-backend-helpers/pkg/common"
)

func TestWithPath(t *testing.T) {

	create := func(path string) common.WithPath {
		p := &common.WithPathBase{}
		p.Init(path)
		if p.Path() != path {
			t.Fatalf("invalid path: expected %s, got %s", path, p.Path())
		}
		if p.FullPath() != path {
			t.Fatalf("invalid full path: expected %s, got %s", path, p.FullPath())
		}

		return p
	}

	path1 := "/section1/section2"
	p1 := create(path1)
	if p1.Separator() != "/" {
		t.Fatalf("invalid separator")
	}
	sections1 := []string{"", "section1", "section2"}
	if !reflect.DeepEqual(p1.Sections(), sections1) {
		t.Fatalf("invalid sections: expected %v, got %v", sections1, p1.Sections())
	}
	constructedPath1 := common.ConstructPath(sections1, "/")
	if constructedPath1 != path1 {
		t.Fatalf("invalid constructed path: expected %s, got %s", path1, constructedPath1)
	}
	paths1 := []string{"/", "/section1", "/section1/section2"}
	if !reflect.DeepEqual(p1.Paths(), paths1) {
		t.Fatalf("invalid paths: expected %v, got %v", paths1, p1.Paths())
	}

	path2 := "/section1/section2/"
	p2 := create(path2)
	sections2 := []string{"", "section1", "section2", ""}
	if !reflect.DeepEqual(p2.Sections(), sections2) {
		t.Fatalf("invalid sections: expected %v, got %v", sections2, p2.Sections())
	}
	constructedPath2 := common.ConstructPath(sections2, "/")
	if constructedPath2 != path2 {
		t.Fatalf("invalid constructed path: expected %s, got %s", path2, constructedPath2)
	}
	paths2 := []string{"/", "/section1", "/section1/section2", "/section1/section2/"}
	if !reflect.DeepEqual(p2.Paths(), paths2) {
		t.Fatalf("invalid paths: expected %v, got %v", paths2, p2.Paths())
	}

	path3 := "/section3/section4"
	p3 := create(path3)
	p3.SetParent(p1)
	if p3.Path() != path3 {
		t.Fatalf("invalid full path: expected %s, got %s", path3, p3.Path())
	}
	fullPath3 := "/section1/section2/section3/section4"
	if p3.FullPath() != fullPath3 {
		t.Fatalf("invalid full path: expected %s, got %s", fullPath3, p3.FullPath())
	}
	sections3 := []string{"", "section1", "section2", "section3", "section4"}
	if !reflect.DeepEqual(p3.Sections(), sections3) {
		t.Fatalf("invalid sections: expected %v, got %v", sections3, p3.Sections())
	}
	paths3 := []string{"/", "/section1", "/section1/section2", "/section1/section2/section3", "/section1/section2/section3/section4"}
	if !reflect.DeepEqual(p3.Paths(), paths3) {
		t.Fatalf("invalid paths: expected %v, got %v", paths3, p3.Paths())
	}

	p3.SetParent(p2)
	if p3.Path() != path3 {
		t.Fatalf("invalid full path: expected %s, got %s", path3, p3.Path())
	}
	if !reflect.DeepEqual(p3.Sections(), sections3) {
		t.Fatalf("invalid sections: expected %v, got %v", sections3, p3.Sections())
	}
	if p3.FullPath() != fullPath3 {
		t.Fatalf("invalid full path: expected %s, got %s", fullPath3, p3.FullPath())
	}
	if !reflect.DeepEqual(p3.Paths(), paths3) {
		t.Fatalf("invalid paths: expected %v, got %v", paths3, p3.Paths())
	}
}
