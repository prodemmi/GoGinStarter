package make

import (
	"GoGinStarter/internal/container"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var StubPath = "app/console/stubs/"

type MakeCommand struct{}

func (s *MakeCommand) NewCommand(container *container.Container) *cobra.Command {
	makeCommand := cobra.Command{
		Use:   "make",
		Short: "Make",
	}

	makeCommand.AddCommand((&MakeModelCommand{}).NewCommand())

	return &makeCommand
}

func CreateFromStub(stub string, target string, values map[string]interface{}) error {
	if _, err := os.Stat(target); err == nil {
		return fmt.Errorf("Target file " + path.Base(target) + " exists")
	}

	err := os.MkdirAll(path.Dir(target), os.ModePerm)
	if err != nil {
		panic(err)
	}

	stubContent, _ := ioutil.ReadFile(stub)

	for key, value := range values {
		stubContent = []byte(strings.ReplaceAll(string(stubContent), "{"+key+"}", value.(string)))
	}

	return ioutil.WriteFile(target, []byte(stubContent), 0644)
}
