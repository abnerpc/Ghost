package main

import "fmt"
import "os/exec"

func main() {
	app := "ls"

	arg0 := "/etc/hosts"

	cmd := exec.Command(app, arg0)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Print(string(out))
}
