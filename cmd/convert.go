package pipelines

import (
	"fmt"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	yaml "github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

// yaml "gopkg.in/yaml.v2"

type Branch struct {
	Default []Step `yaml:",flow"`
}

type Step struct {
	Step Script `yaml:",flow"`
}

type Script struct {
	Script []string `yaml:",flow"`
}

type T struct {
	Image     string
	Pipelines Branch `yaml:",flow"`
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Create a Docker file from a bitbucket-pipelines.yml",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Error("Expected a single bitbucket-pipelines.yml file as an argument")
		}

		data, err := ioutil.ReadFile(args[0])
		if err != nil {
			log.Error("Unable to read ", args[0])
			return
		}

		// j, err := yaml.YAMLToJSON(data)
		// if err != nil {
		// 	fmt.Printf("err: %v\n", err)
		// 	return
		// }
		// fmt.Println(string(j))

		t := T{}
		err = yaml.Unmarshal([]byte(data), &t)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		// fmt.Printf("--- t:\n%v\n\n", t)

		// d, err := yaml.Marshal(&t)
		// if err != nil {
		// 	log.Fatalf("error: %v", err)
		// }
		// fmt.Printf("--- t dump:\n%s\n\n", string(d))

		fmt.Printf("FROM %s\n", t.Image)
		fmt.Printf("ENV DEBIAN_FRONTEND noninteractive\n")
		fmt.Printf("RUN ")
		for i, cmd := range t.Pipelines.Default[0].Step.Script {
			fmt.Printf("%s", cmd)
			if i < len(t.Pipelines.Default[0].Step.Script)-1 {
				fmt.Printf(" && \\")
			}
			fmt.Println()

		}

	},
}

var outputPath string

func init() {
	convertCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "Dockerfile.pipelines", "output")
}
