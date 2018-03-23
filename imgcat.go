package imgcat

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

type ImageCat struct {
	bf bytes.Buffer
}

func (i *ImageCat) Write(p []byte) (n int, err error) {
	return i.bf.Write(p)
}

func (i *ImageCat) Cat() {
	TERM := os.Getenv("TERM")
	encoded := base64.StdEncoding.EncodeToString(i.bf.Bytes())

	_IMG_PRE := "\033]"
	_IMG_POST := "\a"

	if !strings.HasPrefix(TERM, "screen") {
		_IMG_PRE = "\033Ptmux;\033\033]"
		_IMG_POST = "\a\033\\\n"
	}
	inline := 1
	preserveAspectRatio := 0
	data := fmt.Sprintf("\n%s1337;File=inline=%d;preserveAspectRatio=%d:%s%s",
		_IMG_PRE, inline, preserveAspectRatio, encoded, _IMG_POST)
	fmt.Println("=================")
	bio := bufio.NewWriter(os.Stdout)
	bio.WriteString(data)
	bio.Flush()

	return
}
