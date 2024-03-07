package messages

import "fmt"

// ERROR
const (
	ERR      = "\033[1;33m🚨🥶 ERROR:\033[0m "
	DIR      = "%sDirectory %s can't be created --->"
	FILE     = "%sFile %s can't be created"
	LANGINS  = "%s%s programming language is not installed on your computer(%s, %s)"
	LANGAV   = "%s%s programming language is not yet avaiable with gogol"
	TYPEAV   = "%sThe %s kind is not avaiable with golang"
	EXERR    = "%sCan't execut this commande: %s, output : %s"
	FETCHERR = "%sCan't fetch properly datas from : %s --->"
)

// This function  return an error when a directory can't be created
func ErrDir(name string) error {
	return fmt.Errorf(DIR, ERR, name)
}

// This function  return an error when a file can't be created
func ErrFile(name string) error {
	return fmt.Errorf(FILE, ERR, name)
}

// Function return an error when the requirements for the
// programming language selected by the user are not filled
func ErrLangInstall(lang, ops, arch string) error {
	return fmt.Errorf(LANGINS, ERR, lang, ops, arch)
}

// Function return an error when a programming
// language is not avaiable
func ErrLangAvaiable(lang string) error {
	return fmt.Errorf(LANGAV, ERR, lang)
}

// Function return error when a kind of app is not found
// (basic , webapp)
func ErrKindNotAv(kind string) error {
	return fmt.Errorf(TYPEAV, ERR, kind)
}

// Function return error when a command error happens
func ErrExec(cmd, out string) error {
	return fmt.Errorf(EXERR, ERR, cmd, out)
}

// Function for error Fetchin
func ErrFetch(url string) error {
	return fmt.Errorf(FETCHERR, ERR, url)
}
