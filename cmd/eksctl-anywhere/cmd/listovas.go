package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"sigs.k8s.io/yaml"

	eksav1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/version"
)

type listOvasOptions struct {
	fileName string
}

type listOvasOutput struct {
	URI    string
	SHA256 string
	SHA512 string
}

var listOvaOpts = &listOvasOptions{}

func init() {
	listCmd.AddCommand(listOvasCmd)
	listOvasCmd.Flags().StringVarP(&listOvaOpts.fileName, "filename", "f", "", "Filename that contains EKS-A cluster configuration")
	err := listOvasCmd.MarkFlagRequired("filename")
	if err != nil {
		log.Fatalf("Error marking filename flag as required: %v", err)
	}
}

var listOvasCmd = &cobra.Command{
	Use:          "ovas",
	Short:        "List the OVAs that are supported by current version of EKS Anywhere",
	Long:         "This command is used to list the vSphere OVAs from the EKS Anywhere bundle manifest for the current version of the EKS Anywhere CLI",
	PreRunE:      preRunListOvasCmd,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := listOvas(cmd.Context(), listOvaOpts.fileName); err != nil {
			return err
		}
		return nil
	},
}

func listOvas(context context.Context, spec string) error {
	clusterSpec, err := readAndValidateClusterSpec(spec, version.Get())
	if err != nil {
		return err
	}

	bundle := clusterSpec.VersionsBundle

	for _, ova := range bundle.Ovas() {
		if strings.Contains(ova.URI, string(eksav1alpha1.Bottlerocket)) {
			fmt.Printf("%s:\n", strings.Title(string(eksav1alpha1.Bottlerocket)))
		} else {
			fmt.Printf("%s:\n", strings.Title(string(eksav1alpha1.Ubuntu)))
		}
		output := listOvasOutput{
			URI:    ova.URI,
			SHA256: ova.SHA256,
			SHA512: ova.SHA512,
		}
		yamlOutput, err := yaml.Marshal(output)
		if err != nil {
			return err
		}
		fmt.Println(yamlIndent(2, string(yamlOutput)))
	}

	return nil
}

func preRunListOvasCmd(cmd *cobra.Command, args []string) error {
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		err := viper.BindPFlag(flag.Name, flag)
		if err != nil {
			log.Fatalf("Error initializing flags: %v", err)
		}
	})
	return nil
}

func yamlIndent(level int, yamlString string) string {
	indentation := strings.Repeat(" ", level)
	indentedString := fmt.Sprintf("%s%s", indentation, strings.Replace(yamlString, "\n", "\n"+indentation, -1))
	return indentedString
}
