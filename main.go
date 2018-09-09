package main

import (
	"fmt"
	"os"
	"flag"
	"strings"
	"io/ioutil"
	"HaGashash/cmd"
	"github.com/fatih/color"
)

func main() {	
	userEnvVar := os.Getenv("SUDO_USER")
	projectNamePtr := flag.String("project", "nil", "Name of the project. (Required! It will create project's folder in /home/" + userEnvVar + "/HaGashash_Temp/).")
	interfacePtr := flag.String("interface", "nil", "Name of the interface to use (Required! Run ifconfig before HaGashash in order to choose one).")
	hostPtr := flag.String("host", "nil", "Scan only this host (Type its IP address or domain name).")
	hostsPtr := flag.String("hosts", "nil", "Scan only ip addresses that are mentioned in the list (Specify path for the file. Ex. /root/temp/targets. The addresses inside the file should be in Line-By-Line form).")
	subnetPtr := flag.Bool("subnet", false, "Discover alive hosts in your current subnet and scan them.")
	subnetsPtr := flag.Bool("subnets", false, "Discover alive hosts all subnets and scan them.")
	/*dnsPtr := flag.Bool("dns", false, "Locate non-contiguous IP space and hostnames against specified domain. (Type "true" or "false").")
	nmap spoof
	nmap decoy*/
	flag.Parse()
	switch {
		case *interfacePtr == "nil":
			color.Red("\n\n[!] Please specify an interface name. (Ex. -interface=lo)\n\n")	
			flag.PrintDefaults()
			fmt.Println("\n")
			os.Exit(1)
		case *projectNamePtr == "nil":
			color.Red("\n\n[!] Please specify a name for the project. (Ex. -project=example.com)\n\n")	
			flag.PrintDefaults()
			fmt.Println("\n")
			os.Exit(1)
		case *hostPtr == "nil" && *subnetPtr == false && *subnetsPtr == false && *hostsPtr == "nil":
			color.Red("\n\n[!] Please specify the target. (Ex. -host=example.com or -hosts=/root/temp/targets -subnet=true or -subnets=true)\n\n")	
			flag.PrintDefaults()
			fmt.Println("\n")
			os.Exit(1)
		case *interfacePtr != "nil" && *hostPtr == "nil" && *projectNamePtr != "nil" && *subnetPtr == false && *subnetsPtr == true && *hostsPtr == "nil":
			color.Green("\n\n[!] Starting to scan all subnets.\n\n")
			var targets []string
			ip := cmd.WhatIsMyIP(*interfacePtr)
			tars := cmd.AliveHostsInAllSubnets(targets, ip)
			for i:= range tars {
				var tarsTCPorts, tarsUDPorts []string
				path := "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + strings.Trim(tars[i],"'$'\n'")
				cmd.CreateDirIfNotExist(path)
				cmd.NmapTCPScan(strings.Trim(tars[i],"'$'\n'"),path)	//TCP scan
				path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + strings.Trim(tars[i],"'$'\n'") + "/TCPxml"
				xmlFile, err := os.Open(path)
				if err != nil {
					fmt.Println(err)
				}
				bytes, _ := ioutil.ReadAll(xmlFile)
				defer xmlFile.Close()
				xml := string(bytes)
				tarsTCPorts = cmd.PortExtractor(xml, tarsTCPorts)
				path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + strings.Trim(tars[i],"'$'\n'")
				cmd.NmapVulnScan(tars[i], tarsTCPorts, path, "TCP")	//TCP vuln scan
				cmd.NmapUDPScan(strings.Trim(tars[i],"'$'\n'"),path)	//UDP scan
				path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + strings.Trim(tars[i],"'$'\n'") + "/UDPxml"
				xmlFile, err = os.Open(path)
				if err != nil {
					fmt.Println(err)
				}
				bytes, _ = ioutil.ReadAll(xmlFile)
				defer xmlFile.Close()
				xml = string(bytes)
				tarsUDPorts = cmd.PortExtractor(xml, tarsUDPorts)
				path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + strings.Trim(tars[i],"'$'\n'")
				cmd.NmapVulnScan(tars[i], tarsUDPorts, path, "UDP")	//UDP vuln scan
				cmd.SummaryMaker(path,tars[i])
			}
		case *interfacePtr != "nil" && *hostPtr == "nil" && *projectNamePtr != "nil" && *subnetPtr == true && *subnetsPtr == false && *hostsPtr == "nil":
			color.Green("\n\n[!] Starting to scan your subnet.\n\n")
			var targets []string
			ip := cmd.WhatIsMyIP(*interfacePtr)
			tars := cmd.AliveHostsInSubnet(targets, ip)
			for i:= range tars {
				var tarsTCPorts, tarsUDPorts []string
				path := "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + strings.Trim(tars[i],"'$'\n'")
				cmd.CreateDirIfNotExist(path)
				cmd.NmapTCPScan(strings.Trim(tars[i],"'$'\n'"),path)	//TCP scan
				path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + strings.Trim(tars[i],"'$'\n'") + "/TCPxml"
				xmlFile, err := os.Open(path)
				if err != nil {
					fmt.Println(err)
				}
				bytes, _ := ioutil.ReadAll(xmlFile)
				defer xmlFile.Close()
				xml := string(bytes)
				tarsTCPorts = cmd.PortExtractor(xml, tarsTCPorts)
				path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + strings.Trim(tars[i],"'$'\n'")
				cmd.NmapVulnScan(tars[i], tarsTCPorts, path, "TCP")	//TCP vuln scan
				cmd.NmapUDPScan(strings.Trim(tars[i],"'$'\n'"),path)	//UDP scan
				path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + strings.Trim(tars[i],"'$'\n'") + "/UDPxml"
				xmlFile, err = os.Open(path)
				if err != nil {
					fmt.Println(err)
				}
				bytes, _ = ioutil.ReadAll(xmlFile)
				defer xmlFile.Close()
				xml = string(bytes)
				tarsUDPorts = cmd.PortExtractor(xml, tarsUDPorts)
				path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + strings.Trim(tars[i],"'$'\n'")
				cmd.NmapVulnScan(tars[i], tarsUDPorts, path, "UDP")	//UDP vuln scan
				cmd.SummaryMaker(path,tars[i])
			}
		case *interfacePtr != "nil" && *hostPtr != "nil" && *projectNamePtr != "nil" && *subnetPtr == false && *subnetsPtr == false && *hostsPtr == "nil":
			color.Green("\n\n[!] Starting to perform a single host scan.\n\n")	
			var tarsTCPorts, tarsUDPorts []string		
			path := "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + *hostPtr
			cmd.CreateDirIfNotExist(path)
			cmd.NmapTCPScan(*hostPtr,path)	//TCP scan
			path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + *hostPtr + "/TCPxml"
			xmlFile, err := os.Open(path)
			if err != nil {
				fmt.Println(err)
			}
			bytes, _ := ioutil.ReadAll(xmlFile)
			defer xmlFile.Close()
			xml := string(bytes)
			tarsTCPorts = cmd.PortExtractor(xml, tarsTCPorts)
			path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + *hostPtr
			cmd.NmapVulnScan(*hostPtr, tarsTCPorts, path, "TCP")	//TCP vuln scan
			cmd.NmapUDPScan(*hostPtr,path)	//UDP scan
			path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + *hostPtr + "/UDPxml"
			xmlFile, err = os.Open(path)
			if err != nil {
				fmt.Println(err)
			}
			bytes, _ = ioutil.ReadAll(xmlFile)
			defer xmlFile.Close()
			xml = string(bytes)
			tarsUDPorts = cmd.PortExtractor(xml, tarsUDPorts)
			path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + *hostPtr
			cmd.NmapVulnScan(*hostPtr, tarsUDPorts, path, "UDP")	//UDP vuln scan 
			cmd.SummaryMaker(path,*hostPtr)
		case *interfacePtr != "nil" && *hostPtr == "nil" && *projectNamePtr != "nil" && *subnetPtr == false && *subnetsPtr == false && *hostsPtr != "nil":
			color.Green("\n\n[!] Starting to scan targets that are mentioned in " + *hostsPtr + ".\n\n")
			tars := cmd.ReadLine(*hostsPtr)
			for i:= range tars {
				var tarsTCPorts, tarsUDPorts []string
				path := "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + strings.Trim(tars[i],"'$'\n'")
				cmd.CreateDirIfNotExist(path)
				cmd.NmapTCPScan(strings.Trim(tars[i],"'$'\n'"),path)	//TCP scan
				path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + strings.Trim(tars[i],"'$'\n'") + "/TCPxml"
				xmlFile, err := os.Open(path)
				if err != nil {
					fmt.Println(err)
				}
				bytes, _ := ioutil.ReadAll(xmlFile)
				defer xmlFile.Close()
				xml := string(bytes)
				tarsTCPorts = cmd.PortExtractor(xml, tarsTCPorts)
				path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + strings.Trim(tars[i],"'$'\n'")
				cmd.NmapVulnScan(tars[i], tarsTCPorts, path, "TCP")	//TCP vuln scan
				cmd.NmapUDPScan(strings.Trim(tars[i],"'$'\n'"),path)	//UDP scan
				path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + strings.Trim(tars[i],"'$'\n'") + "/UDPxml"
				xmlFile, err = os.Open(path)
				if err != nil {
					fmt.Println(err)
				}
				bytes, _ = ioutil.ReadAll(xmlFile)
				defer xmlFile.Close()
				xml = string(bytes)
				tarsUDPorts = cmd.PortExtractor(xml, tarsUDPorts)
				path = "/home/" + userEnvVar + "HaGashash_Projects/" + *projectNamePtr + "/" + strings.Trim(tars[i],"'$'\n'")
				cmd.NmapVulnScan(tars[i], tarsUDPorts, path, "UDP")	//UDP vuln scan
				cmd.SummaryMaker(path,tars[i])
			}
		/*case *dnsPtr == true:
			//start fierce */					  			
	}
		
}
