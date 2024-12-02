package main

import (
	"fmt"
	"os/exec"
)

func main() {
	r := exec.Command("bash", "-c", `apt-get update && sudo apt-get upgrade && sudo apt-get install exa -y`)

	output, err := r.CombinedOutput()
	if err != nil {
		fmt.Printf("Exeс error: %v\n", err)
		fmt.Println(string(output))
		return
	}
	fmt.Println(string(output))
	fmt.Println("Обновленеи пакетов линукс завершено")
}
