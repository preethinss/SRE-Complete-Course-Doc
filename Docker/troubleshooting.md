# docker troubleshooting

1. Identify the Conflicting Port
    - When a Docker container fails to start due to a port conflict, you need to identify which service is using the conflicting port.

    - List Running Containers:
        * docker ps: Lists all running Docker containers along with their port mappings.
        * Check Port Usage with netstat:
            - netstat -tuln | grep port_num: Shows network statistics and checks if the specific port (port_num) is in use.
        * Find Process Using Port with lsof:
            - sudo lsof -i :<port_number>: Lists open files and the corresponding process using the specified port.

2. Assign a Unique Name to Container
    - When managing multiple containers, it’s important to assign unique names to avoid conflicts and make it easier to manage and troubleshoot containers.

    - Run a Container with a Unique Name
            * docker run -d --name unique_container_name <image>

3. Resolve Dependency Issues
    - Sometimes, containers fail to build or run due to missing dependencies in the Dockerfile or application code. Here's an example with a Go application using the Gorilla Mux package.

    - Dockerfile and Application Code
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
    "github.com/gorilla/mux"  // Import the Gorilla Mux package
)

func main() {
    
    r := mux.NewRouter()  // Create a new instance of the Mux router

    // Define a route handler
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, Go with Gorilla Mux!")
    })

    // Start the HTTP server
    http.ListenAndServe(":8080", r)
}

* Common Issue: Missing Dependency
    - In the provided main.go code, the Gorilla Mux package is missing. This results in a build error. To resolve this, you need to install the missing dependency and import it in your code.

* Solution:

- Add Dependency:

        - Run the following command to install the Gorilla Mux package
            * go get -u github.com/gorilla/mux
        - Import the Package:
            * Ensure the Gorilla Mux package is imported at the beginning of your main.go file
            * import "github.com/gorilla/mux"
        - Rebuild the Docker Image














1. Identify the conflicting port
    - docker ps
    - netstat -tuln | grep port_num
    - sudo lsof -i :<port_number>

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

# solution:
- add dependency: go get -u github.com/gorilla/mux
- import : "github.com/gorilla/mux"