# go-rod-issue

Sample code that makes go-rod fail when utilizing
[luabagg/orcgen](https://github.com/luabagg/orcgen) to generate PDF files from
HTML.

## Usage

```
$ go run main.go
2024/10/15 11:57:04 Server started.
2024/10/15 11:57:06 Server started.
panic: navigation failed: net::ERR_CONNECTION_REFUSED

goroutine 1 [running]:
github.com/go-rod/rod/lib/utils.init.func2({0x805100?, 0xc0000258b0?})
	/home/manuel/workbench/go/pkg/mod/github.com/go-rod/rod@v0.116.0/lib/utils/utils.go:69 +0x1d
github.com/luabagg/orcgen/v2/pkg/webdriver.(*WebDriver).Connect.New.(*Browser).WithPanic.genE.func1({0xc000025a30?, 0x15?, 0x0?})
	/home/manuel/workbench/go/pkg/mod/github.com/go-rod/rod@v0.116.0/must.go:36 +0x62
github.com/go-rod/rod.(*Browser).MustPage(0xc0000b9e60, {0xc0000e5d70?, 0x77093a?, 0xc000012750?})
	/home/manuel/workbench/go/pkg/mod/github.com/go-rod/rod@v0.116.0/must.go:71 +0xd2
github.com/luabagg/orcgen/v2/pkg/webdriver.(*WebDriver).UrlToPage(0xc0000a8180?, {0x8a86c4?, 0xc000100008?})
	/home/manuel/workbench/go/pkg/mod/github.com/luabagg/orcgen/v2@v2.0.2/pkg/webdriver/webdriver.go:60 +0x32
github.com/luabagg/orcgen/v2.ConvertWebpage[...]({0xba2098?, 0xc0000255c0}, {0x8a86c4, 0x15})
	/home/manuel/workbench/go/pkg/mod/github.com/luabagg/orcgen/v2@v2.0.2/orcgen.go:88 +0x89
main.main()
	/home/manuel/workbench/work/go-rod-issue/main.go:81 +0x3f7
exit status 2
```
