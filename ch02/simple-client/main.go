// Copyright 2019 Yandy Ramirez
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// set default kubeconfig path
	home, _ := os.UserHomeDir()
	cfgpath := filepath.Join(home, ".kube", "config")

	// parse flags for kubeconfig path if not set use default
	kubeconfig := flag.String("kubeconfig", cfgpath, "path to kubeconfig file")
	flag.Parse()

	// get a rest.Config from the kubeconfig path
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("Cannot load kubeconfig %v\n", err.Error())
	}

	// generate a clientset for the given config
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Cannot create new config %v\n", err.Error())
	}

	// pod, err := clientset.CoreV1().Pods("book").Get("example", metav1.GetOptions{})
	dep, err := clientset.AppsV1().Deployments("kube-system").Get("coredns", metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Cannot get %v\n", err.Error())
	}

	// fmt.Println(pod)
	fmt.Println(dep.GetCreationTimestamp())
}
