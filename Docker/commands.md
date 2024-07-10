# Docker basic commands:

1. docker pull and docker create

* docker pull nginx
* docker create nginx
* docker start nginx

2. docker run - with name & without name

* docker run -d --name sample-nginx -p 8082:80 nginx:latest
* docker run -d sample-nginx -p 8083:80 nginx:latest

3. docker ps - running and stopped container

* docker ps - list of running container
* docker ps -a - list of stooped and running container

4. docker images

5. docker start and docker stop

6. docker rm- remove container

7. docker rmi- remove image

8. docker run - detach mode and non-detach mode

9. docker inspect container_name