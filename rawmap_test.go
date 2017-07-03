package rawmap

import "bytes"
import "fmt"
import "io/ioutil"
import "os"
import "testing"

func TestMmapSharedReadonly(t *testing.T) {
	test := func(msg []byte, start, length int) {
		tmpfile, err := ioutil.TempFile("", "rawmapTestMmapSharedReadonly")
		if err != nil {
			t.Fatalf("ioutil.TempFile: unexpected error: %s", err)
		}
		defer os.Remove(tmpfile.Name())
		defer tmpfile.Close()
		fmt.Fprintf(tmpfile, "%s", msg)
		b, err := MmapSharedReadonly(tmpfile, start, length)
		if err != nil {
			t.Fatalf("MmapSharedReadonly: unexpected error: %s", err)
		}
		if expected, got := msg[start:start+length], b; bytes.Compare(expected, got) != 0 {
			t.Errorf("expected %#v, but got %#v", string(expected), string(got))
		}
		if err := Munmap(b); err != nil {
			t.Fatalf("Munmap: unexpected error: %s", err)
		}
	}
	msg := bytes.Repeat([]byte("Hello, world!"), 1024)
	test(msg, 0, len(msg))
	test(msg, 4096, 4096)
}
