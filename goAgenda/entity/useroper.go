package entity

import(
	"io/ioutil"
	"encoding/json"
	"os"
)

type User struct{
	Username string `json:"username"`
    Password string `json:"password"`
    Email string `json:"email"`
    Telphone string `json:"telphone"`
}
func ReadUser(filePath string) ([]User,error) {
	var user []User
	s,err := ioutil.ReadFile(filePath)
	if err!=nil {
		return user,err
	}
	jsonStr := string(s)
	
	json.Unmarshal([]byte(jsonStr),&user)
	return user,nil
}

func WriteUser (filePath string, user []User) error{
	if data,err:=json.Marshal(user);err == nil{
		return ioutil.WriteFile(filePath,[]byte(data),os.ModeAppend)
	} else{
		return err
	}
}