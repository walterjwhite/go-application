package main

import (
	"github.com/walterjwhite/go-application/libraries/application"
	"github.com/walterjwhite/go-application/libraries/foreachfile"
	"github.com/walterjwhite/go-application/libraries/logging"
	"github.com/walterjwhite/go-application/libraries/periodic"
	"github.com/walterjwhite/go-application/libraries/runner"

	"flag"
	"strings"
)

var (
	rootDirectoryFlag = flag.String("RootDirectory", ".", "Root Directory to scan files")
	intervalFlag      = flag.String("Interval", "1m", "Interval between execution")
	patternStringFlag = flag.String("Patterns", "", "Patterns")
	commandFlag       = flag.String("Cmd", "", "Command")
	argumentsFlag     = flag.String("Arguments", "", "Arguments")
)

func init() {
	application.Configure()
}

func main() {
	periodic.Periodic(application.Context, periodic.GetInterval(*intervalFlag), runIteration)

	application.Wait()
}

func runIteration() error {
	foreachfile.Execute(*rootDirectoryFlag, exec, getPatterns()...)

	return nil
}

func getPatterns() []string {
	if len(*patternStringFlag) > 0 {
		return strings.Split(*patternStringFlag, ",")
	}

	return []string{}
}

func exec(filePath string) {
	var cmd string
	var arguments []string

	if len(*commandFlag) > 0 {
		cmd = *commandFlag

		if len(*argumentsFlag) > 0 {
			arguments = append(arguments, strings.Split(*argumentsFlag, ",")...)
			arguments = append(arguments, filePath)
		}
	} else {
		cmd = filePath
	}

	_, err := runner.Run(application.Context, cmd, arguments...)
	logging.Panic(err)
}