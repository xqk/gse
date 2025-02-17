// Copyright 2016 ego authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package gse

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
)

// StopWordMap the default stop words.
var StopWordMap = map[string]bool{
	" ": true,
}

// LoadStopArr load stop word by []string
func (seg *Segmenter) LoadStopArr(dict []string) {
	if seg.StopWordMap == nil {
		seg.StopWordMap = make(map[string]bool)
	}

	for _, d := range dict {
		seg.StopWordMap[d] = true
	}
}

// LoadStop load stop word files add token to map
func (seg *Segmenter) LoadStop(files ...string) error {
	if seg.StopWordMap == nil {
		seg.StopWordMap = make(map[string]bool)
	}

	dictDir := path.Join(path.Dir(seg.GetCurrentFilePath()), "data")
	if len(files) <= 0 {
		dictPath := path.Join(dictDir, "dict/zh/stop_word.txt")
		files = append(files, dictPath)
	}

	name := strings.Split(files[0], ", ")
	if name[0] == "zh" {
		name[0] = path.Join(dictDir, "dict/zh/stop_tokens.txt")
	}

	for i := 0; i < len(name); i++ {
		if !seg.SkipLog {
			log.Printf("Load the stop word dictionary: \"%s\" ", name[i])
		}

		file, err := os.Open(name[i])
		if err != nil {
			log.Printf("Could not load dictionaries: \"%s\", %v \n", name[i], err)
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			text := scanner.Text()
			if text != "" {
				seg.StopWordMap[text] = true
			}
		}
	}

	return nil
}

// AddStop add a token to the StopWord dictionary.
func (seg *Segmenter) AddStop(text string) {
	seg.StopWordMap[text] = true
}

// RemoveStop remove a token from the StopWord dictionary.
func (seg *Segmenter) RemoveStop(text string) {
	delete(seg.StopWordMap, text)
}

// EmptyStop empty the stop dictionary
func (seg *Segmenter) EmptyStop() error {
	seg.StopWordMap = nil
	return nil
}

// IsStop check the word is a stop word.
func (seg *Segmenter) IsStop(s string) bool {
	_, ok := seg.StopWordMap[s]
	return ok
	// return StopWordMap[s]
}
