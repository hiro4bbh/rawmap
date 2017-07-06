package rawmap

import "bytes"
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
		if _, err := tmpfile.Write(msg); err != nil {
			t.Fatalf("File.Write: unexpected error: %s", err)
		}
		b, err := MmapShared(tmpfile, start, length, PROTMODE_READONLY)
		if err != nil {
			t.Fatalf("MmapShared: unexpected error: %s", err)
		}
		if expected, got := msg[start:start+length], b; bytes.Compare(expected, got) != 0 {
			t.Errorf("expected %#v, but got %#v", string(expected), string(got))
		}
		if err := Munmap(b); err != nil {
			t.Fatalf("Munmap: unexpected error: %s", err)
		}
	}
	msg := bytes.Repeat([]byte("Hello, world!"), 1024*1024)
	test(msg, 0, len(msg))
	test(msg, 65536, 4096)
}

func TestMmapSharedReadWrite(t *testing.T) {
	test := func(buf, msg []byte, start, length int) {
		tmpfile, err := ioutil.TempFile("", "rawmapTestMmapSharedReadWrite")
		if err != nil {
			t.Fatalf("ioutil.TempFile: unexpected error: %s", err)
		}
		defer os.Remove(tmpfile.Name())
		defer tmpfile.Close()
		if _, err := tmpfile.Write(buf); err != nil {
			t.Fatalf("File.Write: unexpected error: %s", err)
		}
		b, err := MmapShared(tmpfile, start, length, PROTMODE_READWRITE)
		if err != nil {
			t.Fatalf("MmapShared: unexpected error: %s", err)
		}
		copy(b, msg)
		if err := Munmap(b); err != nil {
			t.Fatalf("Munmap: unexpected error: %s", err)
		}
		b, err = MmapShared(tmpfile, start, length, PROTMODE_READWRITE)
		if err != nil {
			t.Fatalf("MmapShared: unexpected error: %s", err)
		}
		if expected, got := msg, b[:len(msg)]; bytes.Compare(expected, got) != 0 {
			t.Errorf("expected %#v, but got %#v", string(expected), string(got))
		}
		if err := Munmap(b); err != nil {
			t.Fatalf("Munmap: unexpected error: %s", err)
		}
	}
	buf := make([]byte, 1024*1024)
	msg := []byte("Hello, world!")
	test(buf, msg, 0, len(buf))
	test(buf, msg, 65536, 4096)
}
