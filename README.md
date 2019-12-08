# Serverless Deployment with AWS API Gateway & Lambda

The basic steps...

### Create REST API(s) in API Gateway

- Create API
- Create resource (with path parameters as needed)
- Create method
	- Integration type: Lambda Function
	- Use Lambda Proxy integration: checked
	- Lambda Region: (make sure it matches with the region in Lambda)
	- Lambda Function: (the function that was created in Lambda)

- Deploy API
	- Deployment stage: (new or existing stage)

### Create Lambda Function(s)

- Create function: Author from scratch
- Function name: (your-lambda-function-name)
- Runtime: Go 1.x

After a function is created, then:

- Add trigger: (the API created in API Gateway above)
- Create executable binary file from your local function code (as shown below) and upload to Lambda Function

### Prepare the local function code

`go get` these 2 modules for your function code

```
$ go get github.com/aws/aws-lambda-go/events
$ go get github.com/aws/aws-lambda-go/lambda
```

Example:

```
func main() {
    lambda.Start(yourFunction)
}

func yourFunction(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	hello = "Hello World!"

	// Return a response with a 200 OK status
	// and the "Hello World!" string in the body.
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       hello,
	}, nil
}
```

### Create executable binary file and place in a .zip file

(instructions from AWS Documentation)

On Linux/Mac

```
# Remember to build your handler executable for Linux!
$ GOOS=linux GOARCH=amd64 go build -o main main.go
$ zip main.zip main
```

On Windows, use the **build-lambda-zip** because Windows may have trouble producing a zip file that marks the binary as executable on Linux.

```
$ go get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
```

then use the tool from your `GOPATH` to create a proper .zip file. If you have a default installation of Go, the tool will be in `%USERPROFILE%\Go\bin`

```
# if using cmd.exe

$ set GOOS=linux
$ set GOARCH=amd64
$ set CGO_ENABLED=0
$ go build -o main main.go
$ %USERPROFILE%\Go\bin\build-lambda-zip.exe -o main.zip main
```

```
# if using PowerShell

$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"
go build -o main main.go
~\Go\Bin\build-lambda-zip.exe -o main.zip main
```

```
# if using Git Bash (my preferred way)

$ GOOS=linux GOARCH=amd64 go build -o main main.go
~/go/bin/build-lambda-zip.exe -o main.zip main
```

The newly created .zip file is ready to upload to Lambda Function

### More

Check out AWS official documentation for more details:
https://docs.aws.amazon.com/lambda/latest/dg/welcome.html