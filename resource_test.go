package app

import "testing"
import "os"
import "path/filepath"

func TestResourcePathJoin(t *testing.T) {
	l := ResourcePath("resources")

	if j := l.Join("css"); j != "resources/css" {
		t.Error("j should be resources/css:", j)
	}
}

func TestResourcePathCSS(t *testing.T) {
	r := ResourcePath("resources")
	cssPath := r.Join("css")

	if err := os.MkdirAll(cssPath, os.ModePerm); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(Resources().Path())

	os.Create(filepath.Join(cssPath, "foo.css"))
	os.Create(filepath.Join(cssPath, "hello"))
	os.Create(filepath.Join(cssPath, "world.txt"))
	os.Mkdir(filepath.Join(cssPath, "bar"), os.ModePerm)

	cssFilenames := r.CSS()

	if l := len(cssFilenames); l != 1 {
		t.Error("cssFilenames should have 1 element:", l)
	}

	if cssFilenames[0] != "foo.css" {
		t.Error("cssFilenames[0] should be foo.css:", cssFilenames[0])
	}
}

func TestResourcePathCSSError(t *testing.T) {
	// No css directory.
	r := ResourcePath("resources")
	r.CSS()

	// css is not a directory.
	cssPath := r.Join("css")
	os.Mkdir(r.Path(), os.ModePerm)
	os.Create(cssPath)
	defer os.RemoveAll(Resources().Path())
	r.CSS()
}

func TestResources(t *testing.T) {
	t.Log(Resources())
}
