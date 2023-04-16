# Single Page App File Service

Often Single Page Apps (SPA) that keep history use the URl path to track the
location in the application.   When people refresh a page or use a bookmark
in the browser this leads to paths like `https://example.com/login` that should
serve the `index.html` page but instead result in 404 errors.   

This middleware allows for a list of different paths that can be remapped to
`index.html` or other files.   Files that are legitimately missing will still
be returned as 404 errors - but items in the list will remapped to a specified
file (usually `index.html`).

## example

```
package main

import (
	"os"

	"github.com/gin-gonic/gin"
	spa "github.com/pschlump/gin-spa"
)

func main() {
	r := gin.Default()
	pathsToRemap = map[string]string{
		"/login": "/index.html",
		"/logout": "/index.html",
	}
	r.Use(spa.StaticServeMiddleware("/", "./build", pathsToRemap, os.Stderr)) // your build of React or other SPA
	r.Run()
}
```
