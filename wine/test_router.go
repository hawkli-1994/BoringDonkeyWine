package wine

import (
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T)  {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("TestParsePattern failed")
	}
}

func TestGetRoute(t *testing.T)  {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/lkrMacBook")
	if n == nil {
		t.Fatal("n == nil fail")
	}
	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}
	if ps["name"] != "lkrMacBook" {
		t.Fatal("name should be equal to 'lkrMacBook'")
	}


}
