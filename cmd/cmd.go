package cmd

import(
	"fmt"
	"do/cluster"
	"github.com/spf13/cobra"
)

// InitCmd is initing commands
func InitCmd(){
	
	var fileName string
	
	 //Define Start-command to start a Master
	 var cmdStart = &cobra.Command{
		Use:   "start ",
		Short: "start master",
		Long:  `create is for starting master server to listen to cli`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cluster.StartMasterDeamon()
		},
	}

	 //Define Kill-command to stop a Master
	 var cmdKill = &cobra.Command{
		Use:   "stop ",
		Short: "stop master",
		Long:  `create is for stopping master server to listen to cli`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cluster.KillMasterDeamon()
		},
	}
	
    //Define Create-command to creat a cluster
	var cmdCreate = &cobra.Command{
		Use:   "create ",
		Short: "create demo cluster by yaml",
		Long:  `create is for creating demo cluster,then you can deploy apps`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
            if fileName == "" {
				fmt.Println("Error: Yaml file is needed")
				return
			}
			cluster.CreateCluster(fileName)
		},
	}
    cmdCreate.Flags().StringVarP(&fileName, "file", "f", "", "file to create the cluster")


    //Define Delete-command to delete a cluster
	var cmdDelete = &cobra.Command{
		Use:   "delete ",
		Short: "delete demo cluster ",
		Long:  `delete is for deleting demo cluster,then you can finish`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {

		//	createcluster.Delete()
		},
	}
 
	var rootCmd = &cobra.Command{Use: "do"}
	rootCmd.AddCommand(cmdKill,cmdStart,cmdCreate,cmdDelete)
	rootCmd.Execute()
}
