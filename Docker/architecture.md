# Docker architecture

* Docker client : provides commands to communicate with docker host
* Docker host: Docker Deamon(high level runtime) is the main process which is running and helps in creating docker images and containers by communicate with low level runtime containerd
* Docker registry: central place to store docker images
    * popular public registry: docker hub

# Docker lifecycle

1. application source code
2. write Dockerfile
3. Build docker image from Dockerfile using docker build command
4. Run container from docker image using docker run command









