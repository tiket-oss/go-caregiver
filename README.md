# Go-Caregiver

Go-Caregiver wraps an application with log rotation capability and provide environment variables via configuration file.

## Installation

To install this application, make sure you have go installed in your machine, and run the following command:

```sh
go get -v github.com/tiket-libre/go-caregiver
```

## Examples

Suppose you have a binary to run your web application called `dummy-server`, and to pass configuration to your application, you only supports environment variables; and you don't have a log policy defined in your application, as you think that is not the main purpose of your application, and you shouldn't be burdened from it.

### Configuration

Go-Caregiver aims to help operations on running said application using a file, containing the application's variables, and how the log rotation policy should be. A sample configuration will look like this.

```json
{
    "service": {
        "VARIABLE": "example",
        "OTHER_VARIABLE": 12,
        "ANOTHER_VARIABLE": 15.5182379128379123
    },
    "log": {
        "filename": "output.log",
        "maxsize": 10,
        "maxbackups": 0,
        "maxage": 0,
        "compress": false
    }
}
```

### Usage

To run go-caregiver, provide the path to configuration file using the `CONFIG` variable and then pass any command you want to wrap after the `go-caregiver` command.

```sh
CONFIG=config.json go-caregiver dummy-server
```
