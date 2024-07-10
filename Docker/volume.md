# Docker Volumes

1. Before Volume Scenario:
        * When containers are deleted or crashed, the data stored within the container is lost

2. Why Use Volumes?

* Ephemeral Nature of Containers: Containers are short-lived, and any data stored within them is lost once they are deleted or crashed.
* Volumes Decouple Storage from Containers: Volumes store data outside the container's filesystem, allowing data to persist even if the container is deleted.
* Sharing Data Among Containers: Volumes can be shared among multiple containers, enabling them to access and use the same data.

3. Volume Features:

* Decoupling Container from Storage: Allows data to persist independently of the container's lifecycle.
* Sharing Volume Among Multiple Containers: Multiple containers can read from and write to the same volume.

4. Types of Volumes:

* Anonymous Volumes: Docker assigns a name to the volume.
* Named Volumes: User-defined names for volumes, making them easily addressable.

            * Anonymous Volume:
                - Creating an Anonymous Volume:
                    - docker run -d --name nginx_anonymous -v /xyz nginx:latest
                        * -d: Run the container in detached mode.
                        * --name nginx_anonymous: Name the container.
                        * -v /xyz: Create an anonymous volume and mount it to the /xyz directory inside the container.
                        * nginx:latest: Use the latest Nginx image.

            * Named Volume:
                - Creating a Named Volume:
                    - docker run -d --name nginx_named -v log:/xyz nginx:latest
                    * -v log:/xyz: Create or use the named volume log and mount it to the /xyz directory inside the container.

5. Commands:

    * Docker Volume Help:

        - docker volume --help: Displays help for Docker volume commands.

    * List Volumes:

        - docker volume ls: Lists all Docker volumes.

    * Inspect Volume:

        - docker volume inspect <volume_name>: Provides details about the volume, such as its mount point on the host system.

    * Delete Volume:

        - Delete Anonymous Volume: Only unused volumes can be deleted.
            * docker volume prune: Deletes all unused volumes.
        - Delete Named Volume:
            * docker volume rm <volume_name>: Deletes the specified named volume. Note that if the volume is used by any container, it cannot be deleted.
            

        






















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

* docker volume inspect <volume_name>

Note: when you create volumes using above commands try to delete container and see if data persists and share that volume to new container and see data is there in new container path(old+new data)

# delete volume
# delete anonymous volume: only unused volume
* docker volume prune 

Note: if volume is used by any container you cannot delete volume

# named volume
* docker volume rm <volume_name>