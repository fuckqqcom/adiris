package commons

import "fmt"

func CheckErr(err error, in interface{}) bool {
	if err != nil {
		//conf.Logger.Error("CheckErr", err)
		fmt.Println("error--->", err, in)
		return false
	}

	return true
}
