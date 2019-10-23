package cmd
import (
	"fmt"
	"github.com/spf13/cobra"
)
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "help user to do something",
	Long: 
	`Usage:
		agenda [command]
	Available Commands:
		user : commands about user operation
	Use "agenda [command] --help" for more information about a command.
	`,
	
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1{
			if args[0] == "user" {
				fmt.Println(
`1. To add a register: 
	instruction:user register -u [Name] -p [Password] -e [Email] -t [Phonenumber]
	[Name] your name
	[Password] your password
	[Email] your email
	[Phonenumber] your phonenumber
2. To delete a register
	instruction:	user delete`)
			}
		} else if len(args) == 2{
			if args[0] == "user" {
				if args[1] == "register" {
					fmt.Println(
`Command : user register
Function: Regist a user.
instruction: agenda user register --name/-u [Name] --password/-p [Password] --email/-e [Email] --telphone/-t [Phonenumber]
Args :
	[Name] 	your name
	[Password] 	your password
	[Email] your email
	[Phonenumber] your phonenumber`)
				} 
			} 
		}
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}
