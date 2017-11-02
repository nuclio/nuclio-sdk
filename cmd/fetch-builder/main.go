/*
Copyright 2017 The Nuclio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Fetch nuclio-build from github
package main

import (
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
	ghUser = "nuclio"
	ghRepo = "nuclio-tools"
)

var (
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
	allRels, _, err := cl.Repositories.ListReleases(ctx, ghRepo, ghUser, nil)
	if err != nil {
		return "", err
	}
	var rels []*github.RepositoryRelease
	for _, rel := range allRels {
		if rel.GetDraft() || rel.GetPrerelease() {
			continue
		}
		rels = append(rels, rel)
	}
	sort.Sort(byTime(rels))
	rel := rels[0]
	log.Printf("latest release: %s", rel.GetName())
	arch := fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
	// nuclio-tools-0.1.0-linux-amd64
	name := fmt.Sprintf("%s-%s-%s", ghRepo, rel.GetName(), arch)
	for _, ast := range rel.Assets {
		if strings.HasSuffix(ast.GetBrowserDownloadURL(), name) {
			return ast.GetBrowserDownloadURL(), nil
		}
	}

	return "", fmt.Errorf("can't find release for %s in version %s", arch, rel.GetName())
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

	if runtime.GOOS == "windows" {
		exeFile += ".exe"
	}
	log.Printf("exracting tar to %q", exeFile)

	out, err := os.Create(exeFile)
	if err != nil {
		die(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		die(err)
	}
	if runtime.GOOS != "windows" {
		out.Chmod(0755)
	}
}
