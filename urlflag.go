package main

import (
	"flag"
	"net/url"
)

// URLFlag is a command-line flag that must be a URL. Additionally,
// if the URL has not been set, the default value is derived from a
// callback.
type URLFlag struct {
	URL     *url.URL
	Default func() *url.URL
}

func (f *URLFlag) GetURL() *url.URL {
	if u := f.URL; u != nil {
		return u
	}
	if fn := f.Default; fn != nil {
		u := fn()
		return u
	}
	return nil
}

var _ flag.Value = (*URLFlag)(nil)

func (f *URLFlag) String() string {
	if u := f.GetURL(); u != nil {
		return u.String()
	}
	return ""
}

func (f *URLFlag) Set(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return err
	}
	f.URL = u
	return nil
}
