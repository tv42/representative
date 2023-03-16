# `representative` -- Static slideshow generator for Go slides

`representative` takes a [Go
Present-style](https://pkg.go.dev/golang.org/x/tools/present)
`*.slide` (or `*.article`) file and creates a static HTML page.

## Install

```
$ go get eagain.net/go/representative/cmd/representative
```

## Usage

Read `foo.slide`, write `foo.html`, write assets into directory named
`static` (which will be created if needed) and assume the HTML can
refer to them as `static/...`:

```
$ representative -assets=static foo.slide
```

What files are created as assets is not guaranteed. We recommend
not trying to have other content in the same directory.

Pass any number of `*.slide` and `*.article` files.

There are two flags you can give it:

- `-assets=DIR`: Write assets to this directory.

  If not given, assets are not created at this time, and must be
  available to the HTML by some other means.

  This can be used with no files to just write assets.

- `-url-to-assets=URL`: Refers to assets from HTML with this base
  URL.

  If `-assets=DIR` was given, defaults to `DIR`, otherwise `static`.

Some more ways you can call it:

```
# Multiple slides at once
$ representative -assets=static *.slide

# Just write out the assets
$ representative -assets=static

# Just convert slides to HTML (assets assumed to be available as
# `static/...`)
$ representative foo.slide

# Write assets into custom directory name
$ representative -assets=my-asset-dir foo.slide

# Asset fetched from a subdir (default `static`), no trailing slash.
$ representative -url-to-assets=my-asset-dir foo.slide

# Assets fetched relative to arbitrary URL.
$ representative -url-to-assets=https://cdn.example.com/slide-assets/ foo.slide
```


## History

`representative` was inspired by https://github.com/cmars/represent
but is an independent implementation with completely different
command-line interface.

Notable differences:

- `representative` has new slideshow features added to `present` after
  `represent` was last updated
- `representative` does not need its source tree available to access
  assets
- `representative` does not copy everything in the source directory
  into the publish directory; it just writes HTML
