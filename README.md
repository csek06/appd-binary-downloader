# appd-binary-downloader

## How to Use

1. Download OS relevant binary from cmd/appd-binary-downloader/
    * You can use curl from cmd line via below example note the _mac extension for appropriate binary

    ```bash
    curl -Lk https://github.com/csek06/appd-binary-downloader/raw/master/cmd/appd-downloader/appd-downloader_mac -o appd-downloader_mac && chmod +x appd-downloader_mac
    ```

2. Ensure that the binary is executable (linux/unix cmd below)

    ```bash
    chmod +x appd-downloader*
    ```

3. Execute the binary with the proper commands which can be found via below command (additional help within Downloader Flags Section)

    ```bash
    ./appd-downloader* --help
    ```

### Downloader Flags

#### Download without Authentication

Currently there exists a method to download the agents without authentication, this program will attempt to do this. Below is an example command that you would use to download the database, java, and machine agents all in one go. You will notice I am also using the '-automate' flag that will take some assumptions on what you're searching for by detecting the host environment. Additionally there is the '-o' flag (optional) will place the binaries in the targeted output folder.

```bash
./appd-downloader_mac -automate -db -java -ma -o='agent-folder'
```

Above command output

```bash
Host Details
OS: darwin
Arch: amd64
Following Agent Components will be Downloaded:
        java agent
        database agent
        machine agent
 20.59 MiB / 20.59 MiB [=================================================================================================================] 100.00% 51.10 MiB/s 0s
 147.84 MiB / 147.84 MiB [===============================================================================================================] 100.00% 29.96 MiB/s 4s
 130.51 MiB / 130.51 MiB [===============================================================================================================] 100.00% 43.18 MiB/s 3s
 ```

#### Download via Authentication

Example from below flags without having run this downloader previously. With the create-password flag, it will output an encrypted password and the 'auth' flag that you can later use.

example credentials
    - username: user@appdynamics.com
    - password: password123

```bash
./appd-downloader_mac -username='user@appdynamics.com' -create-password
password123
```

Above command output

```bash
./appd-downloader_mac -username='user@appdynamics.com' -create-password
user: user@appdynamics.com pass:
Password not passed into CLI, what is your AppDynamics Community Password?
password123
Going forward you can pass your encrypted password via CLI as
-encrypted-password='kc7QBWZJMpTmcx7v2fNf9TyoKHTLKtv0gjYw511yjEozaGZBwM3+OjgAqgDhF4XkYehj38Rzd6IN8424Dpc/OiiNRMVdErWy'
Going forward you can pass your encrypted credentials via CLI as
-auth='hCZJA4JA/zGRR79rEGig0eYNjty8c8r3D8LWTFqqPf/EjLR7baFzsAaqWQq1yQkvK99B7n6sFQM62I7TR6GRIgAnEl0LvZk5HRjBRSZWAwZ+Fdm2y+oNwr8=:kc7QBWZJMpTmcx7v2fNf9TyoKHTLKtv0gjYw511yjEozaGZBwM3+OjgAqgDhF4XkYehj38Rzd6IN8424Dpc/OiiNRMVdErWy'
```

You would then execute the below command to download a java agent, notice I am not using '-encrypted-password' and only the '-auth' flag. If you use the '-encrypted-password' flag, you are required to also use the '-username' flag.

```bash
./appd-downloader_mac -auth='hCZJA4JA/zGRR79rEGig0eYNjty8c8r3D8LWTFqqPf/EjLR7baFzsAaqWQq1yQkvK99B7n6sFQM62I7TR6GRIgAnEl0LvZk5HRjBRSZWAwZ+Fdm2y+oNwr8=:kc7QBWZJMpTmcx7v2fNf9TyoKHTLKtv0gjYw511yjEozaGZBwM3+OjgAqgDhF4XkYehj38Rzd6IN8424Dpc/OiiNRMVdErWy' -java
```

#### Authentication Flags

- "username" "AppDynamics Community Username (email)"
- "encrypted-password" "Your Encrypted Password created by this Program via -create-password"
- "decrypted-password" "Your AppDynamics Community Password to be Encrypted"
- "create-password" "Flag to create an Encrypted Password to be used for this program"
- "auth" "Flag that is combined from your Username and Encrypted Password to be used for this program"

#### Platform Components

- "all-platform" "Flag to Download All Platform Components"
- "ec" "Flag to Download Enterprise Console"
- "es" "Flag to Download Events Service"
- "eum" "Flag to Download EUM Server"
- "synthetics" "Flag to Download Synthetic Server"
- "cluster-manager" "Flag to Download Cluster Manager"

#### Agent Components

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
- "synthetic-agent" "Flag to download the Private Synthetic Agent"

#### Automation Assistance Flags

The below flags will assist with increased automation to reduce/remove any extra CLI input

- "detect-host" "Flag to detect Host OS / Arch and reduce binary search results"
- "direct-binary" "Flag to download a binary directly via link produced from previous output"
- "automate" "Flag to make assumptions based upon best practice installations (e.g. only show RPM if available)"
- "tos" "Flag to set the target OS binary type (e.g. -tos=linux)"
- "tbit" "Flag to set the target OS Bit binary type (e.g. -tbit=32)"
- "extension" "Flag to set file extension zip or rpm or other applicable"

#### Other Helpful Flags

- "o" "Flag to set the output folder of binaries, default is current directory"

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
