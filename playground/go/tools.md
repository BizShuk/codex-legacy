# Tools

##### Go build
Now, test that the package compiles with go build:
`go build github.com/user/stringutil`
Or, if you are working in the package's source directory, just:

`go build`
This won't produce an output file.


##### Go get
go get [options] <package> , download packaget to $GOPATH with hierarchical path
options:
-d , download only                                                               
~                                   


##### Go doc
Get comments before package name. If there is a bug in some function, `// BUGS(who)` comment will be captured ,too


##### Go install
places the package object inside the pkg directory of the workspace.
If it's a executable program , place it inside the bin directory
