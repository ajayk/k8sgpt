/*
Copyright 2023 The K8sGPT Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cache

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/k8sgpt-ai/k8sgpt/pkg/cache"
	"github.com/spf13/cobra"
)

var (
	region         string
	bucketName     string
	storageAccount string
	containerName  string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a remote cache",
	Long: `This command allows you to add a remote cache to store the results of an analysis.
	The supported cache types are:
	- Azure Blob storage
	- S3`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(color.YellowString("Adding remote based cache"))
		remoteCache := cache.NewCacheProvider(bucketname, region, storageAccount, containerName)
		err := cache.AddRemoteCache(remoteCache)
		if err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	CacheCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&region, "region", "r", "", "The region to use for the AWS S3 cache")
	addCmd.Flags().StringVarP(&bucketname, "bucket", "b", "", "The name of the AWS S3 bucket to use for the cache")
	addCmd.MarkFlagsRequiredTogether("region", "bucket")
	addCmd.Flags().StringVarP(&storageAccount, "storageacc", "s", "", "The Azure storage account name of the container")
	addCmd.Flags().StringVarP(&containerName, "container", "c", "", "The Azure container name to use for the cache")
	addCmd.MarkFlagsRequiredTogether("storageacc", "container")
	// Tedious check to ensure we don't include arguments from different providers
	addCmd.MarkFlagsMutuallyExclusive("region", "storageacc")
	addCmd.MarkFlagsMutuallyExclusive("region", "container")
	addCmd.MarkFlagsMutuallyExclusive("bucket", "storageacc")
	addCmd.MarkFlagsMutuallyExclusive("bucket", "container")
}
