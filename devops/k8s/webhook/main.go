package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"k8s.io/api/admission/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	runtimeScheme = runtime.NewScheme()
	codecs        = serializer.NewCodecFactory(runtimeScheme)
	deserializer  = codecs.UniversalDeserializer()
	defaulter     = runtime.ObjectDefaulter(runtimeScheme)
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	nsList, err := clientset.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})

	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", nsList)
}

func g() {
	g := gin.New()
	//https://github.com/yaoice/webhook-demo/blob/master/pkg/webhook/webhook.go
	g.GET("/validate", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		requestedAdmissionReview := v1beta1.AdmissionReview{}

		deserializer := codecs.UniversalDeserializer()
		deserializer.Decode(body, nil, &requestedAdmissionReview)

	})
}
