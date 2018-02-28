package fsutil

import "io"

// Close is intented to use like this
// f, err := os.Open("test.txt")
// if err != nil {
//     panic(err)
// }
// defer fsutil.Close(f, &err)
func Close(x io.Closer, p *error) {
	clsErr := x.Close()
	if *p == nil {
		*p = clsErr
	}
}

type SyncCloser interface {
	Sync() error
	Close() error
}

// SyncClose is intented to use like this
// f, err := os.Create("content.txt")
// if err != nil {
//     (panic)
// }
// f.WriteString("hello")
// defer fsutil.SyncClose(f, &err)
func SyncClose(x SyncCloser, p *error) {
	syncErr := x.Sync()
	if *p == nil {
		*p = syncErr
	}
	Close(x, p)
}
