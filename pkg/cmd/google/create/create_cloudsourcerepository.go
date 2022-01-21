package create

import (
	"fmt"

	"github.com/spf13/cobra"

	m "aip/pkg/cmd/google/models"
	"aip/pkg/services/google/sourcerepo"
)

func NewCloudSourceRepository() *cobra.Command {

	csrCmd := &cobra.Command{
		Use:   "csr",
		Short: "Cloud Source Repository",
		Long: `This command allows you to create an Cloud Source Repository (CSR) on Google Cloud Platform (GCP). You must provide the necessary configuration file as parameter in order to create the repository. The file must be provided in JSON or YAML extension.
		
		Example: 
			aip google create csr -c="config.yaml"
			aip google create csr --config="config.yaml"

			aip google create csr -c="config.json"
			aip google create csr --config="config.json"  `,

		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Creating repository...")

			fileName, _ := cmd.Flags().GetString("config")

			cfg, sourcerepo := setupCsr(fileName)

			err := execCsrProcess(cfg, sourcerepo)

			if err != nil {
				fmt.Println("One or more errors have occurred during the process.")
			} else {
				fmt.Println("Process finished.")
			}

		},
	}

	csrCmd.PersistentFlags().StringP("config", "c", "", "Possible values: your-file.yaml, your-file.json")
	csrCmd.MarkPersistentFlagRequired("config")

	return csrCmd
}

func setupCsr(fileName string) (*m.CSRConfig, sourcerepo.SourceRepoResources) {

	cfg := m.NewCSRConfig(fileName)

	csrName := cfg.GetName()
	project := cfg.GetProject()
	projectId := project.GetId()

	fmt.Println(csrName)

	sourcerepo := sourcerepo.NewSourceRepoResources(projectId, csrName)

	return cfg, sourcerepo
}

func execCsrProcess(cfg *m.CSRConfig, sourcerepo sourcerepo.SourceRepoResources) error {

	cfg.NewCSR(sourcerepo)

	err := cfg.InitCSR()

	if err != nil {
		fmt.Println("An error has occurred during CSR initialization.")
		fmt.Println(err)
	} else {
		fmt.Println("CSR was initialized sucessfuly.")

		if cfg.HasTeam() {

			cfg.UpdateTeam()

			team := cfg.GetTeam()

			req, err := sourcerepo.AddDevelopers(team)

			if err != nil {
				fmt.Println(err, req)

				return err

			} else {
				fmt.Println("The developers were added sucessfully to the repository.")
			}
		}
	}

	return nil
}
