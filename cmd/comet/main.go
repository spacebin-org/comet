/*
 * Copyright 2021 Luke Whrit, Jack Dorland

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 *   http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/GreatGodApollo/gospacebin"
	arg "github.com/alexflint/go-arg"
	"github.com/spacebin-org/comet/internal/lib"
)

var spacebin = gospacebin.NewClient("https://spaceb.in")

type arguments struct {
	Input     []string `arg:"positional" help:"the file to upload"`
	Raw       bool     `arg:"-r,--raw" help:"whether to return plain-text versions"`
	Copy      bool     `arg:"-c,--copy" help:"copy the url to clipboard after upload"`
	Instance  string   `arg:"env:COMET_INSTANCE" help:"the spacebin instance" default:"https://spaceb.in"`
	ResultURL string   `arg:"env:COMET_RESULT_URL" help:"the base url for the response" default:"https://spaceb.in"`
}

func (arguments) Version() string {
	// @todo Add license information to this message

	return "comet 0.1.0"
}

func main() {
	var args arguments

	arg.MustParse(&args)

	if lib.IsInputFromPipe() {
		reader := bufio.NewReader(os.Stdin)
		var output []rune

		for {
			input, _, err := reader.ReadRune()

			if err != nil && err == io.EOF {
				break
			}

			output = append(output, input)
		}

		for j := 0; j < len(output); j++ {
			fmt.Printf("%c", output[j])
		}
	} else {
		for _, file := range args.Input {
			file, err := filepath.Abs(file)

			if err != nil {
				log.Fatalln(err)
			}

			if lib.FileExists(file) {
				content, err := ioutil.ReadFile(file)

				if err != nil {
					log.Fatalln(err)
				}

				document, err := spacebin.CreateDocument(&gospacebin.CreateDocumentOpts{
					Content:   string(content),
					Extension: lib.Extensions[filepath.Ext(file)],
				})

				if err != nil {
					log.Fatalln(err)
				}

				fmt.Println(document.ID)
			}
		}
	}
}
