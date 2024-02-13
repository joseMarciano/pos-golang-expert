package main

import (
	"fmt"
	"github.com/joseMarciano/pos-golang-expert/apis/configs/authconfig"
	"github.com/joseMarciano/pos-golang-expert/apis/configs/databaseconfig"
	"github.com/joseMarciano/pos-golang-expert/apis/configs/serverconfig"
)

func init() {
	println("Main")
}

func main() {
	fmt.Println(databaseconfig.GetDbConfig())
	fmt.Println(authconfig.GetAuthConfig())
	fmt.Println(serverconfig.GetServerConfig())
}
