# template-go-vercel
Template to get started with Go / Golang on Vercel for free!

I created this repo because I found Vercel's documentation quit limited and the repo made by [@riccardogiorato](https://github.com/riccardogiorato/template-go-vercel) also didn't really providing me with answers to the questions I had about how to get started with Go / Golang on Vercel.

### Disclaimer
I don't know Riccardo Giorato but I am sure he is a nice guy, and this is no critique of his work, I am just elaborating on his work.

## Getting Started
I don't expect you to clone this repo, instead I will be taking you through my learnings from building [this hobby project]([https://github.com/kristiannissen/brewblog](https://github.com/kristiannissen/template-go-vercel)) on Vercel, and the starting point for my tutorial will be [this guide](https://vercel.com/docs/functions/runtimes/go) from Golang on how to create a module.

But before we get started you should read the [documentation from Vercel](https://vercel.com/docs/functions/runtimes/go) to get an understanding of the constraints/requirements you will be dealing with. Especially [this part](https://vercel.com/docs/functions/runtimes/go#advanced-go-usage) is interresting.

Go files that reside inside the API folder have to implement the [http.HandlerFunc](https://pkg.go.dev/net/http@go1.22.1#HandlerFunc) signaure.
```
type HandlerFunc func(ResponseWriter, *Request)
```
Which means that if I create a file called hello.go there needs to be a func that implements the http.HandlerFunc signature in that file. Here is an example - disregard the package name in my example, that is not relevant.
```
package api
 
import (
  "fmt"
  "net/http"
)
 
func Hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello Kitty")
}
```
I chose to use the same name for my .go file and the function name, but as you can see in Vercel's documentation, that is not required. If you try to create a .go file in a subfolder that does not contain a function with that signature you will get the following error when deploying
```
Error: Could not find an exported function in "api/pkg/cmd/main.go"
```

## Creating your module
As the documentation from Vercel states, you have to add the go.mod file at the project root level, if you don't it will not be picked up and Go will not be detected.

### Organizing a Go module
This is where I ran into the biggest headache, this part really took me a long time to figure out. I wanted to follow the guidelies from [Organizing a Go module](https://go.dev/doc/modules/layout) for my Go module, but this isn't possible on Vercel because all files that reside inside the api folder have to export the http.HandlerFunc signature.

This means that you cannot store your model in a folder called models, you can't even have a .go file that only contains a struct, that file has to have a function that implements the http.HandlerFunc signature. So you model.go file would look like this
```
package pkg

type Model struct {
	Name string
}

func NewModel() *Model {
	return & Model{Name: "Kitty"}
}

func Pimp(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Pimp says Hello"))
}
```
If you try you will see an error like this "Error: Could not find an exported function in "api/pkg/cmd/main.go""

### Workaround
I don't remember if I read this somewhere or just tried it as part of my endless attempts to figure out how to work around that constraint, but I realised that adding a "_" as a suffix to the folder name worked.

So if I moved my model.go file into "api/_pkg" I could create a model.go file and import that in my hello.go
```
package api

import (
	"net/http"

	p "template-go-vercel/api/_pkg"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)

	model := *p.NewModel()

	resp := "Hello "+ model.Name
	w.Write([]byte(resp))
}
```
I settled on the folder name "pkg" based on the guidelies from [golang standards](https://github.com/golang-standards/project-layout?tab=readme-ov-file#pkg)
So now I have a way to structure my Go project and use it on Vercel!
```
.
├── LICENSE
├── README.md
├── api
│   ├── _pkg
│   │   ├── cmd
│   │   │   └── main.go
│   │   └── model.go
│   ├── hello.go
│   └── hello_test.go
├── cmd
│   ├── api
│   │   └── main.go
│   └── cli
│       └── main.go
├── go.mod
└── public
    └── index.html
```
## Testing
I like writing tests for my code! This folder structure gives me the ability to test the logic inside the different submodules individually, or as I ended up doing here, in a test file inside the _pkg folder. Inside the api folder I stored the files that all have the http.HandlerFunc signature and these functions and the test cases reside in the hello_test.go file.

Because all of the files in the api folder must implement the http.HandlerFunc signature they rely on [net/http/httptest](https://pkg.go.dev/net/http/httptest@go1.22.1) which is why it made sense to me to have a separate test file for those test cases.

### Working locally
I prefer writing tests for my code when I work locally, but I added a small mux that you can run from your command-line, use go run ./cmd/api -port=1234 or which ever port number you prefer to run the API locally. You will have to add the handler functions yourself ;)

**Notes**

You have to check this before deploying anything https://vercel.com/guides/serverless-function-contains-invalid-runtime-error
