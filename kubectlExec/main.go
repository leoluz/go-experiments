package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("kubectl", "exec", "-n", "ambassador", "-i", "deploy/ambassador", "--", "sh", "-c", "curl -s localhost:8005/snapshot-external | python -c 'import sys, json; print(json.load(sys.stdin)[\"AmbassadorMeta\"][\"cluster_id\"])'")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error: %+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", output)
}
