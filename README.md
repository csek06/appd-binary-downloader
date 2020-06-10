# appd-binary-downloader

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
