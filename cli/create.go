package cli

import (
	"errors"
	"path"
	"runtime"
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

const (
	AppRoot        = "services"
	ControllersDir = "controllers"
	ModelsDir      = "models"
	RoutesDir      = "routes"
	ServicesDir    = "services"
	ValidationDir  = "validation"
	TemplateDir    = "../template"
	RootDir        = "github.com/JubaerHossain/gomd"
)

var AppName string

var Create = &cobra.Command{
	Use:  "create",
	Args: cobra.MinimumNArgs(2),
	RunE: Run,
}

func Run(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("not enough arguments")
	}
	AppName = args[0]
	name := args[1]
	fs := afero.NewBasePathFs(afero.NewOsFs(), AppRoot+"/")
	if err := createFolders(fs, name); err != nil {
		return err
	}
	if err := createFiles(fs, name); err != nil {
		return err
	}
	return nil
}

func createFolders(fs afero.Fs, name string) error {
	fs.Mkdir(name, 0755)
	dirs := []string{ControllersDir, ModelsDir, RoutesDir, ServicesDir, ValidationDir}
	for _, dir := range dirs {
		if err := fs.Mkdir(path.Join(name, dir), 0755); err != nil {
			return err
		}
	}
	return nil
}

func createFiles(fs afero.Fs, name string) error {
	createFile(fs, name, path.Join(TemplateDir, "controller.stub"), path.Join(name, ControllersDir, name+"_controller.go"))
	createFile(fs, name, path.Join(TemplateDir, "model.stub"), path.Join(name, ModelsDir, name+".go"))
	createFile(fs, name, path.Join(TemplateDir, "route.stub"), path.Join(name, RoutesDir, "api.go"))
	createFile(fs, name, path.Join(TemplateDir, "validation.stub"), path.Join(name, "validation", name+"_validation.go"))
	createFile(fs, name, path.Join(TemplateDir, "service.stub"), path.Join(name, ServicesDir, name+"_service.go"))
	return nil
}

func createFile(fs afero.Fs, name, stubPath, filePath string) error {
	fs.Create(filePath)

	_, filename, _, _ := runtime.Caller(1)
	stubPath = path.Join(path.Dir(filename), stubPath)

	contents, err := fileContents(stubPath)
	if err != nil {
		return err
	}
	contents = replaceStub(contents, name)

	if err := overwrite(AppRoot+"/"+filePath, contents); err != nil {
		return err
	}
	return nil
}

func fileContents(file string) (string, error) {
	a := afero.NewOsFs()
	contents, err := afero.ReadFile(a, file)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func overwrite(file string, message string) error {
	a := afero.NewOsFs()
	return afero.WriteFile(a, file, []byte(message), 0666)
}

func replaceStub(content string, name string) string {

	content = strings.Replace(content, "{{TitleName}}", Title(name), -1)
	content = strings.Replace(content, "{{PluralLowerName}}", Lower(Plural(name)), -1)
	content = strings.Replace(content, "{{SingularLowerName}}", Lower(Singular(name)), -1)
	content = strings.Replace(content, "{{AppName}}", RootDir, -1)
	content = strings.Replace(content, "{{AppRoot}}", AppRoot, -1)
	return content
}

func Plural(name string) string {
	pluralize := pluralize.NewClient()
	return pluralize.Plural(name)
}

func Singular(name string) string {
	pluralize := pluralize.NewClient()
	return pluralize.Singular(name)
}

func Lower(name string) string {
	return strings.ToLower(name)
}

func Title(name string) string {
	return strings.Title(Lower(name))
}
