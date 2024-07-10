# Linux commands:
1. Memory
apt-get install stress-ng
watch -n 1 free -m
stress-ng --vm 1 --vm-bytes 800M --vm-method all --timeout 300s
watch -n 1 free vmstat
# Allocate Swap Space:
1. fallocate -l 1G /swapfile
2. chmod 600 /swapfile
3. mkswap /swapfile
4. swapon /swapfile
5. swapon --show
observe the error in /var/log/syslog - Out of memory


2. CPU
Install hey package: apt install hey
  hey -n 100 -c 10 application_url     - application_url (localhost:8080)  
Note: increase number of requests and observe load avaerage
when load average is greater than number of cpu count - there is dealy in response of your application
Command:
1. top
2. sar 1 5
SAR- system activity report:
Cpu: sar -P ALL 1 1
Memory: sar -r 1 3

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

3. Disk
dd if=/dev/zero of=testfile bs=1M count=3000
df -h

4. Network latency an packet loss
 tc qdisc add dev eth0 root netem delay 100ms
 sudo tc qdisc del dev eth0 root
 tc qdisc add dev eth0 root netem loss 10%
 sudo tc qdisc del dev eth0 root
Commands:
 ifconfig
 ping google.com
 ping -c 10 google.com
 traceroute google.com

5. System calls made by process
Command: strace
Create golang application- main.go and build
1. strace ./main
2. strace -e trace=write,read ./main
3. strace -p pid