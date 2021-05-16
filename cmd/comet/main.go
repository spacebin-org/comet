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
	"github.com/alexflint/go-arg"
	"github.com/spacebin-org/comet/internal/lib"
)

type arguments struct {
	Input     []string `arg:"positional" help:"the file to upload"`
	Raw       bool     `arg:"-r,--raw" help:"whether to return plain-text versions"`
	Copy      bool     `arg:"-c,--copy" help:"copy the url to clipboard after upload"`
	Instance  string   `arg:"env:COMET_INSTANCE" help:"the spacebin instance" default:"https://spaceb.in"`
	ResultURL string   `arg:"env:COMET_RESULT_URL" help:"the base url for the response" default:"https://spaceb.in"`
}

var args arguments

func init() {
	arg.MustParse(&args)
}

func (arguments) Version() string {
	// @todo Add license information to this message

	return "comet 0.1.0"
}

func main() {
	spacebin := &lib.Spacebin{
		InstanceURL: args.Instance,
		ResultURL:   args.ResultURL,
	}

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

		idString, err := spacebin.Upload(&gospacebin.CreateDocumentOpts{
			Content:   lib.RunesToString(output),
			Extension: "txt",
		})

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(idString)
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

				idString, err := spacebin.Upload(&gospacebin.CreateDocumentOpts{
					Content:   string(content),
					Extension: lib.Extensions[filepath.Ext(file)],
				})

				if err != nil {
					log.Fatalln(err)
				}

				fmt.Println(idString)
			}
		}
	}
}
