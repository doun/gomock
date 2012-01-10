GoMock is a mocking framework for the [Go programming language][golang]. It
integrates well with Go's built-in `testing` package, but can be used in other
contexts too.


Installation
------------

Once you have [installed Go][golang-install], run these commands
to install the `gomock` package and the `mockgen` tool:

    goinstall github.com/dsymonds/gomock/gomock
    goinstall github.com/dsymonds/gomock/mockgen


Documentation
-------------

After installing, you can use `go doc` to get documentation:

    go doc github.com/dsymonds/gomock/gomock

Alternatively, there is an online reference for the package hosted on GoPkgDoc
[here][gomock-ref].


Running mockgen
---------------

The `mockgen` command is used to generate source code for a mock class given a
Go package and a list of its interfaces to be mocked. If the package `foo/bar`
is installed on your system and you want to mock its interfaces `Baz` and `Qux`,
run the following:

    mockgen foo/bar Baz Qux

Source code will be printed to stdout. Use the `-packageOut` flag to control the
name of the resulting package, or the default `mock_bar` will be used. For an
example of the use of `mockgen`, see the `sample/` directory. 


TODO: Brief overview of how to create mock objects and set up expectations, and
an example.

[golang]: http://golang.org/
[golang-install]: http://golang.org/doc/install.html#releases
[gomock-ref]: http://gopkgdoc.appspot.com/pkg/github.com/dsymonds/gomock/gomock
