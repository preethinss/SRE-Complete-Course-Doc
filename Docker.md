# Docker architecture
Docker client : provides commands to communicate with docker host
Docker host: Docker Deamon(high level runtime) is the main process which is running and helps in creating docker images and containers by communicate with low level runtime containerd
Docker registry: central place to store docker images
    popular public registry: docker hub

# Docker lifecycle
1. application source code
2. write Dockerfile
3. Build docker image from Dockerfile using docker build command
4. Run container from docker image using dokcer run command

# Docker installation as root user
apt update
apt install docker.io -y

# Docker installation as ubuntu user:
sudo apt update
sudo apt install docker.io -y
docker run hello-world
usermod -aG docker ubuntu

Docker basic commands:
1. docker pull and docker create
docker pull nginx
Docker create nginx
Docker start nginx
2. docker run - with name & without name
docker run -d --name sample-nginx -p 8082:80 nginx:latest
docker run -d sample-nginx -p 8083:80 nginx:latest
3. docker ps - running and stopped container
docker ps - list of running container
Docker ps -a - list of stooped and running container
4. docker images
5. docker start and docker stop
6. docker rm- remove container
7. docker rmi- remove image
8. docker run - detach mode and non-detach mode
9. docker inspect container_name

# network
ip addr
docker network ls

# bridge - default network
docker run -d --name bridge-conatiner -p 8082:80 nginx     (or)
docker run -d --name bridge-conatiner -p 8082:80 --network bridge nginx

# host
docker run -d --name bridge-conatiner -p 8082:80 --network host nginx

# custom network: provides isolation between container 
# containers are in bridge network cannot communicate with custom network because both are in diffrent subnet 
Docker network --help
docker network create mynet - custom network
docker run -d --net=mynet nginx:latest (or) 
docker run -d --network mynet -p 8083:80 nginx:latest
docker network rm mynet or docker network prune

# check containers in custom network reachable from bridge network
1. docker exec -it bridge_container_id /bin/bash
2. Install ping: apt-get install iputils-ping
3. ping Ip_address_custom_container

# check containers are reachable in same network 
1. docker exec -it bridge_container_id /bin/bash
2. ping install
3. ping Ip_address_bridge_container

----------------------------------------------------------------------
# volume
# before volume scenario:
1. when containers are deleted or crashed due to any reason the data associated with conatiner is also deleted
# why volume?- conatiners are ephemeral in nature(short-lived) and once containers are deleted, application data is stored in container gets deleted 
# volume features:
1. Decoupling container from stoarge
2. share volume among multiple container

two types of volumes: named and anonymous
1. anonymous : the volume name given by docker
2. named: you can customize name for volume and its easily addresable

commands:
1. docker volume --help
2. docker volume ls

Note: by default no volume will be there

# anonymous volume
1. docker run -d --name nginx_anonymous -v /xyz nginx:latest

/xyz is container path - you can provide existing path for your application where data gets stored ex: mongo: /data/db

# names volume
1. docker run -d --name nginx_named -v log:/xyz nginx:latest

log is a volume and /xyz is container path

# inspect volume: get volume details(provide info on where volume is created)
docker volume inspect <volume_name>

Note: when you create volumes using above commands try to delete container and see if data persists and share that volume to new container and see data is there is new container path(old+new data)

# delete volume
# delete anonymous volume: only unused volume
docker volume prune 

Note: if volume is used by any container you cannot delete volume

# named volume
docker volume rm <volume_name>

--------------------------------------------------------------------------
# Dockerfile
# without multistage
1. Create go application
    steps:
        go mod init sample
        vi main.go

package main

import ( 
    "fmt"
    "net/http"
)
func helloDocker(w http.ResponseWriter, r *http.Request)
{ 
    fmt.Fprintf(w, "Hello, go!")
}
func main() {
    http.HandleFunc("/", helloDocker)
    fmt.Println("Server running on port 8080")
    http.ListenAndServe(":8080", nil)
}

# copy and paste above program into main.go

2. Create Dockerfile and copy paste below content
FROM golang:latest
WORKDIR /app
COPY . . 
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]

3. build image from Dockerfile
docker build -t <repository>/<image_name>:<tag>
ex:
    docker build -t preethi/without_multi_stage:1.0.0 .

Note: execute docker build command where Dockerfile is present 

4. push image to docker hub
docker push preethi/without_multi_stage:1.0.0

# with multistage
# why multistage? - reduce size of image, provides security

1. Create golang application
 steps:
        go mod init sample
        vi main.go

package main

import ( 
    "fmt"
    "net/http"
)
func helloDocker(w http.ResponseWriter, r *http.Request)
{ 
    fmt.Fprintf(w, "Hello, go!")
}
func main() {
    http.HandleFunc("/", helloDocker)
    fmt.Println("Server running on port 8080")
    http.ListenAndServe(":8080", nil)
}

# copy and paste above program into main.go

2. Create Dockerfile and copy paste below content
# build stage
FROM golang:latest
WORKDIR /app
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# final stage
FROM scratch
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 8080
CMD ["./main"]

Note: Scratch is minimalistic image which provides security at container level

3. build image from Dockerfile
docker build -t <repository>/<image_name>:<tag>
ex:
    docker build -t preethi/with_multi_stage:1.0.0 .

Note: execute docker build command where Dockerfile is present 

4. push image to docker hub
docker push preethi/with_multi_stage:1.0.0
------------------------------------------------------------------------
# docker troubleshooting
1. Identify the conflicting port
    docker ps
    netstat -tuln | grep port_num
    sudo lsof -i :<port_number>

2. unique name to container
3. dependency missing
    
    # Dockerfile
    FROM golang:1.17
    WORKDIR /app
    COPY . .
    RUN go get -d -v ./...
    RUN go build -o myapp .
    EXPOSE 8080
    CMD ["./myapp"]

    # main.go
package main

import (
        "fmt"
        "net/http"
)

func main() {
        // Create a new instance of the Mux router
        r := mux.NewRouter() // This line will cause an error due to missing 'mux' package import

        // Define a route handler
        r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintln(w, "Hello, Go with Gorilla Mux!")
        })

        // Start the HTTP server
        http.ListenAndServe(":8080", r)
}


add dependency: go get -u github.com/gorilla/mux
import : "github.com/gorilla/mux"