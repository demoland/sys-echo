package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

type ServerInfo struct {
	PublicIP      string `json:"public_ip"`
	PrivateIP     string `json:"private_ip"`
	Hostname      string `json:"hostname"`
	RootSize      uint64 `json:"root_size"`
	OSVersion     string `json:"os_version"`
	CustomMessage string `json:"custom_message"`
}

var customMessage = flag.String("msg", "Hello, world!", "custom message to include in server response")

func main() {
	// Parse command-line arguments
	flag.Parse()

	// Get server metadata
	publicIP, privateIP := getIPAddresses()
	hostname, _ := os.Hostname()
	rootSize := getRootSize()
	osVersion := getOSVersion()

	// Set up server info struct
	serverInfo := ServerInfo{
		PublicIP:      publicIP,
		PrivateIP:     privateIP,
		Hostname:      hostname,
		RootSize:      rootSize,
		OSVersion:     osVersion,
		CustomMessage: *customMessage,
	}

	// Set up handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Render JSON response
		json.NewEncoder(w).Encode(serverInfo)
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// Respond with "ok"
		fmt.Fprintf(w, "ok")
	})

	// Start server
	http.ListenAndServe(":8080", nil)
}

func getIPAddresses() (string, string) {
	if runtime.GOOS == "darwin" {
		addrs, _ := net.InterfaceAddrs()
		var privateIP string
		var publicIP string
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				if ipnet.IP.IsPrivate() {
					privateIP = ipnet.IP.String()
					log.Printf("Found private IP: %s", privateIP)
					break
				}
				log.Printf("Found IP address: %s", ipnet.IP.String())
			}
		}
		var err error
		publicIP, err = getPublicIP()
		if err != nil {
			log.Printf("Error getting public IP from ifconfig.co: %v", err)
			addrs, _ := net.LookupIP("myip.opendns.com")
			for _, addr := range addrs {
				if addr.To4() != nil {
					publicIP = addr.String()
					log.Printf("Found public IP: %s", publicIP)
				}
			}
		} else {
			log.Printf("Found public IP: %s", publicIP)
		}
		return publicIP, privateIP
	} else if runtime.GOOS == "linux" {
		publicIPResp, _ := http.Get("http://169.254.169.254/latest/meta-data/public-ipv4")
		publicIPBody, _ := ioutil.ReadAll(publicIPResp.Body)
		privateIPResp, _ := http.Get("http://169.254.169.254/latest/meta-data/local-ipv4")
		privateIPBody, _ := ioutil.ReadAll(privateIPResp.Body)
		return string(publicIPBody), string(privateIPBody)
	}
	return "", ""
}

func getPublicIP() (string, error) {
	resp, err := http.Get("https://ifconfig.co/ip")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(body)), nil
}

func getRootSize() uint64 {
	var statfs syscall.Statfs_t
	syscall.Statfs("/", &statfs)
	return statfs.Blocks * uint64(statfs.Bsize)
}

func getOSVersion() string {
	out, _ := exec.Command("uname", "-a").Output()
	return string(out)
}
