package create

import (
	"fmt"

	"github.com/spf13/cobra"
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

			cfg, cloudbuild := setup(fileName)

			err := execProcess(cfg, cloudbuild)

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
