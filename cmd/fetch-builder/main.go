// Fetch nuclio-build from github
package main

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"sort"
	"strings"

	"github.com/google/go-github/github"
)

const (
	ghUser  = "nuclio"
	ghRepo  = "nuclio-tools"
	exeFile = "nuclio-build"
)

// Sort
type byTime []*github.RepositoryRelease

func (rs byTime) Len() int      { return len(rs) }
func (rs byTime) Swap(i, j int) { rs[i], rs[j] = rs[j], rs[i] }
func (rs byTime) Less(i, j int) bool {
	ti := rs[i].GetPublishedAt().Time
	tj := rs[j].GetPublishedAt().Time
	// Sort in descending order
	return tj.Before(ti)
}

func fetchURL() (string, error) {
	cl := github.NewClient(nil)
	ctx := context.Background()
	log.Printf("getting releases")
	rels, _, err := cl.Repositories.ListReleases(ctx, ghRepo, ghUser, nil)
	if err != nil {
		return "", err
	}
	sort.Sort(byTime(rels))
	rel := rels[0]
	log.Printf("latest release: %s", rel.GetName())
	arch := fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
	// nuclio-tools-v0.1.0-linux-amd64.tar.gz
	name := fmt.Sprintf("%s-%s-%s.tar.gz", ghRepo, rel.GetName(), arch)
	for _, ast := range rel.Assets {
		if strings.HasSuffix(ast.GetBrowserDownloadURL(), name) {
			return ast.GetBrowserDownloadURL(), nil
		}
	}

	return "", fmt.Errorf("can't find release for %s in version %s", arch, rel.GetName())
}

func extract(rdr io.Reader, wtr io.Writer) error {
	gz, err := gzip.NewReader(rdr)
	if err != nil {
		return err
	}
	defer gz.Close()
	tr := tar.NewReader(gz)
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if h.Name == exeFile {
			_, err := io.Copy(wtr, tr)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf(`can't find %q in tar`, exeFile)
}

func die(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "be more verbose")
	flag.Usage = func() {
		fmt.Printf("usage: %s\n\n", path.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()

	if !verbose {
		log.SetOutput(ioutil.Discard)
	}

	url, err := fetchURL()
	if err != nil {
		die(err)
	}
	resp, err := http.Get(url)
	if err != nil {
		die(err)
	}
	defer resp.Body.Close()

	log.Printf("exracting tar to %q", exeFile)

	out, err := os.Create(exeFile)
	if err != nil {
		die(err)
	}
	defer out.Close()

	if err = extract(resp.Body, out); err != nil {
		die(err)
	}
	if runtime.GOOS != "windows" {
		out.Chmod(0755)
	}
}
