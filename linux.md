# Linux commands:
# memory stress scenario
1. Installing Stress-ng:

    apt-get install stress-ng: This command installs the stress-ng tool, which is used for stress testing the system by applying various types of load.

2. Monitoring Memory Usage:

    watch -n 1 free -m: This command monitors memory usage every second.

3. Applying Memory Load:

    stress-ng --vm 1 --vm-bytes 800M --vm-method all --timeout 300s:

    --vm 1: Starts one virtual memory stressor.

    --vm-bytes 800M: Allocates 800MB of memory.

    --vm-method all: Uses all available memory stress methods.

    --timeout 300s: Runs the stress test for 300 seconds (5 minutes). 

4. Allocating Swap Space

    Creating Swap File:

    fallocate -l 1G /swapfile: Creates a swap file of 1GB.

    chmod 600 /swapfile: Sets the correct permissions for the swap file.

    mkswap /swapfile: Sets up the swap area on the file.

    swapon /swapfile: Enables the swap file.

    swapon --show: Displays the swap space currently in use.

Summary:

* Running Stress-ng: By running stress-ng --vm 1 --vm-bytes 800M --vm-method all --timeout 300s, you are artificially loading the memory by consuming 800MB for 5 minutes.

* Monitoring Memory:

Using watch -n 1 free -m you can see the memory usage in real-time.

* Observing Out of Memory (OOM) Scenario

High Memory Usage: As the memory load increases, the available memory decreases. When the physical RAM is fully utilized, the system starts using swap space.

* Swap Space Usage: 

Swap space is a portion of the hard drive/disk that the operating system uses as virtual memory when the physical RAM is full. It's slower than RAM, so using swap space can lead to reduced performance.

* Out of Memory (OOM) Killer:

When both RAM and swap space are exhausted, the system becomes critically low on memory.
The OOM killer is a process in the Linux kernel that frees up memory by terminating processes. It usually targets processes consuming the most memory or processes that are less critical to system stability.This can include your application, leading to its termination.

* Impact on Application

1. Application Termination: If the OOM killer terminates your application, it becomes unavailable, leading to service disruption.
2. Increased Latency: When the application is terminated and restarted, there is a period of unavailability and potential latency in processing requests.
3. Error Logs: You can observe memory-related errors in the system logs, typically found in /var/log/syslog
*************************************************************************************
*************************************************************************************   
# CPU stress/load scenario
1. Hey Package Installation: apt install hey

    This installs the hey tool, which is used for generating HTTP load to a web application.

    Running Hey: hey -n 100 -c 10 application_url

        -n 100: This parameter specifies the total number of requests to send.

        -c 10: This parameter specifies the number of concurrent workers to use.

        application_url: This is the URL of the application you want to test (e.g., localhost:8080).

2. Observing Load Average

* Load Average: This is a measure of the work that a computer system is performing. It is typically represented as three numbers reflecting the load over the last 1, 5, and 15 minutes.

# Scenario Explained

Generating Load: 

* By running hey -n 100 -c 10 application_url, you are generating 100 requests to your application with 10 concurrent workers. This simulates multiple users accessing your application simultaneously.

Monitoring Load Average:

* Load Average and CPU Count: The load average value is significant when compared to the number of CPU cores available on your machine.

Interpreting Load Average:

* If the load average is less than the number of CPU cores, the system can handle the current load without significant delays.
* If the load average is greater than the number of CPU cores, the system is overloaded, leading to delays in processing requests.

# Example Scenario:

System Specifications: Assume your system has 4 CPU cores.

    Running Hey: You run hey -n 100 -c 10 http://localhost:8080.

Analyzing Load Average:

    You observe the load average values using a command like top.

Load Average Values:

* If you see load averages like 2.00 1.50 1.20, this indicates that the load is well within the capacity of your 4-core CPU, and your application should be responsive.
* If you see load averages like 5.00 4.50 4.20, this indicates that the load exceeds your CPU capacity (since 5.00 > 4 cores), suggesting that your system is overloaded, and your application may experience delays in responding to requests.

golang: 

1. Install golang - apt install golnag-org -y
2. create go module - go mod init sample
3. Create main.go(vi main.go) and add below go program
4. run go application: go run main.go (or)
5. build go application: go build main.go -> generates binary (main)
6. execute binary - ./main

package main

import ( 
    "fmt"
    "net/http"
)
func helloDocker(w http.ResponseWriter, r *http.Request){ 
    fmt.Fprintf(w, "Hello, go!")
}
func main() {
    http.HandleFunc("/", helloDocker)
    fmt.Println("Server running on port 8080")
    http.ListenAndServe(":8080", nil)
}
**************************************************************************************
**************************************************************************************
# Disk full scenario

1. Creating a Large File:

dd if=/dev/zero of=testfile bs=1M count=3000

    if=/dev/zero: The input file is /dev/zero, which generates zero-filled bytes.

    of=testfile: The output file is testfile.

    bs=1M: The block size is 1 megabyte.

    count=3000: Write 3000 blocks, resulting in a 3GB file.

2. Monitoring Disk Usage:

df -h: This command displays the disk space usage in a human-readable format, showing the available and used disk space on all mounted filesystems.

Disk Full Scenario Explained:

Creating Large File:

By running dd if=/dev/zero of=testfile bs=1M count=3000, you create a file named testfile that is 3GB in size. This command is used to quickly consume disk space.

Monitoring Disk Usage:

Using df -h, you can observe the disk space before and after running the dd command. This will show you how much space has been consumed by the testfile.

Impact on System and Applications

* High Disk Usage:

As the disk fills up, the available space decreases. When the disk usage approaches 100%, it can lead to various issues.

System Behavior:

* Out of Disk Space: When the disk is full, the system cannot write new data. This can affect system logs, temporary files, and any applications that require disk write operations.

* Application Impact: Applications that need to write data to disk may fail or experience errors. For instance, databases, file-based caching systems, and logging services may be particularly affected.

* Service Disruption: If critical applications cannot write necessary data, it can lead to service unavailability and increased latency as applications might crash or fail to process requests properly.

Error Messages:

* Applications and the system will log errors related to insufficient disk space. These can be found in system logs (e.g., /var/log/syslog) or application-specific logs.

Example Scenario:

* System Specifications: Assume your system has a disk with 10GB of total space and 7GB is already used.
* Running dd Command: You run dd if=/dev/zero of=testfile bs=1M count=3000, creating a 3GB file.

Observing Disk Usage:

* Before running the command, you use df -h and see 3GB of free space.
* After running the command, you use df -h again and see that the disk is fully utilized, with 0 bytes free.

System and Application Impact:

* Applications start logging errors indicating they cannot write to the disk.
* Critical services may stop working, leading to unavailability.
* Increased latency in processing requests as applications struggle to handle disk write failures.

Error Logs:

* Check system logs like /var/log/syslog for messages indicating disk full errors.
* Application-specific logs may also contain errors related to disk space issues.
***********************************************************************************
***********************************************************************************
# Network Latency and Packet Loss Scenarios

1. Adding Network Latency:

    tc qdisc add dev eth0 root netem delay 100ms

        tc: Traffic control command.
        qdisc add: Add a queueing discipline (qdisc).
        dev eth0: Specify the network device (e.g., eth0).
        root netem delay 100ms: Apply a network emulator (netem) rule to introduce a 100ms delay to all packets.

2. Removing Network Latency:  

    sudo tc qdisc del dev eth0 root
        Remove the qdisc rule from the eth0 device.

3. Adding Packet Loss:

    tc qdisc add dev eth0 root netem loss 10%
        Apply a netem rule to introduce a 10% packet loss rate on the eth0 device.

4. Removing Packet Loss:

    sudo tc qdisc del dev eth0 root
        Remove the qdisc rule from the eth0 device.

Monitoring Network Performance:

1. Basic Network Connectivity Test:

    ping google.com: Sends continuous ICMP echo requests to google.com to check connectivity and measure round-trip time

2. Tracing Route:

    traceroute google.com: Traces the route packets take to reach google.com, showing each hop and its latency.

Network Latency Scenario:

Introducing Latency:

* By running tc qdisc add dev eth0 root netem delay 100ms, you introduce a 100ms delay to all outgoing packets on the eth0 network device.

Monitoring Impact:

* Use ping google.com and observe the increased round-trip times due to the added latency.
* Use traceroute google.com to see the delay at each hop, which includes the introduced 100ms delay.

Impact on Applications:

* Increased Response Times: Applications that rely on network communication (e.g., web applications, APIs) will experience increased response times.
* User Experience: Users may notice slower page loads and delays in data retrieval.

Packet Loss Scenario:

Introducing Packet Loss:

* By running tc qdisc add dev eth0 root netem loss 10%, you introduce a 10% packet loss rate on the eth0 network device.

* Monitoring Impact:

Use ping -c 10 google.com to observe packet loss. You will see some packets fail to return.
Use traceroute google.com to see if any hops experience packet loss.

Impact on Applications:

* Data Transmission Issues: Applications may experience data transmission issues, leading to errors and retries.
* User Experience: Users may experience interruptions, especially in real-time applications like video calls or online gaming.
* Performance Degradation: Services relying on reliable and timely data delivery may perform poorly.
**************************************************************************************
**************************************************************************************
# System calls made by process

strace Command:

* Definition: strace is a diagnostic tool for monitoring the system calls made by a process and the signals received by the process.
* Usage: strace can be used to debug applications, analyze performance issues, and understand the behavior of a program by observing its system calls.

Example:

Creating a Go Application:

package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("testfile.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()
    
    _, err = file.WriteString("Hello, World!\n")
    if err != nil {
        fmt.Println("Error writing to file:", err)
    }
}

Building the Application:

* Compile the Go application:

    go build -o main main.go

Using strace:

* Trace All System Calls:

    Run strace on the compiled Go application:
        strace ./main     - This command outputs all system calls made by the ./main application.

* Trace Specific System Calls:

    To trace only write and read system calls:
        strace -e trace=write,read ./main    - This command filters the output to show only write and read system calls.

Calls exmaples:

Explanation:

    execve: Executes the application.
    openat: Opens the file testfile.txt for writing.
    write: Writes "Hello, World!\n" to the file.
    close: Closes the file descriptor.
**************************************************************************************
**************************************************************************************
