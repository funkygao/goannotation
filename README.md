goannotation
============

Java annotation feature introduced into golang.

Users implement plugin/annotation in annotations/ directory to auto generate go source code based on annotation tags.

Usually work with go generate tool.


### Example

    go build
    go generate
    ./goannotation contrib/sample.go

