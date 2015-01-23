package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Parity Control
// http://www.codeabbey.com/index/task_view/parity-control
func main() {
	fmt.Println(getSentence(os.Args[1:]))
}

func getSentence(in []string) string {
	buf := make([]byte, 0, len(in))

	for _, byt := range in {
		var tmp byte
		fmt.Sscanf(byt, "%d", &tmp)

		if dec, ok := checkByte(tmp); ok {
			buf = append(buf, dec)
		}
	}

	return string(buf)
}

func checkByte(b byte) (byte, bool) {
	bin := strconv.FormatInt(int64(b), 2)

	if strings.Count(bin, "1")%2 != 0 {
		return b, false
	}

	// Shift off the 8th bit, then shift back.
	// This works because byte is an alias for uint8,
	// so we know that we can safely shift.
	b = (b << 1) >> 1
	return b, true
}
