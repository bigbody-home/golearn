package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string

	//fmt.Println(homedir.HomeDir())
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods("redis").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	for _, v := range pods.Items {
		fmt.Println(v.Name)
	}
	svc, err := clientset.CoreV1().Services("redis").List(context.TODO(), metav1.ListOptions{})
	for _, v := range svc.Items {
		fmt.Printf("redis的service是%v\n", v.Name)
	}
	//create namespace
	nsname := "dev"
	ns, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: nsname,
		},
	}, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("namespace %v create successfully!\n", ns.Name)
	err = clientset.CoreV1().Namespaces().Delete(context.TODO(), nsname, metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("namespace %v delete successfully!", nsname)
	// create pod
	lb := make(map[string]string)
	lb["nginx"] = "luke"
	fmt.Println(lb["nginx"])
	var num int32 = 1
	var replica *int32
	replica = &num
	var cport []v1.ContainerPort
	cport1 := v1.ContainerPort{
		Name:          "nginx-port",
		HostPort:      80,
		ContainerPort: 80,
		Protocol:      v1.ProtocolTCP,
	}
	cport = append(cport, cport1)
	container := v1.Container{
		Name:  "nginx",
		Image: "nginx",
		Ports: cport,
	}
	var containers []v1.Container
	containers = append(containers, container)
	deployes := appsV1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx-deployment",
			Namespace: "default",
			Labels:    lb,
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: replica,
			Selector: &metav1.LabelSelector{
				MatchLabels: lb,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nginx",
					Namespace: "default",
					Labels:    lb,
				},
				Spec: v1.PodSpec{
					Containers: containers,
				},
			},
		},
	}
	deploy, err := clientset.AppsV1().Deployments("default").Create(context.TODO(), &deployes, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("deployment %v create successfully\n", deploy.Name)

}
