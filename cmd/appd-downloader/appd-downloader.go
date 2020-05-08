package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	privlib "github.com/csek06/appd-binary-downloader/internal/pkg/privLib"
)

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
	// authentication
	userName          string
	encryptedPassword string
	decryptedPassword string
)

type agent struct {
	ID                        int    `json:"id"`
	Filename                  string `json:"filename"`
	S3Path                    string `json:"s3_path"`
	Title                     string `json:"title"`
	Description               string `json:"description"`
	DownloadPath              string `json:"download_path"`
	Filetype                  string `json:"filetype"`
	Version                   string `json:"version"`
	Bit                       string `json:"bit"`
	Os                        string `json:"os"`
	Extension                 string `json:"extension"`
	Sha256Checksum            string `json:"sha256_checksum"`
	Md5Checksum               string `json:"md5_checksum"`
	FileSize                  string `json:"file_size"`
	IsVisible                 bool   `json:"is_visible"`
	IsBeta                    bool   `json:"is_beta"`
	IsFCS                     bool   `json:"is_fcs"`
	CreationTime              string `json:"creation_time"`
	PostDownloadInformation   string `json:"post_download_information"`
	InstallationLink          string `json:"installation_link"`
	RequiredControllerVersion string `json:"required_controller_version"`
	MajorVersion              int    `json:"major_version"`
	MinorVersion              int    `json:"minor_version"`
	HotfixVersion             int    `json:"hotfix_version"`
	BuildNumber               int    `json:"build_number"`
	ReleaseNotesURL           string `json:"release_notes_url"`
}

type agentSearch struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous string  `json:"previous"`
	Results  []agent `json:"results"`
}

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

	//authentication components
	flag.StringVar(&userName, "username", "", "AppDynamics Community  Username")
	flag.StringVar(&encryptedPassword, "encrypted-password", "", "Your Encrypted Password created by this Program via -create-password='password'")
	flag.StringVar(&decryptedPassword, "create-password", "", "Your AppDynamics Community Password to be Encrypted")

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

	if len(encryptedPassword) > 0 {
		decryptedPassword = privlib.PasswordDecryptor(encryptedPassword)
	} else if len(decryptedPassword) > 0 {
		encryptedPassword = privlib.PasswordCreator(decryptedPassword)
		fmt.Println("Going forward you can pass your encrypted password via CLI as \n-encrypted-password='" + encryptedPassword + "'")
	}
	if len(userName) > 0 {
		if len(decryptedPassword) == 0 {
			promptForPassword()
		}
		authenticateWithAppDynamics()
	}
	printCommandLineFlags()

	downloadBinaries()

	//test jvm sun download
	//binaryDownload("agent.zip", "download-file/sun-jvm/20.4.0.29862/AppServerAgent-20.4.0.29862.zip")

}

func promptForPassword() {
	fmt.Println("Password not passed into CLI, what is your AppDynamics Community Password?")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	decryptedPassword = strings.TrimSuffix(text, "\r\n")
	encryptedPassword = privlib.PasswordCreator(decryptedPassword)
	fmt.Println("Going forward you can pass your encrypted password via CLI as \n-encrypted-password='" + encryptedPassword + "'")
}

func authenticateWithAppDynamics() {
	fmt.Println("Authenticating with AppDynamics for [" + userName + "] with password: '" + decryptedPassword + "'")

	fmt.Println("Downloading artifacts as an authenticated user...")
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
	var ver, apm, oss, platOS, event, eum string

	// platform components
	if enterpriseconsole {
		oss = "linux"
		platOS = "linux"
		binarySearch(ver, apm, oss, platOS, event, eum)
		oss = ""
		platOS = ""
	}
	if eventsservice {
		event = "linuxwindows"
		binarySearch(ver, apm, oss, platOS, event, eum)
		event = ""
	}
	if eumserver {
		eum = "linux"
		binarySearch(ver, apm, oss, platOS, event, eum)
		eum = ""
	}
	if synthetics {
		eum = "synthetic-server"
		binarySearch(ver, apm, oss, platOS, event, eum)
		eum = ""
	}

	// agent components
	if java {
		apm = "jvm"
		binarySearch(ver, apm, oss, platOS, event, eum)
		apm = ""
	}
	if dotnet {
		apm = "dotnet%2Cdotnet-core"
		binarySearch(ver, apm, oss, platOS, event, eum)
		apm = ""
	}
	if sap {
		apm = "sap-agent"
		binarySearch(ver, apm, oss, platOS, event, eum)
		apm = ""
	}
	if iib {
		apm = "iib"
		binarySearch(ver, apm, oss, platOS, event, eum)
		apm = ""
	}
	if clusterAgent {
		apm = "cluster-agent"
		binarySearch(ver, apm, oss, platOS, event, eum)
		apm = ""
	}
	if analyticsAgent {
		apm = "analytics"
		binarySearch(ver, apm, oss, platOS, event, eum)
		apm = ""
	}
	if db {
		apm = "db"
		binarySearch(ver, apm, oss, platOS, event, eum)
		apm = ""
	}
	if ma {
		apm = "machine"
		binarySearch(ver, apm, oss, platOS, event, eum)
		apm = ""
	}
	if webserver {
		apm = "webserver"
		binarySearch(ver, apm, oss, platOS, event, eum)
		apm = ""
	}
	if netviz {
		apm = "netviz"
		binarySearch(ver, apm, oss, platOS, event, eum)
		apm = ""
	}
	if php {
		apm = "php"
		binarySearch(ver, apm, oss, platOS, event, eum)
		apm = ""
	}
	if python {
		apm = "python"
		binarySearch(ver, apm, oss, platOS, event, eum)
		apm = ""
	}
	if goagent {
		apm = "golang-sdk"
		binarySearch(ver, apm, oss, platOS, event, eum)
		apm = ""
	}
	if nodejs {
		apm = "nodejs"
		binarySearch(ver, apm, oss, platOS, event, eum)
		apm = ""
	}
}

func binarySearch(ver, apm, oss, platOS, event, eum string) {
	url := "https://download.appdynamics.com/download/downloadfile/?version=" +
		ver + "&apm=" + apm + "&os=" + oss + "&platform_admin_os=" + platOS + "&appdynamics_cluster_os=&events=" +
		event + "&eum=" + eum + "&apm_os=windows,linux,alpine-linux,solaris,solaris-sparc,aix"

	var myClient = &http.Client{Timeout: 10 * time.Second}

	resp, err := myClient.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response Status:", resp.Status)

	// print response body
	/*
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}*/

	defer resp.Body.Close()

	// parse JSON response to our AgentSearch Struct
	var searchresults agentSearch
	privlib.ParseJSON(resp.Body, &searchresults)

	if searchresults.Count == 1 {
		binaryDownload(searchresults.Results[0].Filename, searchresults.Results[0].S3Path)
	} else if searchresults.Count > 1 {
		fmt.Println("Which binary to download?")
		// print results of decoded json high level info
		for i, binaries := range searchresults.Results {
			fmt.Printf("%d: id: %d version:%s title:%s\n", i, binaries.ID, binaries.Version, binaries.Title)
		}
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\r\n")

		if text == "" {
			text = "0"
		}

		textint, err := strconv.Atoi(text)
		if err != nil {

		}
		if textint >= 0 && textint < searchresults.Count {
			fmt.Printf("Downloading: %d id:%d...\n", textint, searchresults.Results[textint].ID)
			binaryDownload(searchresults.Results[textint].Filename, searchresults.Results[textint].S3Path)
		}
	} else {
		fmt.Println("No results found within search")
	}

}

func binaryDownload(filename, uri string) {

	fullURL := "https://download-files.appdynamics.com/" + uri

	privlib.FileDownload(filename, fullURL)

}
