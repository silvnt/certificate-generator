# certifier-go

A web application developed in Go to create and send bulk certificates.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites & Installing

* Go 1.10.3 or later (https://golang.org/dl/);
* wkhtmltopdf 0.12.5 or later (https://wkhtmltopdf.org/downloads.html);
* Go libraries

**Note:** Verify that your GOPATH is set and that the wkhtmltopdf installation was successful through the following commands:
```
go env
wkhtmltopdf -V
```

After installing the previous ones, obtain the following libraries:

* [go-wkhtmltopdf](github.com/SebastiaanKlippert/go-wkhtmltopdf): Golang commandline wrapper for wkhtmltopdf
```
go get -u github.com/SebastiaanKlippert/go-wkhtmltopdf
```
* [sendgrid-go](github.com/sendgrid/sendgrid-go): The Official SendGrid Led, Community Driven Golang API Library
```
go get -u github.com/sendgrid/sendgrid-go
```

Finally, obtain this project through the following command:
```
go get -u github.com/silvnt/certifier-go
```

## Running

Before executing, open the command line in project folder ($GOPATH/src/github.com/silvnt/certifier-go), enter the values of SENDGRID_API_KEY and ALTER_SERVER_ADDRESS (if the latter is empty, the default is 'localhost:3000') in local.env, and run the following command:
```
source local.env
```
**Note:** You don't have to run this command again until you open a new terminal instance.

Now, start the service:
```
go run certifier.go
```

## Authors

* **Silvano Neto**

## Acknowledgments

* João Alexandre (github.com/joaoaneto) for tips of standardization and organization of web projects
* Dayvid Clementino (github.com/dayvidcds) for tips in front-end development
* Even3 for the opportunity to meet this challenge!
