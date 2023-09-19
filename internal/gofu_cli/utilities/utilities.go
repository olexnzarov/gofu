package utilities

var exitCode = 0

func SetExitCode(code int) {
	exitCode = code
}

func GetExitCode() int {
	return exitCode
}
