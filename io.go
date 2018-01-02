package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"math"
	"net/http"
	"os"
	"sync"

	"github.com/BurntSushi/graphics-go/graphics"
	"github.com/pkg/errors"
	"github.com/unidoc/unidoc/pdf/creator"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func scrap(cli *http.Client, url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("invalid HTTP status %d", resp.StatusCode)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	matcher := func(n *html.Node) bool {
		if n.DataAtom != atom.Img {
			return false
		}
		return scrape.Attr(n, "class") == "slide_image"
	}
	images := scrape.FindAll(root, matcher)
	var URLs []string
	for _, image := range images {
		URLs = append(URLs, scrape.Attr(image, "data-full"))
	}

	return URLs, nil
}

func imagesToPDF(inputPaths []string, outputPath string) error {
	c := creator.New()

	for _, imgPath := range inputPaths {
		img, err := creator.NewImageFromFile(imgPath)
		if err != nil {
			return err
		}
		img.ScaleToWidth(612.0)

		// Use page width of 612 points, and calculate the height proportionally based on the image.
		// Standard PPI is 72 points per inch, thus a width of 8.5"
		height := 612.0 * img.Height() / img.Width()
		c.SetPageSize(creator.PageSize{612, height})
		c.NewPage()
		img.SetPos(0, 0)
		_ = c.Draw(img)
	}

	return c.WriteToFile(outputPath)
}

func readWrite(ch chan string, wg *sync.WaitGroup, r reader, w writer) {
	for URL := range ch {
		number, url := URL[0:2], URL[2:]
		data, err := r.read(url)
		if err != nil {
			log.Println(err)
		} else {
			err = w.write(number+".jpg", data)
			if err != nil {
				log.Println(err)
			}
		}
		wg.Done()
	}
}

// A reader returns the content of an image by file name.
type reader interface {
	read(string) ([]byte, error)
}

// A writer writes the content of an image with a file name and type.
type writer interface {
	write(string, []byte) error
}

// errors
var (
	errInvalidClient = errors.New("invalid client")
	errEmptyPath     = errors.New("empty images dir path")
)

// errReadNetwork is an network error for a certain image.
type errReadNetwork string

func (e errReadNetwork) Error() string {
	return fmt.Sprintf("network error reading image `%s`", string(e))
}

// httpReader reads the image url.
type httpReader struct {
	cli *http.Client
}

// newHTTPReader returns a new reader with the provided client.
func newHTTPReader(cli *http.Client) (*httpReader, error) {
	if cli == nil {
		return nil, errInvalidClient
	}
	return &httpReader{cli}, nil
}

func (h *httpReader) read(imageURL string) ([]byte, error) {
	resp, err := h.cli.Get(imageURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, errReadNetwork(imageURL)
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	dimensions := img.Bounds()
	rotated := image.NewRGBA(image.Rect(0, 0, dimensions.Dy(), dimensions.Dx()))
	err = graphics.Rotate(rotated, img, &graphics.RotateOptions{math.Pi / 2.0})
	if err != nil {
		return nil, err
	}

	var buff bytes.Buffer
	err = jpeg.Encode(&buff, rotated, &jpeg.Options{Quality: jpeg.DefaultQuality})

	return buff.Bytes(), err
}

// A fileWriter writes an image as a file in a path.
type fileWriter struct {
	path string
}

// newFileWriter returns a FileWriter that writes to the provided path.
func newFileWriter(path string) (*fileWriter, error) {
	if path == "" {
		return nil, errEmptyPath
	}
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return &fileWriter{path}, nil
}

func (fw *fileWriter) write(imageName string, data []byte) error {
	err := os.MkdirAll(fw.path, os.ModePerm)
	if err != nil {
		return err
	}

	fpath := fmt.Sprintf("%s/%s", fw.path, imageName)
	f, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	_, err = f.Write(data)

	return err
}
