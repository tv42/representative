package templates

//go:generate go run github.com/tv42/becky -var=Action action.tmpl
//go:generate go run github.com/tv42/becky -var=Article article.tmpl
//go:generate go run github.com/tv42/becky -var=Slides slides.tmpl

func tmpl(a asset) string {
	return a.Content
}
