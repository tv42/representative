package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"eagain.net/go/representative"
	"github.com/dchest/safefile"
	"golang.org/x/tools/present"
)

func init() {
	present.PlayEnabled = true
	// TODO present.NotesEnabled should probably come from a flag
}

func convert(src string, urlToAssets *url.URL) error {
	dst := src + ".html"
	dstF, err := safefile.Create(dst, 0644)
	if err != nil {
		return err
	}
	defer dstF.Close()
	if err := representative.Convert(dstF, src, urlToAssets); err != nil {
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
		if err := representative.WriteAssets(dir); err != nil {
			log.Fatalf("writing assets: %v", err)
		}
	}
}
