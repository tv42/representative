// Package representative creates static HTML for Go "represent"
// format slides and articles.
//
// For more information, see
// https://pkg.go.dev/golang.org/x/tools/present
//
// The HTML and assets are co-dependent and should be created with the
// same version of representative.
package representative

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/dchest/safefile"
	"golang.org/x/tools/present"
)

func playable(c present.Code) bool {
	return c.Play && c.Ext == ".go"
}

//go:embed internal/golang_org_x_tools/templates/action.tmpl
var tmplAction string

//go:embed internal/golang_org_x_tools/templates/article.tmpl
var tmplArticle string

//go:embed internal/golang_org_x_tools/templates/slides.tmpl
var tmplSlides string

// Convert writes the HTML for the slides or article at file path src
// to w.
//
// For assets that are included in the HTML (such as code examples),
// it reads assets from files relative to src. The final HTML uses
// urlToAssets as a base URL to load assets.
func Convert(w io.Writer, src string, urlToAssets *url.URL) error {
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
	if _, err := tmpl.Parse(tmplAction); err != nil {
		return err
	}
	ext := filepath.Ext(src)
	if ext == "" {
		return fmt.Errorf("source file must have an extension: %s", src)
	}
	byExt := map[string]string{
		".article": tmplArticle,
		".slide":   tmplSlides,
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

	if err := doc.Render(w, tmpl); err != nil {
		return err
	}
	return nil
}

//go:embed internal/golang_org_x_tools/static
var assets embed.FS

// WriteAssets writes the asset files into dir. The directory will be
// created if needed. The exact set of files created should not be
// relied on, and this directory should not be used for other
// purposes.
func WriteAssets(dir string) error {
	if err := os.Mkdir(dir, 0755); err != nil && !errors.Is(err, os.ErrExist) {
		return fmt.Errorf("cannot make asset directory: %v", err)
	}
	const staticDir = "internal/golang_org_x_tools/static"
	entries, err := assets.ReadDir(staticDir)
	if err != nil {
		return fmt.Errorf("cannot list assets: %v", err)
	}
	for _, entry := range entries {
		if !entry.Type().IsRegular() {
			return fmt.Errorf("unexpected non-file asset: %q", entry.Name())
		}
		name := entry.Name()
		content, err := assets.ReadFile(path.Join(staticDir, name))
		if err != nil {
			return fmt.Errorf("cannot read asset: %s: %v", name, err)
		}
		if err := safefile.WriteFile(filepath.Join(dir, name), []byte(content), 0644); err != nil {
			return err
		}
	}
	return nil
}
