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

			pipeline, sourcerepo, cloudbuild := setupCiCd(fileName, steps)

			err := execCiCdProcess(pipeline, sourcerepo, cloudbuild)

			if err != nil {
				fmt.Println("One or more errors have occurred during the process.")
			} else {
				fmt.Println("Process finished.")
			}

		},
	}

	cicdpipelineCmd.PersistentFlags().StringP("config", "c", "", "Possible values: your-file.yaml, your-file.json")
	cicdpipelineCmd.MarkPersistentFlagRequired("config")

	cicdpipelineCmd.PersistentFlags().StringP("steps", "s", "", "Possible values: your-file.yaml, your-file.json")

	return cicdpipelineCmd
}

func setupCiCd(fileName, steps string) (*m.CiCdPipelineConfig, sourcerepo.SourceRepoResources, cloudbuild.CloudBuildTriggerResources) {
	pipeline := m.NewCiCdPipeline(fileName)

	project := pipeline.GetProject()
	project.SetNumber()
	projectId, projectNumber := project.GetId(), project.GetNumber()

	csr := pipeline.GetCsr()
	trigger := pipeline.GetTrigger()

	csrName, csrBranch := csr.GetName(), csr.GetBranch()

	csrCfg := m.NewCSRConfigWithoutParameters()

	csrCfg.SetCsr(csr)
	csrCfg.SetProject(project)

	cbtCfg := m.NewCBTConfigWithoutParameters()

	cicdpipelineCfg := m.NewCiCdPipelineConfig(*csrCfg, *cbtCfg)

	triggerName, triggerDescription := trigger.GetName(), trigger.GetDescription()

	sourcerepoResources := sourcerepo.NewSourceRepoResources(projectId, csrName)

	cloudbuildResources := cloudbuild.NewCloudBuildTriggerResources(triggerName, triggerDescription, csrBranch, csrName, projectId, projectNumber, steps)

	return cicdpipelineCfg, sourcerepoResources, cloudbuildResources

}

func execCiCdProcess(pipelineCfg *m.CiCdPipelineConfig, sourcerepo sourcerepo.SourceRepoResources, cloudbuild cloudbuild.CloudBuildTriggerResources) error {

	csrCfg := pipelineCfg.GetCSRConfig()
	cbtCfg := pipelineCfg.GetCBTConfig()

	csrCfg.NewCSR(sourcerepo)

	err := csrCfg.InitCSR()

	if err != nil {
		fmt.Println("An error has occurred during CSR initialization.")
		fmt.Println(err)
	} else {
		fmt.Println("CSR was initialized sucessfuly.")

		if csrCfg.HasTeam() {

			csrCfg.UpdateTeam()

			err = csrCfg.AddTeam(sourcerepo)

			if err != nil {
				return err
			}
		}
	}

	err = cbtCfg.NewCBT(cloudbuild)

	if err != nil {
		return err
	}

	return nil
}
