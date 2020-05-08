# appd-binary-downloader

## how to compile

```bash
go build
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
