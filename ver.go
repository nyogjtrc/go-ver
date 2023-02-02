package ver

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

// Version info default as unknown
// will overwrite by ldflags
var (
	Version string = "unknown"
	BuildAt string = "unknown"
	Commit  string = "unknown"
)

// Print version info
func Print() {
	f := `  Version:	%s
  Go version:	%s
  Build at:	%s
  Commit:	%s
  OS/ARCH:	%s/%s
`

	fmt.Printf(
		f,
		Version,
		runtime.Version(),
		BuildAt,
		Commit,
		runtime.GOOS,
		runtime.GOARCH,
	)
}

func PrintOneLine() {
	fmt.Printf("%s build from %s on %s/%s at %s, %s\n",
		Version,
		Commit,
		runtime.GOOS,
		runtime.GOARCH,
		BuildAt,
		runtime.Version(),
	)
}

// Maps of version info
func Maps() map[string]string {
	return map[string]string{
		"version":    Version,
		"go_version": runtime.Version(),
		"build_at":   BuildAt,
		"commit":     Commit,
		"os/arch":    fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// Router for gin
func Router(r *gin.Engine) {
	r.GET("/api/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, Maps())
	})
}
