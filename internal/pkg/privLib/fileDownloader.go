package privlib

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	pb "gopkg.in/cheggaaa/pb.v1"
)

const refreshRate = time.Millisecond * 100

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer
// interface and we can pass this into io.TeeReader() which will report progress on each
// write cycle.
type WriteCounter struct {
	n   int // bytes read so far
	bar *pb.ProgressBar
}

/*
CheckCreateFolder Given a destination automatically create a folder if it doesn't exist and return path of that folder
*/
func CheckCreateFolder(outputFolder string) string {
	if strings.HasPrefix(outputFolder, "~/") {

		home, err := homedir.Dir()
		if err != nil {
			// handle err
		}
		fmt.Println(home)
		outputFolder = filepath.Join(home, outputFolder[2:])
	}
	os.MkdirAll(outputFolder, os.ModePerm)
	abs, _ := filepath.Abs(outputFolder)
	return abs + string(os.PathSeparator)
}

/*
FileDownload Given a filename and url, this method will attempt to download a file to a given named location
*/
func FileDownload(filename, url, token string) {

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filename + ".tmp")
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	resp, err := http.Get("http://google.com")
	if len(token) == 0 {
		resp, err = http.Get(url)
	} else {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			// handle err
		}
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			// handle err
		}
	}
	// Get the data
	//resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	defer resp.Body.Close()

	fsize, _ := strconv.Atoi(resp.Header.Get("Content-Length"))

	// Create our progress reporter and pass it to be used alongside our writer
	counter := newWriteCounter(fsize)
	counter.Start()

	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		fmt.Println("error: " + err.Error())
	}

	counter.Finish()
	out.Close()

	err = os.Rename(filename+".tmp", filename)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}

}

func newWriteCounter(total int) *WriteCounter {
	b := pb.New(total)
	b.SetRefreshRate(refreshRate)
	b.ShowTimeLeft = true
	b.ShowSpeed = true
	b.SetUnits(pb.U_BYTES)

	return &WriteCounter{
		bar: b,
	}
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	wc.n += len(p)
	wc.bar.Set(wc.n)
	return wc.n, nil
}

/*
Start starts the timer of writing the file
*/
func (wc *WriteCounter) Start() {
	wc.bar.Start()
}

/*
Finish finishes the timer of writing the file
*/
func (wc *WriteCounter) Finish() {
	wc.bar.Finish()
}
