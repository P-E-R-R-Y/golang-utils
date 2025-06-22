package tools

// Retry calls fn up to attempts times until no error returned.
func Retry(attempts int, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
	}
	return err
}