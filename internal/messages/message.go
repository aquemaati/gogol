package messages

import "fmt"

const (
	// Creating messages
	STARTING  = "🚀🚀 \033[1;32mStarting a new %s application named %s with %s programming language\033[0m\n"
	BUILDFILE = "🪚 Building the new file: %s\n"
	BUILDDIR  = "🚧 Building the new directory: %s\n"
	CONGRAT   = "💪 Congratulation, your %s application named %s with %s programming language has been created successfully \n"
	EXECCMD   = "🔫 Executing command: %s\n%s\n"
	SELECT    = "Selec an option or exit gogol with CTRL-c\n"

	// Fetching message
	FETCH = "📡\033[34mFetching datas from: %s\033[0m\n"
)

// TODO Select option
func SlectOpt() string {
	return SELECT
}

// Function that say a new projectr is created
func StartingProject(name, kind, lang string) string {
	return fmt.Sprintf(STARTING, kind, name, lang)
}

// This function return a message when a file is being build
func FileBuilding(name string) string {
	return fmt.Sprintf(BUILDFILE, name)
}

// This function return a message when a directory is being build
func DirBuilding(name string) string {
	return fmt.Sprintf(BUILDDIR, name)
}

// This function return a message when fetching datas
func Fetching(target string) string {
	return fmt.Sprintf(FETCH, target)
}

// Tell the user everything went ok
func ExecCmd(cmd, out string) string {
	return fmt.Sprintf(EXECCMD, cmd, out)
}

// Tell the user everything went ok
func Congrat(kind, name, lang string) string {
	return fmt.Sprintf(CONGRAT, kind, name, lang)
}
