package cmd
import (
	"fmt"
	. "goAgenda/entity"
	"github.com/spf13/cobra"
	"errors"
)

const userPlace = "user.txt"
func userLegalCheck(userInfo []User,username string, password string,email string ,telphone string) (bool,error){
	if (username == "" || username[0] == '-'){
		return false,errors.New("The name is empty")
	}
	if (password == "" || password[0] == '-'){
		return false,errors.New("The password is empty")
	}
	if email == "" || email[0] == '-'{ 
		return false,errors.New("The mail is empty")
	}
	if telphone == ""{
		return false,errors.New("The telphone is empty")
	}
	for _,user := range userInfo {
		if user.Username == username{
			return false,errors.New("The name has already been registed")
		}
	}
	return true,nil
}
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "help a user to regist",
	Run: func(cmd *cobra.Command, args []string) {
		userInfo,userReadingerr := ReadUser(userPlace)
		if userReadingerr!=nil {
			fmt.Println(userReadingerr)
			return
		}
		username, _ := cmd.Flags().GetString("username")
		password, _:= cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		telphone,_ := cmd.Flags().GetString("telphone")
		if len(args)>0 {
			switch (args[0]){
				case "register":{
					if pass,err := userLegalCheck(userInfo,username,password,email,telphone); err!=nil{
						fmt.Println(err)
						return
					}else if !pass{
						fmt.Println("Regist Failed")
						return
					}
					userInfo = append(userInfo,User{username,password,email,telphone})
					WriteUser(userPlace,userInfo)
					fmt.Println("Regist success")
				}
				default:{
					fmt.Println("The command is undefined")
				}
			}
		}
	},
}
func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.Flags().StringP("username","u","","Help message for username")
	userCmd.Flags().StringP("password","p","","Help message for password")
	userCmd.Flags().StringP("email","e","","Help message for email")
	userCmd.Flags().StringP("telphone","t","","Help message for telphone")
}
