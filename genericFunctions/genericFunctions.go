package genericFunctions

// Check generic function check for errors and panic
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// This interface extends Close to zip archive
type myCloser interface {
	Close() error
}

// CloseFile function close the zip archive
func CloseFile(f myCloser) {
	err := f.Close()
	Check(err)
}
