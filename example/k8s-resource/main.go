package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	kubeConfStr := "kubectl --kubeconfig ~/.kube/test.yaml "
	kubeNsStr := " -n tr "
	getRequestStr := "get deploy -o=custom-columns=name:.metadata.name,ns:.metadata.namespace,replicas:.spec.replicas,request-cpu:.spec.template.spec.containers[0].resources.requests.cpu,request-memory:.spec.template.spec.containers[0].resources.requests.memory"
	getReqCmd := fmt.Sprintf(kubeConfStr + getRequestStr + kubeNsStr)
	cmd := exec.Command("/bin/bash", "-c", getReqCmd)
	var getReqRes []byte
	var err error
	if getReqRes, err = cmd.Output(); err != nil {
		fmt.Println(err)
		// os.Exit(1)
	}
	for _, requestInfo := range strings.Split(string(getReqRes), "\n") {

		fmt.Println()
	}
}
