package representative_test

import (
	"bytes"
	"io/ioutil"
	"net/url"
	"strings"
	"testing"

	"eagain.net/go/representative"
	"github.com/andybalholm/cascadia"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

func TestSlidesStaticPath(t *testing.T) {
	buf := new(bytes.Buffer)
	if err := representative.Convert(
		buf,
		"testdata/simple.slide",
		&url.URL{Path: "xyzzy/foo/bar"},
	); err != nil {
		t.Fatalf("convert: %v", err)
	}
	tree, err := html.Parse(buf)
	if err != nil {
		t.Fatalf("parse slide html: %v", err)
	}
	sel := cascadia.MustCompile(`script[src]`)
	nodes := cascadia.QueryAll(tree, sel)
	for _, n := range nodes {
		for _, attr := range n.Attr {
			if attr.Namespace == "" && attr.Key == "src" {
				if !strings.HasPrefix(attr.Val, "xyzzy/foo/bar/") {
					t.Errorf("script source points outside static url: %v", attr.Val)
				}
			}
		}
	}
}

func readdirnames(dir string) ([]string, error) {
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var r []string
	for _, fi := range fis {
		r = append(r, fi.Name())
	}
	return r, nil
}

func TestWriteAssets(t *testing.T) {
	const staticDir = "testdata/static"
	if err := representative.WriteAssets(staticDir); err != nil {
		t.Fatalf("writeAssets: %v", err)
	}
	names, err := readdirnames(staticDir)
	if err != nil {
		t.Fatalf("readdir: %v", err)
	}

	want := []string{
		"article.css",
		"favicon.ico",
		"jquery-ui.js",
		"jquery.js",
		"notes.css",
		"notes.js",
		"play.js",
		"playground.js",
		"slides.js",
		"styles.css",
	}
	if diff := cmp.Diff(want, names); diff != "" {
		t.Errorf("wrong assets:\n%s", diff)
	}
}
