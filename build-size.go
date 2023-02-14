package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type sizeMap map[string]int64

func main() {
	work := os.Getenv("WORK")
	links, err := filepath.Glob(work + "/*/importcfg.link")
	if err != nil {
		panic(err)
	}
	// assume library being built with -buildmode=archive
	if len(links) == 0 {
		links = []string{work + "/b001/importcfg"}
		if _, err := os.Stat(links[0]); err != nil {
			panic(err)
		}
	}
	link := links[0]

	file, err := os.Open(link)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sizes := make(sizeMap)

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if !strings.HasPrefix(line, "packagefile") {
			continue
		}
		line = strings.Replace(line, "packagefile ", "", 1)
		fields := strings.Split(strings.TrimRight(line, "\n"), "=")
		module := fields[0]
		library := fields[1]

		lib, err := os.Open(library)
		if err != nil {
			panic(err)
		}
		defer lib.Close()
		fi, err := lib.Stat()
		if err != nil {
			panic(err)
		}

		sizes[module] = fi.Size()
	}

	type kv struct {
		Key   string
		Value int64
	}

	var ss []kv
	for k, v := range sizes {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(a, b int) bool {
		return ss[a].Value > ss[b].Value
	})

	for _, kv := range ss {
		module := kv.Key
		size := kv.Value
		fmt.Printf("%d\t%s\n", size/1024, module)
	}
}
