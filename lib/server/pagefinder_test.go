package server

import (
	b "homestead/lib/blogfsys"
	"homestead/lib/generator"
	"strings"
	"testing"
)

const (
	root      string = "../../test"
	testindex string = "index.html"
	testpost  string = "post"
)

func TestFindInFsys(t *testing.T) {
	var want int = 1

	setEnv()
	finder := pageFinder{
		fsys: b.NewBlogFsys(root),
	}

	if found, err := finder.findInFsys(b.Index); err != nil {
		t.Fatal(err)
	} else if len(found) != want {
		t.Fatalf("Expected %d, got: %d", want, len(found))
	} else if !strings.Contains(found[0].GetPath(), b.Index) {
		t.Fatalf("Expected %s, found: %s", b.Index, found[0].GetPath())
	}

	if found, err := finder.findInFsys(b.Posts); err != nil {
		t.Fatal(err)
	} else if len(found) != want {
		t.Fatalf("Expected %d, got: %d", want, len(found))
	} else if !strings.Contains(found[0].GetPath(), b.Posts) {
		t.Fatalf("Expected %s, found: %s", b.Index, found[0].GetPath())
	}
}

func TestGetIndex(t *testing.T) {
	setEnv()

	finder := NewPageFinder(root)

	idx := finder.GetIndex()
	if idx == nil {
		t.Fatal("Expected index.html, found nil")
	} else if !strings.Contains(idx.GetPath(), testindex) {
		t.Fatalf("Expected index.html, found %s", idx.GetPath())
	}
}

func TestGetPost(t *testing.T) {
	setEnv()

	finder := NewPageFinder(root)

	post := finder.GetPost(testpost)
	if post == nil {
		t.Fatal("Expected a post, found nil")
	}
}

func setEnv() {
	// Make sure that the public is generated before tests
	gen := generator.NewGenerator(root)
	gen.GenerateStaticContent()
}
