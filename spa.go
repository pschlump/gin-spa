package spa

import (
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/pschlump/dbgo"
)

type SPAList struct {
	FromTo map[string]string
	mutex  sync.Mutex
}

var logFilePtr *os.File

func InSpaList(s string, spa SPAList) (to string, found bool) {
	spa.mutex.Lock()
	defer spa.mutex.Unlock()
	to, found = spa.FromTo[s]
	return
}

func StaticServeMiddleware(urlPrefix, spaDirectory string, spa map[string]string, logFile *os.File) gin.HandlerFunc {
	var spalist SPAList
	spalist.FromTo = spa
	logFilePtr = logFile
	directory := static.LocalFile(spaDirectory, true)
	fileserver := http.FileServer(directory)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {
		if directory.Exists(urlPrefix, c.Request.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		} else if val, found := InSpaList(c.Request.URL.Path, spalist); found {
			dbgo.Fprintf(logFilePtr, "SPA Remap: from ->%s<- to ->%s<- at:%(LF)\n", c.Request.URL.Path, val)
			c.Request.URL.Path = val
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}

func ResetLogFile(newFp *os.File) {
	logFilePtr = newFp
}

/* vim: set noai ts=4 sw=4: */
