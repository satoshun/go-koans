package koans

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)

const (
	intintint    int    = -1
	boolboolbool bool   = false
	stringstring string = ""
	deleteme     bool   = false
)

func TestKoans(t *testing.T) {
	aboutInterface()
	aboutStruct()
	aboutDefer()
	aboutTemplate()
	aboutRecover()
}

func assert(o bool) {
	if !o {
		fmt.Printf("\n%c[35m%s%c[0m\n\n", 27, __getRecentLine(), 27)
		os.Exit(1)
	}
}

func __getRecentLine() string {
	_, file, line, _ := runtime.Caller(2)
	buf, _ := ioutil.ReadFile(file)
	code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
	return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}
