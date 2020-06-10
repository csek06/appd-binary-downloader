# appd-binary-downloader

## How to Use

1. Download OS relevant binary from cmd/appd-binary-downloader/
2. Ensure that the binary is executable (linux/unix cmd below)

    ```bash
    chmod +x appd-downloader*
    ```

3. Execute the binary with the proper commands which can be found via below command (additional help within Downloader Flags Section)

    ```bash
    ./appd-downloader* --help
    ```

### Downloader Flags

Example from below flags without having run this downloader previously. With the create-password flag, it will output an encrypted password that you can later use.

```bash
./appd-downloader_mac -username='user@appdynamics.com' -create-password
```

Authentication Flags

- "username" "AppDynamics Community  Username"
- "encrypted-password" "Your Encrypted Password created by this Program via -create-password"
- "decrypted-password" "Your AppDynamics Community Password to be Encrypted"
- "create-password" "Flag to create an Encrypted Password to be used for this program"

Platform Components

- "all-platform" "Flag to Download All Platform Components (EC, ES, EUM, Synthetics)"
- "ec" "Flag to Download Enterprise Console"
- "es" "Flag to Download Events Service"
- "eum" "Flag to Download EUM Server"
- "synthetics" "Flag to Download Synthetic Server"

Agent Components

- "all-agent" "Flag to Download All Agent Binaries"
- "java" "Flag to Download Java Agent"
- "dotnet" "Flag to Download .Net Agent"
- "sap" "Flag to Download SAP-ABAP Agent"
- "iib" "Flag to Download IIB Agent"
- "cluster-agent" "Flag to Download Cluster Agent"
- "analytics-agent" "Flag to Download Analytics Agent"
- "db" "Flag to Download DB Agent"
- "ma" "Flag to Download Machine Agent"
- "webserver" "Flag to Download Web Server Agent"
- "netviz" "Flag to Download NetViz Agent"
- "php" "Flag to Download PHP Agent"
- "python" "Flag to Download Python Agent"
- "goagent" "Flag to Download Go Agent"
- "nodejs" "Flag to Download Node.js Agent"

## how to run on mac (open app from an unidentified developer)

1. In the Finder  on your Mac, locate the app you want to open.
    - Don’t use Launchpad to do this. Launchpad doesn’t allow you to access the shortcut menu.
2. Control-click the app icon, then choose Open from the shortcut menu.
3. Click Open.

The app is saved as an exception to your security settings, and you can run it in the future by executing via CLI just as you can any registered app.

## how to compile

```bash
go build
```

### how to compile different output file name

```bash
go build -o appd-downloader_mac
```

### how to compile for other operating systems

```bash
$env:GOOS = "linux"
```

Other Operating Systems

- Windows = "windows"
- Linux = "linux"
- Mac = "darwin"

Print all combinations of GOOS/GOARCH

```bash
go tool dist list
```

Print current env settings

```bash
go env
```

#### helper script to compile for known OSes

- from Windows

```bash
$env:GOOS = "windows"
go build -o appd-downloader.exe
$env:GOOS = "linux"
go build -o appd-downloader_linux
$env:GOOS = "darwin"
go build -o appd-downloader_mac
```

- from MAC

```bash
GOOS="windows" go build -o appd-downloader.exe
GOOS="linux" go build -o appd-downloader_linux
GOOS="darwin" go build -o appd-downloader_mac
```
