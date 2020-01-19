package static

//go:generate go run github.com/tv42/becky -var=_ -wrap=addAsset article.css favicon.ico jquery.js jquery-ui.js notes.css notes.js slides.js styles.css play.js playground.js

var Assets = make(map[string]string)

func addAsset(a asset) struct{} {
	Assets[a.Name] = a.Content
	return struct{}{}
}
