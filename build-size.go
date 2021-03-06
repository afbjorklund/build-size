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
	link := links[0]

	file, err := os.Open(link)
	defer file.Close()
	if err != nil {
		panic(err)
	}

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

		line = strings.Replace(line, "packagefile ", "", 1)
		fields := strings.Split(strings.TrimRight(line, "\n"), "=")
		module := fields[0]
		library := fields[1]

		lib, err := os.Open(library)
		defer lib.Close()
		if err != nil {
			panic(err)
		}
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
