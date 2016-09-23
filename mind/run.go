package mind

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

// Run - run a Rachel command
func Run(name string) {
	commandFile := fmt.Sprintf("/rachel/%s.rachel", name)

	command, err := ioutil.ReadFile(commandFile)
	if err != nil {
		fmt.Println(err)
	}

	output, err := exec.Command("sh", "-c", string(command)).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}
