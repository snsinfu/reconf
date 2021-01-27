package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/docopt/docopt-go"
)

const usage = `Generate files and run command.
Usage: reconf [-f -w <file> ...] <command>...

  <command>...   Command to execute.

Options:
  -w, --render <file>  Generate <file> (if it does not exist) by rendering
                       template file named "<file>.template".
  -f, --force          Force generating files, overwriting existing ones.
  -h, --help           Show this usage message and exit.
`

const (
	version        = "v0.1"
	errorCode      = 120
	templateSuffix = ".template"
)

type Config struct {
	Files   []string `docopt:"--render"`
	Force   bool     `docopt:"--force"`
	Command []string `docopt:"<command>"`
}

func main() {
	parser := docopt.Parser{
		OptionsFirst: true,
	}
	opts, err := parser.ParseArgs(usage, os.Args[1:], version)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := opts.Bind(&config); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(errorCode)
	}

	if err := run(config); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(errorCode)
	}
}

func run(config Config) error {
	envv := os.Environ()
	vars := map[string]interface{}{
		"env": mapEnviron(envv),
	}

	for _, filename := range config.Files {
		// Leave existing file as-is (unless forced).
		if _, err := os.Stat(filename); os.IsNotExist(err) || config.Force {
			if err := generate(filename, vars); err != nil {
				return err
			}
		}
	}

	paths := strings.Split(os.Getenv("PATH"), ":")

	return execvpe(config.Command[0], paths, config.Command, envv)
}

// Generates file by rendering corresponding template.
func generate(filename string, vars map[string]interface{}) error {
	tmplname := filename + templateSuffix
	tmpl := template.New(tmplname)

	// Custom functions must be set before parsing template.
	tmpl.Funcs(templateFuncs)

	// ParseFiles() uses basename of the file as the name of the template. We
	// want the path of the file as-is.
	text, err := ioutil.ReadFile(tmplname)
	if err != nil {
		return err
	}

	if _, err := tmpl.Parse(string(text)); err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := tmpl.Execute(file, vars); err != nil {
		os.Remove(filename)
		return err
	}

	return nil
}
