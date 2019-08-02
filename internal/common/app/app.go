package app

var ShutdownCallback []func()
func RegShutdownCallback(f func()) {
	ShutdownCallback = append(ShutdownCallback, f)
}
func RunShutdownCallback() {
	for _, fn := range ShutdownCallback {
		fn()
	}
}