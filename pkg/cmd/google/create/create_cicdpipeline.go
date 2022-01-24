package create

import (
	"fmt"

	"github.com/spf13/cobra"

	m "aip/pkg/cmd/google/models"

	"aip/pkg/cmd/google/services/cloudbuild"
	"aip/pkg/cmd/google/services/sourcerepo"
)

func NewCICDPipelineCommand() *cobra.Command {

	cicdpipelineCmd := &cobra.Command{
		Use:   "ci-cd-pipeline",
		Short: "Continuous Integration and Continuous Deployment pipeline.",
		Long: `This command allows you to create an entire CI/CD pipeline on Google Cloud Platform (GCP). You must provide the necessary files as parameters in order to create the desire pipeline The files can be provided in JSON or YAML extension.
		
		Example: 
			aip google create ci-cd-pipeline -c="config.yaml" -s="cloudbuild.yaml"
			aip google create ci-cd-pipeline --config="config.yaml" --steps="cloudbuild.yaml" 

			aip google create ci-cd-pipeline -c="config.json" -s="cloudbuild.json"
			aip google create ci-cd-pipeline --config="config.json" --steps="cloudbuild.json" `,

		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Creating pipeline...")

			fileName, _ := cmd.Flags().GetString("config")
			steps, _ := cmd.Flags().GetString("steps")

			pipeline := m.NewCiCdPipeline(fileName)

			project := pipeline.GetProject()
			project.SetNumber()
			projectId, projectNumber := project.GetId(), project.GetNumber()

			csr := pipeline.GetCsr()
			trigger := pipeline.GetTrigger()

			csrName, csrBranch, team := csr.GetCsrName(), csr.GetCsrBranch(), csr.GetCsrTeam()

			triggerName, triggerDescription := trigger.GetName(), trigger.GetDescription()

			sourcerepoResources := sourcerepo.NewSourceRepoResources(projectId, csrName)

			csrCfg := m.NewCSRConfig(fileName)

			newSourceRepository(sourcerepoResources)
			addDevsToRepo(sourcerepoResources, team)

			err := csrCfg.InitCSR()

			if csrCfg.HasTeam() {
				csrCfg.UpdateTeam()
			}

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("The repository was initialized successfully.")
			}

			cloudbuildResources := cloudbuild.NewCloudBuildTriggerResources(triggerName, triggerDescription, csrBranch, csrName, projectId, projectNumber, steps)

			addTrigger(cloudbuildResources)

		},
	}

	cicdpipelineCmd.PersistentFlags().StringP("config", "c", "", "Possible values: your-file.yaml, your-file.json")
	cicdpipelineCmd.MarkPersistentFlagRequired("config")

	cicdpipelineCmd.PersistentFlags().StringP("steps", "s", "", "Possible values: your-file.yaml, your-file.json")

	return cicdpipelineCmd
}

func newSourceRepository(sourcerepoResources sourcerepo.SourceRepoResources) {

	req, err := sourcerepoResources.FindByName()

	if err != nil {

		req, err = sourcerepoResources.AddRepository()

		if err != nil {
			fmt.Println("Error while creating the repository.")
		} else {
			fmt.Println("The repository was created sucessfully.")
		}

	} else {
		fmt.Println(req, err)
	}

}

func addDevsToRepo(sourcerepoResources sourcerepo.SourceRepoResources, team []string) {

	req, err := sourcerepoResources.AddDevelopers(team)

	if err != nil {
		fmt.Println(err, req)
	} else {
		fmt.Println("The developers were added sucessfully to the repository.")
	}
}

func addTrigger(cloudbuildResources cloudbuild.CloudBuildTriggerResources) {

	req, err := cloudbuildResources.AddTrigger()

	if err != nil {
		fmt.Println(req, err)
	} else {
		fmt.Println("The trigger was created sucessfully.")
	}

	err = cloudbuildResources.AuthorizeCloudBuildServiceAccount()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Cloud build service account is authorized to trigger deploys.")
	}
}
