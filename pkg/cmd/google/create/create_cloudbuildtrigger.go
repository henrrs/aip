package create

import (
	"fmt"

	m "aip/pkg/cmd/google/models"

	"github.com/spf13/cobra"
	"google.golang.org/api/cloudbuild/v1"
)

func NewCloudBuildTrigger() *cobra.Command {

	cbtCmd := &cobra.Command{
		Use:   "csr",
		Short: "Cloud Build Trigger",
		Long: `This command allows you to create an Cloud Builg Trigger on Google Cloud Platform (GCP). You must provide the necessary configuration file as parameter in order to create the trigger. The file must be provided in JSON or YAML extension.
		
		Example: 
			aip google create cbt -c="config.yaml"
			aip google create cbt --config="config.yaml"

			aip google create cbt -c="config.json"
			aip google create cbt --config="config.json"  `,

		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Creating trigger...")

			fileName, _ := cmd.Flags().GetString("config")

			cfg, cloudbuild := setupCbt(fileName)

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

	return cbtCmd
}

func setupCbt(fileName string) (*m.TriggerConfig, *cloudbuild.BuildTrigger) {
	return nil, nil
}

func execCbtProcess(*m.TriggerConfig, *cloudbuild.BuildTrigger) error {

	return nil
}
