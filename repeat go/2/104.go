package main

import (
	"fmt"
	"html"
	"strings"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ct := resp.Header.Get("content=Type")
	if ct != "text/html" && !strings.HuaPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s,not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML:%v", url, err)
	}
	return nil
}
