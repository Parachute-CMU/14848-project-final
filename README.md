# 14848-Project Option 1-Final

## Contianer URLs
1. HADOOP: <br>
namenode: https://hub.docker.com/r/bde2020/hadoop-namenode<br>
datanode: https://hub.docker.com/r/bde2020/hadoop-datanode<br>
2. SPARK: https://hub.docker.com/r/bitnami/spark
3. Jupyter Notebook: https://hub.docker.com/r/jupyter/scipy-notebook
4. Sonarqube: https://hub.docker.com/_/sonarqube
5. Driver: https://hub.docker.com/r/parachute0719/driver

## Video demonstrating
[this is the video link :)](https://youtu.be/y-xqMuPL0RY)

## Part One
### Containerization in GCP
I firstly depployed the four applications containers in GCP. Specifically, **I deployed 1 namenode and 2 datanodes**.
![](https://github.com/Parachute-CMU/14848-project-final/blob/master/final/screenshots/step1/pods%20for%204%20containers.png)
To have the containers exposed, I used 4 load balancers.
![](https://github.com/Parachute-CMU/14848-project-final/blob/master/final/screenshots/step1/4%20load%20balancers.png)
For every container application, here is the routine for deployment.(used datanode for example)<br>
a) pull the image from docker hub: `docker pull bde2020/hadoop-datanode`<br>
b) tag the image: `docker tag bde2020/hadoop-datanode gcr.io/planar-berm-327519/xuan/hadoop-datanode:1.0`<br>
c) push the image to the container registry: `docker push gcr.io/planar-berm-327519/xuan/hadoop-datanode:1.0`<br>
d) deploy the k8s pods(set environment varaibles, modify yaml files for replica numbers, set pre-conditions in ymal files)<br>
e) deploy the service by exposing the pods with ports mapping<br>

check again for the Hadoop:<br>
1 namenode:
![](https://github.com/Parachute-CMU/14848-project-final/blob/master/final/screenshots/step1/1%20namenode.png)
2 datanodes:
![](https://github.com/Parachute-CMU/14848-project-final/blob/master/final/screenshots/step1/2%20datanodes.png)

### Part Two
### Build driver to connect the 4 applications
source code for my driver:(network programming in golang is concise :)
```golang
package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	// build the static main page
	go func() {
		fs := http.FileServer(http.Dir("assets/"))
		http.Handle("/hello/", http.StripPrefix("/hello/", fs))
		http.ListenAndServe(":8080", nil)
	} ()

	// driver is listening to port 80
	addr, _ := net.ResolveTCPAddr("tcp4", ":80")
	listener, err := net.ListenTCP("tcp", addr)
	defer listener.Close()
	if err != nil {
		return
	}

	containerRouting()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connection error")
		}
		conn.Close()
	}

}

func containerRouting() {
	// environment variables defined in driver YAML
	hadoop := os.Getenv("HADOOP") + ":9870"
	spark := os.Getenv("SPARK") + ":8080"
	jupyter := os.Getenv("JUPYTER") + ":8888"
	sonarqube := os.Getenv("SONARQUBE") + ":9000"

	// map the internal port of each container to an external IP
	go routingHelper(hadoop, ":81")
	go routingHelper(spark, ":82")
	go routingHelper(jupyter, ":83")
	go routingHelper(sonarqube, ":84")

}

func routingHelper(ip string, port string) {
	url, err := url.Parse(ip)
	if err != nil {
		return
	}
	http.ListenAndServe(port, httputil.NewSingleHostReverseProxy(url))
}
```
Instead of exposing the 4 containers as load balancers, I exposed the 4 contianers as Cluster IP, 
and connect them with the driver using the environment variables defined in the driver deployment yaml file.<br>
**Critical part in the driver deployment yaml file**
```
spec:
      containers:
        - image: parachute0719/driver:18.0
          imagePullPolicy: Always
          name: p1-driver
          env:
            - name: HADOOP
              value: "http://namenode-l24ms" // hadoop service name
            - name: SPARK
              value: "http://spark-zgm67" // spark service name
            - name: JUPYTER
              value: "http://jupyter-p7cm4" // jupyter service name
            - name: SONARQUBE
              value: "http://sonarqube-chf5c" // sonarqube service name
          ports:
            - containerPort: 8080 // static main page port
            - containerPort: 81 // hadoop port
            - containerPort: 82 // spark port
            - containerPort: 83 // jupyter port
            - containerPort: 84 // sonarqube port
```
exposing the dricver
![](https://github.com/Parachute-CMU/14848-project-final/blob/master/final/screenshots/step2/how%20to%20expose%20driver.png)
after exposing, driver has 5 exposed ports
![](https://github.com/Parachute-CMU/14848-project-final/blob/master/final/screenshots/step2/driver.png)
### Results
1 loadbalancer
![](https://github.com/Parachute-CMU/14848-project-final/blob/master/final/screenshots/step2/1%20load%20balancer.png)
the main page (at port 79/hello)
![](https://github.com/Parachute-CMU/14848-project-final/blob/master/final/screenshots/step2/mainpage79.png)
hadoop (at port 81)
![](https://github.com/Parachute-CMU/14848-project-final/blob/master/final/screenshots/step2/hadoop81.png)
spark (at port 82)
![](https://github.com/Parachute-CMU/14848-project-final/blob/master/final/screenshots/step2/spark82.png)
jupyter notebook (at port 83)
![](https://github.com/Parachute-CMU/14848-project-final/blob/master/final/screenshots/step2/jupyter83.png)
sonarqube (at port 84)
![](https://github.com/Parachute-CMU/14848-project-final/blob/master/final/screenshots/step2/sonarqube84.png)
