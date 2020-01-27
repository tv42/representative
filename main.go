package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"eagain.net/go/representative/internal/golang_org_x_tools/static"
	"eagain.net/go/representative/internal/golang_org_x_tools/templates"
	"github.com/dchest/safefile"
	"golang.org/x/tools/present"
)

func init() {
	present.PlayEnabled = true
	// TODO present.NotesEnabled should probably come from a flag
}

func writeAssets(dir string) error {
	if err := os.Mkdir(dir, 0755); err != nil && !errors.Is(err, os.ErrExist) {
		return fmt.Errorf("cannot make asset directory: %v", err)
	}
	for name, content := range static.Assets {
		if err := safefile.WriteFile(filepath.Join(dir, name), []byte(content), 0644); err != nil {
			return err
		}
	}
	return nil
}

func playable(c present.Code) bool {
	return c.Play && c.Ext == ".go"
}

func convert(src string, urlToAssets *url.URL) error {
	tmpl := present.Template()
	tmpl.Funcs(template.FuncMap{
		// required by present
		"playable": playable,

		// Since present controls the data passed to the
		// template, we need to pass things through a
		// function.
		//
		// Alternatively, we could reimplement the minimal
		// amount of code in present.Doc.Render.
		"static": func() string {
			return urlToAssets.String()
		},
	})
	if _, err := tmpl.Parse(templates.Action); err != nil {
		return err
	}
	ext := filepath.Ext(src)
	if ext == "" {
		return fmt.Errorf("source file must have an extension: %s", src)
	}
	byExt := map[string]string{
		".article": templates.Article,
		".slide":   templates.Slides,
	}
	content, ok := byExt[ext]
	if !ok {
		return fmt.Errorf("unknown extension: %s", ext)
	}
	if _, err := tmpl.Parse(content); err != nil {
		return err
	}

	srcF, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcF.Close()
	doc, err := present.Parse(srcF, src, 0)
	if err != nil {
		return err
	}
	srcF.Close()

	dst := src + ".html"
	dstF, err := safefile.Create(dst, 0644)
	if err != nil {
		return err
	}
	defer dstF.Close()

	if err := doc.Render(dstF, tmpl); err != nil {
		return err
	}

	if err := dstF.Commit(); err != nil {
		return err
	}

	return nil
}

const prog = "representative"

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", prog)
	fmt.Fprintf(flag.CommandLine.Output(), "  %s [OPTS] SLIDES_AND_ARTICLES..\n", prog)
	fmt.Fprintf(flag.CommandLine.Output(), "  %s -assets=DIR\n", prog)
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	var assets = flag.String("assets", "", "write assets to `DIR`")
	var urlToAssets URLFlag
	urlToAssets.Default = func() *url.URL {
		if s := *assets; s != "" {
			return &url.URL{Path: s}
		}
		return &url.URL{Path: "static"}
	}
	flag.Var(&urlToAssets, "url-to-assets", "base `URL` to fetch assets from; -assets sets this too")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 && *assets == "" {
		flag.Usage()
		os.Exit(2)
	}

	for _, arg := range flag.Args() {
		if err := convert(arg, urlToAssets.GetURL()); err != nil {
			log.Fatalf("converting: %v", err)
		}
	}
	if dir := *assets; dir != "" {
		if err := writeAssets(dir); err != nil {
			log.Fatalf("writing assets: %v", err)
		}
	}
}
