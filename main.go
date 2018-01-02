// Convert slideshare presentations to PDF in landscape mode, for the ones that have the download link disabled.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		panic("need slideshare full urls as args")
	}
	cli := &http.Client{
		Timeout: time.Second * 20,
	}
	for _, arg := range os.Args[1:] {
		url, err := url.Parse(arg)
		if err != nil || url.Host != "www.slideshare.net" {
			log.Printf("invalid slideshare url `%s`\n", url)
			continue
		}
		log.Printf("Processing %s\n", url)
		err = run(cli, url)
		if err != nil {
			log.Printf("Could not process url `%s` because of `%s`\n", url, err.Error())
		}
	}
}

func run(cli *http.Client, url *url.URL) error {
	imageURLs, err := scrap(cli, url.String())
	if err != nil {
		return err
	}

	reader, err := newHTTPReader(cli)
	if err != nil {
		return err
	}

	dir, err := ioutil.TempDir("", "slideshare2kindle")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(dir)
	}()
	writer, err := newFileWriter(dir)
	if err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	ch := make(chan string)
	for i := 0; i < runtime.NumCPU(); i++ {
		go readWrite(ch, wg, reader, writer)
	}

	numImages := len(imageURLs)
	log.Printf("Gathering %d images from %s\n", numImages, url.String())
	wg.Add(numImages)
	for i, image := range imageURLs {
		ch <- fmt.Sprintf("%02d%s", i, image)
	}
	wg.Wait()

	var imagesPaths []string
	for i := 0; i < numImages; i++ {
		imagesPaths = append(imagesPaths, fmt.Sprintf("%s/%02d.jpg", dir, i))
	}
	pdfPath := fmt.Sprintf("%s.pdf", path.Base(url.Path))
	log.Printf("Creating output PDF\n")
	err = imagesToPDF(imagesPaths, pdfPath)
	if err == nil {
		log.Printf("Successfully created %s\n", pdfPath)
	}
	return err
}
