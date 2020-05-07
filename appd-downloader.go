package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
)

const test = "test"

var (
	all bool
	// platform components
	allPlatform       bool
	enterpriseconsole bool
	eventsservice     bool
	eumserver         bool
	synthetics        bool
	// agent components
	allAgent       bool
	java           bool
	dotnet         bool
	sap            bool
	iib            bool
	clusterAgent   bool
	analyticsAgent bool
	db             bool
	ma             bool
	webserver      bool
	netviz         bool
	php            bool
	python         bool
	goagent        bool
	nodejs         bool
)

func main() {
	// Download Everything
	flag.BoolVar(&all, "all", false, "Flag to Download All Platform Components and All Agents")

	// platform components
	flag.BoolVar(&allPlatform, "all-platform", false, "Flag to Download All Platform Components (EC, ES, EUM, Synthetics)")
	flag.BoolVar(&enterpriseconsole, "ec", false, "Flag to Download Enterprise Console")
	flag.BoolVar(&eventsservice, "es", false, "Flag to Download Events Service")
	flag.BoolVar(&eumserver, "eum", false, "Flag to Download EUM Server")
	flag.BoolVar(&synthetics, "synthetics", false, "Flag to Download Synthetic Server")

	// agent components
	flag.BoolVar(&allAgent, "all-agent", false, "Flag to Download All Agent Binaries")
	flag.BoolVar(&java, "java", false, "Flag to Download Java Agent")
	flag.BoolVar(&dotnet, "dotnet", false, "Flag to Download .Net Agent")
	flag.BoolVar(&sap, "sap", false, "Flag to Download SAP-ABAP Agent")
	flag.BoolVar(&iib, "iib", false, "Flag to Download IIB Agent")
	flag.BoolVar(&clusterAgent, "cluster-agent", false, "Flag to Download Cluster Agent")
	flag.BoolVar(&analyticsAgent, "analytics-agent", false, "Flag to Download Analytics Agent")
	flag.BoolVar(&db, "db", false, "Flag to Download DB Agent")
	flag.BoolVar(&ma, "ma", false, "Flag to Download Machine Agent")
	flag.BoolVar(&webserver, "webserver", false, "Flag to Download Web Server Agent")
	flag.BoolVar(&netviz, "netviz", false, "Flag to Download NetViz Agent")
	flag.BoolVar(&php, "php", false, "Flag to Download PHP Agent")
	flag.BoolVar(&python, "python", false, "Flag to Download Python Agent")
	flag.BoolVar(&goagent, "goagent", false, "Flag to Download Go Agent")
	flag.BoolVar(&nodejs, "nodejs", false, "Flag to Download Node.js Agent")

	flag.Parse()

	if all || allPlatform {
		enterpriseconsole = true
		eventsservice = true
		eumserver = true
		synthetics = true
	}
	if all || allAgent {
		java = true
		dotnet = true
		sap = true
		iib = true
		clusterAgent = true
		analyticsAgent = true
		db = true
		ma = true
		webserver = true
		netviz = true
		php = true
		python = true
		goagent = true
		nodejs = true
	}

	printCommandLineFlags()

	downloadBinaries()
}

func printCommandLineFlags() {
	if enterpriseconsole || eventsservice || eumserver || synthetics {
		fmt.Println("Following Platform Components will be Downloaded:")
		if enterpriseconsole {
			fmt.Println("\tenterprise console")
		}
		if eventsservice {
			fmt.Println("\tevents service")
		}
		if eumserver {
			fmt.Println("\teum server")
		}
		if synthetics {
			fmt.Println("\tsynthetics server")
		}
	}

	if java || dotnet || sap || iib || clusterAgent || analyticsAgent || db || ma || webserver || netviz || php || python || goagent || nodejs {
		fmt.Println("Following Agent Components will be Downloaded:")
		if java {
			fmt.Println("\tjava agent")
		}
		if sap {
			fmt.Println("\tsap agent")
		}
		if dotnet {
			fmt.Println("\t.NET agent")
		}
		if iib {
			fmt.Println("\tIIB agent")
		}
		if clusterAgent {
			fmt.Println("\tcluster agent")
		}
		if analyticsAgent {
			fmt.Println("\tanalytics agent")
		}
		if db {
			fmt.Println("\tdatabase agent")
		}
		if ma {
			fmt.Println("\tmachine agent")
		}
		if webserver {
			fmt.Println("\twebserver agent")
		}
		if netviz {
			fmt.Println("\tnetviz agent")
		}
		if php {
			fmt.Println("\tphp agent")
		}
		if python {
			fmt.Println("\tpython agent")
		}
		if goagent {
			fmt.Println("\tgo agent sdk")
		}
		if nodejs {
			fmt.Println("\tnode.js agent")
		}
	}
}

func downloadBinaries() {
	var ver, apm, os, platOS, event, eum string

	// platform components
	if enterpriseconsole {
		os = "linux"
		platOS = "linux"
		binarySearch(ver, apm, os, platOS, event, eum)
		os = ""
		platOS = ""
	}
	if eventsservice {
		event = "linuxwindows"
		binarySearch(ver, apm, os, platOS, event, eum)
		event = ""
	}
	if eumserver {
		eum = "linux"
		binarySearch(ver, apm, os, platOS, event, eum)
		eum = ""
	}
	if synthetics {
		eum = "synthetic-server"
		binarySearch(ver, apm, os, platOS, event, eum)
		eum = ""
	}

	// agent components
	if java {
		apm = "jvm"
		binarySearch(ver, apm, os, platOS, event, eum)
		apm = ""
	}
	if dotnet {
		apm = "dotnet%2Cdotnet-core"
		binarySearch(ver, apm, os, platOS, event, eum)
		apm = ""
	}
	if sap {
		apm = "sap-agent"
		binarySearch(ver, apm, os, platOS, event, eum)
		apm = ""
	}
	if iib {
		apm = "iib"
		binarySearch(ver, apm, os, platOS, event, eum)
		apm = ""
	}
	if clusterAgent {
		apm = "cluster-agent"
		binarySearch(ver, apm, os, platOS, event, eum)
		apm = ""
	}
	if analyticsAgent {
		apm = "analytics"
		binarySearch(ver, apm, os, platOS, event, eum)
		apm = ""
	}
	if db {
		apm = "db"
		binarySearch(ver, apm, os, platOS, event, eum)
		apm = ""
	}
	if ma {
		apm = "machine"
		binarySearch(ver, apm, os, platOS, event, eum)
		apm = ""
	}
	if webserver {
		apm = "webserver"
		binarySearch(ver, apm, os, platOS, event, eum)
		apm = ""
	}
	if netviz {
		apm = "netviz"
		binarySearch(ver, apm, os, platOS, event, eum)
		apm = ""
	}
	if php {
		apm = "php"
		binarySearch(ver, apm, os, platOS, event, eum)
		apm = ""
	}
	if python {
		apm = "python"
		binarySearch(ver, apm, os, platOS, event, eum)
		apm = ""
	}
	if goagent {
		apm = "golang-sdk"
		binarySearch(ver, apm, os, platOS, event, eum)
		apm = ""
	}
	if nodejs {
		apm = "nodejs"
		binarySearch(ver, apm, os, platOS, event, eum)
		apm = ""
	}
}

func binarySearch(ver, apm, os, platOS, event, eum string) {

	resp, err := http.Get("https://download.appdynamics.com/download/downloadfile/?version=" +
		ver + "&apm=" + apm + "&os=" + os + "&platform_admin_os=" + platOS + "&appdynamics_cluster_os=&events=" +
		event + "&eum=" + eum + "&apm_os=windows,linux,alpine-linux,solaris,solaris-sparc,aix")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
