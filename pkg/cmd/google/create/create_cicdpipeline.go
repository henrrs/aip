package create

import (
	"fmt"

	"github.com/spf13/cobra"

	"aip/pkg/utils"

	"aip/pkg/services/google/cloudbuild"
	"aip/pkg/services/google/sourcerepo"
)

type Trigger struct {
	Name          string
	Description   string
	Branch        string
	Substitutions []string
}

type Repository struct {
	Name string
}

type Pipeline struct {
	ProjectId  string
	Team       []string
	Repository Repository
	Trigger    Trigger
}

type Config struct {
	Pipeline Pipeline
}

func NewConfig(fileName string) *Config {
	c := &Config{}
	c = utils.ReadFile(fileName, c).(*Config)
	c.Pipeline.Team = utils.UpdateTeam(c.Pipeline.Team)

	return c
}

func NewCICDPipelineCommand() *cobra.Command {

	cicdpipelineCmd := &cobra.Command{
		Use:   "ci-cd-pipeline",
		Short: "Continuous Integration and Continuous Deployment pipeline.",
		Long: `This command allows you to create an entire CI/CD pipeline on Google Cloud Platform (GCP). You must provide the necessary files as parameters in order to create the desire pipeline The files can be provided in JSON or YAML extension.
		
		Example: 
			aip google create ci-cd-pipeline -c="config.yaml" --p="cloudbuild.yaml"
			aip google create ci-cd-pipeline --config="config.yaml" --pipeline="cloudbuild.yaml" 

			aip google create ci-cd-pipeline -c="config.json" --p="cloudbuild.json"
			aip google create ci-cd-pipeline --config="config.json" --pipeline="cloudbuild.json" `,

		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Creating pipeline...")

			fileName, _ := cmd.Flags().GetString("config")
			cloudBuild, _ := cmd.Flags().GetString("pipeline")

			c := NewConfig(fileName)

			sourcerepoResources := sourcerepo.NewSourceRepoService(c.Pipeline.ProjectId, c.Pipeline.Repository.Name)

			newSourceRepository(sourcerepoResources)
			addDevsToRepo(sourcerepoResources, c.Pipeline.Team)

			err := sourcerepo.InitRepo(c.Pipeline.ProjectId, c.Pipeline.Repository.Name)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("The repository was initialized successfully.")
			}

			cloudbuildResources := cloudbuild.NewCloudBuildService(c.Pipeline.Trigger.Branch, c.Pipeline.Repository.Name, c.Pipeline.ProjectId, c.Pipeline.Trigger.Description, c.Pipeline.Trigger.Name, cloudBuild)

			addTrigger(cloudbuildResources)

		},
	}

	cicdpipelineCmd.PersistentFlags().StringP("config", "c", "", "Possible values: your-file.yaml, your-file.json")
	cicdpipelineCmd.MarkPersistentFlagRequired("config")

	cicdpipelineCmd.PersistentFlags().StringP("pipeline", "p", "", "Possible values: your-file.yaml, your-file.json")
	cicdpipelineCmd.MarkPersistentFlagRequired("pipeline")

	return cicdpipelineCmd
}

func newSourceRepository(sourcerepoResources sourcerepo.ServiceResources) {

	req, err := sourcerepo.FindByName(sourcerepoResources)

	if err != nil {

		req, err = sourcerepo.AddRepository(sourcerepoResources)

		if err != nil {
			fmt.Println("Error while creating the repository.")
		} else {
			fmt.Println("The repository was created sucessfully.")
		}

	} else {
		fmt.Println(req, err)
	}

}

func addDevsToRepo(sourcerepoResources sourcerepo.ServiceResources, team []string) {

	req, err := sourcerepo.AddDevelopers(sourcerepoResources, team)

	if err != nil {
		fmt.Println(err, req)
	} else {
		fmt.Println("The developers were added sucessfully to the repository.")
	}
}

func addTrigger(cloudbuildResources cloudbuild.ServiceResources) {

	req, err := cloudbuild.AddTrigger(cloudbuildResources)

	if err != nil {
		fmt.Println(req, err)
	} else {
		fmt.Println("The trigger was created sucessfully.")
	}

	err = cloudbuild.AuthorizeCloudBuildServiceAccount(cloudbuildResources)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Cloud build service account is authorized to trigger deploys.")
	}
}
