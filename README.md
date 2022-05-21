# Initialize Go Project

A command line app to initialize a new go project with proper convention.

### Golang Version

Version 1.18

### Packages

No external package has been used.

### Usage

- At present the only option is to build from source.
- The environment variables `GOPATH` & `GITHUB_USERNAME` must be set.

### Steps to Build App

1. Clone the repository: `git clone https://github.com/yashrajpahwa/init-go-project.git`
2. Go to the directory: `cd init-go-project`
3. Run the command: ` go build -ldflags "-s -w"` to build the app or run: `go install .` to directly install the app.

### License

Licensed under the MIT license.
