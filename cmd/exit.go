package cmd

import (
	"fmt"
	"lightyear/core/global"
)

func Destroy() {

	fmt.Println("Destroying app...")
	defer global.AsynqClient.Close()

}
