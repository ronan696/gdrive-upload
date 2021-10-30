# gdrive-upload

Go program for uploading files to Google Drive .

## Installation

### Install Go
[Homebrew](http://mxcl.github.io/homebrew/):

```Shell
sudo brew install go
```

[apt](http://packages.qa.debian.org/a/apt.html)-get:

```Shell
sudo apt-get install golang
```

[Install Golang manually](https://golang.org/doc/install)
or
[compile it yourself](https://golang.org/doc/install/source)

### Clone the project
```Shell
git clone https://github.com/ronan696/gdrive-upload.git
```

## Usage

### Setup
1. [Create a new project](https://developers.google.com/workspace/guides/create-project) (or use an existing one) in Google Developer Console.
1. Enable Google Drive API.
1. Create a [Service Account](https://developers.google.com/identity/protocols/oauth2/service-account) for the project. 
1. Download the Service Account key in json format.
1. Rename the json file to `credentials.json` and place it in the go project directory.
1. Modify the below lines in `gdrive-upload.go` with the required information
   ```go
	fileName := "<FileName>"
	mimeType := "<MimeType>"
	parentId := "<ParentID>"
   ```
### Running the program
Execute the below commands in the go project directory to download dependencies and run the program
```Shell
go get -u ./...
go run gdrive-upload.go
```

### Build an executable
Execute the below command in the go project directory to generate an executable file
```Shell
go build
```

## License
[MIT](https://github.com/ronan696/gdrive-upload/blob/master/LICENSE)