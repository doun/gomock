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
	"bytes"
	"flag"
	"fmt"
	"log"
	"text/template"

	// Make sure goinstall installs the generate package, even though it's not
	// used directly here. It's used indirectly through generated code.
	_ "github.com/dsymonds/gomock/generate"
)

const (
	gomockImportPath = "github.com/dsymonds/gomock/gomock"
)

var (
	packageOut  = flag.String("package", "", "Package of the generated code; defaults to the package of the input file with a 'mock_' prefix.")
)

// srcTemplate is a template for the source code generated by mockgen that is
// compiled and run to itself generate the mock class source code. We must do
// this because there is no way to get ahold of reflect.Type values for
// arbitrary interfaces at runtime.
const srcTemplate =
`
package main

import (
	"fmt"
	"github.com/dsymonds/gomock/generate"
	"reflect"
	pkg "{{.Pkg}}"
)

func main() {
	numTypes := {{.InterfaceNames | len}}
	types := make([]reflect.Type, numTypes)

	{{range $i, $typeName := .InterfaceNames}}
	var ptr{{$i}} *pkg.{{$typeName}}
	types[{{$i}}] := reflect.TypeOf(ptr{{$i}}).Elem()
	{{end}}

	output, err := generate.GenerateMockSource("{{.OutputPkg}}", types)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Print(output)
}
`

type templateArg struct {
	Pkg string
	InterfaceNames []string
	OutputPkg string
}

func main() {
	flag.Parse()

	// Turn off dates in logging output.
	log.SetFlags(0)

	// Parse the command-line args.
	if flag.NArg() < 2 {
		log.Fatal("Usage: mockgen <package> interface [...]")
	}

	arg := &templateArg{
		Pkg: flag.Arg(0),
		InterfaceNames: make([]string, flag.NArg()-1),
		OutputPkg: "mock_" + flag.Arg(0),
	}

	// Handle non-standard package names.
	if *packageOut != "" {
		arg.OutputPkg = *packageOut
	}

	// Copy interface names.
	for i := 1; i < flag.NArg(); i++ {
		arg.InterfaceNames[i-1] = flag.Arg(i)
	}

	// Parse the generated code template, and run it with the arg struct.
	t := template.Must(template.New("code").Parse(srcTemplate))

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, arg); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	// TODO(jacobsa): Compile this instead.
	fmt.Print(buf.String())
}
