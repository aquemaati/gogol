package messages

import "fmt"

const (
	STARTING = "\n🚀🚀 \033[1;32mStarting a new %s application named %s with %s programming language\033[0m\n"
	// Creating messages
	BUILDFILE = "🪚 Building the new file: %s"
	BUILDDIR  = "🚧 Building the new directory: %s"

	CONGRAT = "\n💪 Congratulation, your %s application named %s with %s programming language has been created successfully"

	// Action
	EXECCMD = "🔫 Executing command: %s\n%s"
	SELECT  = "Select an option or exit gogol with CTRL-c\n"
	CHECKOS = "\n🧐 Checking if your computer has the requirements for %s programming language"

	// Fetching message
	FETCH = "📡 \033[34mFetching datas from: %s\033[0m"
)

// TODO Select option
func SlectOpt() string {
	return SELECT
}

// Function return string to say binary is checking computer
// For programming languages requirements
func CheckLangMess(lang string) string {
	return fmt.Sprintf(CHECKOS, lang)
}

// Function that say a new projectr is created
func StartingProject(kind, name, lang string) string {
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
