// Copyright 2010 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// mockgen generates mock implementations of Go interfaces. The package to
// which the mocked interfaces belong must be installed beforehand.
//
// To mock interfaces Tweaker and Frobnicator from the package
// github.com/foo/bar, you invoke mockgen as follows:
//
//     mockgen github.com/foo/bar Tweaker Frobnicator
//
// Output suitable for saving to a .go file will be written to stdout.

package main

import (
	"flag"
	"fmt"
	"github.com/dsymonds/gomock/generate"
	"io"
	"log"
	"reflect"
)

const (
	gomockImportPath = "github.com/dsymonds/gomock/gomock"
)

var (
	packageOut  = flag.String("package", "", "Package of the generated code; defaults to the package of the input file with a 'mock_' prefix.")
)

func main() {
	flag.Parse()

	// TODO(jacobsa): Pay attention to the package and interface arguments.
	var someReaderPtr *io.Reader
	readerType := reflect.TypeOf(someReaderPtr).Elem()
	types := []reflect.Type{readerType}

	output, err := generate.GenerateMockSource("mock_io", types)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Print(output)
}
