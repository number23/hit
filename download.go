package hit

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	fileUrl := "https://golangcode.com/images/avatar.jpg"

	err := DownloadFile("avatar.jpg", fileUrl)
	if err != nil {
		panic(err)
	}

}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func DownloadFileFromFtp(c *ftp.ServerConn, fname string) error {
	r, err := c.Retr(fname)
	defer r.Close()

	if err != nil {
		return err
	}

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fname, buf, 0644)
}
