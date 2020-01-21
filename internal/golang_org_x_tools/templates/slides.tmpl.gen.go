// Code generated by github.com/tv42/becky -- DO NOT EDIT.

package templates

var Slides = tmpl(asset{Name: "slides.tmpl", Content: "" +
	"{/* This is the slide template. It defines how presentations are formatted. */}\n\n{{define \"root\"}}\n<!DOCTYPE html>\n<html>\n  <head>\n    <title>{{.Title}}</title>\n    <meta charset='utf-8'>\n    <link rel=\"icon\" href=\"{{static}}/favicon.ico\">\n    <script>\n      var notesEnabled = {{.NotesEnabled}};\n    </script>\n    <script src='{{static}}/slides.js'></script>\n\n    {{if .NotesEnabled}}\n    <script>\n      var sections = {{.Sections}};\n      var titleNotes = {{.TitleNotes}}\n    </script>\n    <script src='{{static}}/notes.js'></script>\n    {{end}}\n\n    <script>\n      // Initialize Google Analytics tracking code on production site only.\n      if (window[\"location\"] && window[\"location\"][\"hostname\"] == \"talks.golang.org\") {\n        var _gaq = _gaq || [];\n        _gaq.push([\"_setAccount\", \"UA-11222381-6\"]);\n        _gaq.push([\"b._setAccount\", \"UA-49880327-6\"]);\n        window.trackPageview = function() {\n          _gaq.push([\"_trackPageview\", location.pathname+location.hash]);\n          _gaq.push([\"b._trackPageview\", location.pathname+location.hash]);\n        };\n        window.trackPageview();\n        window.trackEvent = function(category, action, opt_label, opt_value, opt_noninteraction) {\n          _gaq.push([\"_trackEvent\", category, action, opt_label, opt_value, opt_noninteraction]);\n          _gaq.push([\"b._trackEvent\", category, action, opt_label, opt_value, opt_noninteraction]);\n        };\n      }\n    </script>\n  </head>\n\n  <body style='display: none'>\n\n    <section class='slides layout-widescreen'>\n\n      <article>\n        <h1>{{.Title}}</h1>\n        {{with .Subtitle}}<h3>{{.}}</h3>{{end}}\n        {{if not .Time.IsZero}}<h3>{{.Time.Format \"2 January 2006\"}}</h3>{{end}}\n        {{range .Authors}}\n          <div class=\"presenter\">\n            {{range .TextElem}}{{elem $.Template .}}{{end}}\n          </div>\n        {{end}}\n      </article>\n\n  {{range $i, $s := .Sections}}\n  <!-- start of slide {{$s.Number}} -->\n      <article {{$s.HTMLAttributes}}>\n      {{if $s.Elem}}\n        <h3>{{$s.Title}}</h3>\n        {{range $s.Elem}}{{elem $.Template .}}{{end}}\n      {{else}}\n        <h2>{{$s.Title}}</h2>\n      {{end}}\n      <span class=\"pagenumber\">{{pagenum $s 1}}</span>\n      </article>\n  <!-- end of slide {{$s.Number}} -->\n  {{end}}{{/* of Slide block */}}\n\n      <article>\n        <h3>Thank you</h3>\n        {{range .Authors}}\n          <div class=\"presenter\">\n            {{range .Elem}}{{elem $.Template .}}{{end}}\n          </div>\n        {{end}}\n      </article>\n\n    </section>\n\n    <div id=\"help\">\n      Use the left and right arrow keys or click the left and right\n      edges of the page to navigate between slides.<br>\n      (Press 'H' or navigate to hide this message.)\n    </div>\n\n    {{if .PlayEnabled}}\n    <script src='{{static}}/jquery.js'></script>\n    <script src='{{static}}/jquery-ui.js'></script>\n    <script src='{{static}}/playground.js'></script>\n    <script src='{{static}}/play.js'></script>\n    <script>initPlayground(new HTTPTransport());</script>\n    {{end}}\n\n    <script>\n      (function() {\n        // Load Google Analytics tracking code on production site only.\n        if (window[\"location\"] && window[\"location\"][\"hostname\"] == \"talks.golang.org\") {\n          var ga = document.createElement(\"script\"); ga.type = \"text/javascript\"; ga.async = true;\n          ga.src = (\"https:\" == document.location.protocol ? \"https://ssl\" : \"http://www\") + \".google-analytics.com/ga.js\";\n          var s = document.getElementsByTagName(\"script\")[0]; s.parentNode.insertBefore(ga, s);\n        }\n      })();\n    </script>\n  </body>\n</html>\n{{end}}\n\n{{define \"newline\"}}\n<br>\n{{end}}\n" +
	"", etag: `"AP7WjjEhHaQ="`})