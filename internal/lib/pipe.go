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

package lib

import (
	"errors"
	"os"
)

func FileExists(filepath string) bool {
	info, e := os.Stat(filepath)

	if os.IsNotExist(e) {
		return false
	}

	return !info.IsDir()
}

func IsInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()

	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func GetFile(filePath string) (*os.File, error) {
	if !FileExists(filePath) {
		return nil, errors.New("the file provided does not exist")
	}

	file, e := os.Open(filePath)

	if e != nil {
		return nil, e
	}

	return file, nil
}
