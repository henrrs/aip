package create

import (
	"fmt"

	m "aip/pkg/cmd/google/models"
	"aip/pkg/cmd/google/services/cloudbuild"

	"github.com/spf13/cobra"
)

func NewCloudBuildTrigger() *cobra.Command {

	cbtCmd := &cobra.Command{
		Use:   "cbt",
		Short: "Cloud Build Trigger",
		Long: `This command allows you to create an Cloud Builg Trigger on Google Cloud Platform (GCP). You must provide the necessary configuration file as parameter in order to create the trigger. The file must be provided in JSON or YAML extension.
		
		Example: 
			aip google create cbt -c="config.yaml -s="cloudbuild.yaml""
			aip google create cbt --config="config.yaml" --steps="cloudbuild.yaml"

			aip google create cbt -c="config.json -s="cloudbuild.json""
			aip google create cbt --config="config.json" --steps="cloudbuild.json"`,

		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Creating trigger...")

			fileName, _ := cmd.Flags().GetString("config")
			steps, _ := cmd.Flags().GetString("steps")

			cfg, cloudbuild := setupCbt(fileName, steps)

			err := execCbtProcess(cfg, cloudbuild)

			if err != nil {
				fmt.Println("One or more errors have occurred during the process.")
			} else {
				fmt.Println("Process finished.")
			}

		},
	}

	cbtCmd.PersistentFlags().StringP("config", "c", "", "Possible values: your-file.yaml, your-file.json")
	cbtCmd.MarkPersistentFlagRequired("config")

	cbtCmd.PersistentFlags().StringP("steps", "s", "", "Possible values: your-file.yaml, your-file.json")

	return cbtCmd
}

func setupCbt(fileName, steps string) (*m.CbtConfig, cloudbuild.CloudBuildTriggerResources) {

	var cloudbuildResources cloudbuild.CloudBuildTriggerResources

	cbt := m.NewCbtConfig(fileName)

	csr := cbt.GetCsr()
	repoName, branchName := csr.GetCsrName(), csr.GetCsrBranch()

	project := cbt.GetProject()
	project.SetNumber()
	projectId, projectNumber := project.GetId(), project.GetNumber()

	trigger := cbt.GetTrigger()
	triggerName, triggerDescription := trigger.GetName(), trigger.GetDescription()

	if steps != "" {
		cloudbuildResources = cloudbuild.NewCloudBuildTriggerResources(triggerName, triggerDescription, branchName, repoName, projectId, projectNumber, steps)
	} else {
		cloudbuildResources = cloudbuild.NewCloudBuildTriggerResources(triggerName, triggerDescription, branchName, repoName, projectId, projectNumber, "")
	}

	return cbt, cloudbuildResources
}

func execCbtProcess(cfg *m.CbtConfig, cloudbuild cloudbuild.CloudBuildTriggerResources) error {

	cfg.NewCBT(cloudbuild)

	return nil
}
