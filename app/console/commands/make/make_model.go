package make

import (
	"GoGinStarter/internal/utils"
	"github.com/gertd/go-pluralize"
	"github.com/spf13/cobra"
	"strings"
)

type MakeModelCommand struct{}

func (s *MakeModelCommand) RunE(cmd *cobra.Command, args []string) error {
	pClient := pluralize.NewClient()
	singular := pClient.Singular(args[0])

	modelName := utils.SFirstLetterUpper(singular)
	filename := strings.ToLower(singular)
	tableName := pClient.Plural(filename)

	values := map[string]interface{}{
		"MODEL_NAME": modelName,
		"TABLE_NAME": tableName,
	}

	return CreateFromStub(s.Stub()+"make_model.stub", "app/models/"+filename+".go", values)
}

func (s *MakeModelCommand) Stub() string {
	return StubPath
}

func (s *MakeModelCommand) NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "model",
		Short: "Generate a new model",
		Long:  "Generate a new gorm model struct",
		RunE:  s.RunE,
	}
}
