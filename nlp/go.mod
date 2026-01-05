module github.com/ardanlabs/nlp

go 1.24.5

// Risks of using 3rd party packages
// - security (ignornace, intentional)
// - bugs
// - compatibility (api changes)
// - Legal (license)
// ~ Might be gone, user can delete packages
//  ```go mod vendor``` - downloads all the dependancy packages - a lot of space


require github.com/stretchr/testify v1.11.1

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
