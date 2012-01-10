This directory contains an example of a package containing a non-trivial
interface that can be mocked with GoMock. The interesting files are:

 *  `user.go`: Source code for the sample package, containing interfaces to be
    mocked. This file depends on the packages named imp[1-4] for various things.

 *  `user_test.go`: A test for the sample package, in which mocks of the
    interfaces from `user.go` are used. This demonstrates how to create mock
    objects, set up expectations, and so on.

Use `goinstall` to get a clone of the GoMock repository into your `$GOPATH`
directory:

    goinstall github.com/dsymonds/gomock

You can build the sample package as follows:

    go build github.com/dsymonds/gomock/sample

To run the test, you'll need to first use MockGen to generate the `mock_user`
package used by the test:

    cd $GOPATH/src/github.com/dsymonds/gomock/sample
    mkdir -p mock_user
    mockgen github.com/dsymonds/gomock/sample Index Embed > mock_user/mock_user.go

You can now verify that the mock package builds:

    go build github.com/dsymonds/gomock/sample/mock_user

You can invoke the following command to run the tests in `user_test`.go:

    go test
