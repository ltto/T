package str

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5(str string) string {
	w := md5.New()
	_, _ = io.WriteString(w, str)
	return fmt.Sprintf("%x", w.Sum(nil))
}
