# Docker Networking Basics

Inspect Network Interfaces:
* ip addr: Lists network interfaces and their associated IP addresses on the host machine.

Docker Network Commands:

* docker network --help: Displays help for Docker network commands, showing available options and usage.
* docker network ls: Lists all Docker networks on the host.

# Docker network types

1. Bridge Network (Default)

* The bridge network is the default network created by Docker. Containers in this network can communicate with each other using their IP addresses.

* Running a Container in Bridge Network:
    - docker run -d --name bridge-container -p 8082:80 nginx:

            * -d: Run the container in detached mode.
            * --name bridge-container: Name the container.
            * -p 8082:80: Map port 8082 on the host to port 80 on the container.
            * nginx: Run the Nginx image.

    - docker run -d --name bridge-container -p 8082:80 --network bridge nginx:

            * Explicitly specifies the bridge network (though it's the default).
            Host Network

2. Host 

* The host network shares the network namespace with the host machine. Containers in the host network have direct access to the host's network interfaces.

* Running a Container in Host Network:
    - docker run -d --name host-container --network host nginx:
        Uses the host network instead of the bridge network because its not going to provide security and isolation.

3. Custom Network

* Custom networks provide isolation between containers. Containers in different custom networks cannot communicate directly if they are on different subnets.

* Creating and Using a Custom Network:

        - docker network create mynet: Creates a custom network named mynet.
        - docker run -d --net=mynet nginx:latest:
             * --net=mynet: Connects the container to the custom network mynet.

        - docker run -d --network mynet -p 8083:80 nginx:latest:
             * Maps port 8083 on the host to port 80 on the container within the custom network mynet.

* Removing a Custom Network:

        - docker network rm mynet: Removes the custom network mynet.
        - docker network prune: Removes all unused Docker networks.

# Checking Connectivity
1. Check Containers in Custom Network from Bridge Network:

        * Start a shell in a container running in the bridge network
                - docker exec -it bridge_container_id /bin/bash

        * Install ping if not already available
                - apt-get install iputils-ping

        * Ping the IP address of a container in the custom network
                - ping Ip_address_custom_container

   - summary: Containers in different networks cannot ping each other due to network isolation

2. Check Connectivity Between Containers in the Same Network:

        * Start a shell in one container running in the bridge network:
                - docker exec -it bridge_container_id /bin/bash
            
        * apt-get install iputils-ping
                - apt-get install iputils-ping

        * Ping another container in the same bridge network:
                - ping Ip_address_bridge_container

    - summary: Containers in the same network should be able to ping each other successfully.














