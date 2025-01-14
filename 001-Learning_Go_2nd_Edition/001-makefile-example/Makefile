# The .DEFAULT_GOAL defines which target is run when no target is specified.
# In our case, we are going to run the build target.
.DEFAULT_GOAL := build

# The .PHONY line keeps make from getting confused if there is a directory or
# file in your project with same name as one of the listed targets.
.PHONY: fmt vet build clean

# Next we have the target definitions. The word before the colon (:) is the name of the target.
# Any words after the target (like vet in the line build: vet) are the other targets that must
# be run before the specified target runs. The tasks that are performed by the target are on the
#indented lines after the target.

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

clean:
	go clean