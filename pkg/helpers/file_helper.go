package helpers

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/yosa12978/webbpics/pkg/crypto"
)

func AddNewPicture(src io.Reader, filename string) (string, error) {
	fext := strings.Split(filename, ".")
	path := os.Getenv("MEDIA_DIR") + "/" + crypto.NewToken(16) + "." + fext[len(fext)-1]
	fmt.Println(path)
	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}
	return path[1:], err
}
