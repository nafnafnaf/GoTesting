# My Go Project

This is a simple Go project that demonstrates how to create an executable application using Go.

## Project Structure

```
my-go-project
├── cmd
│   └── main.go        # Entry point of the application
├── pkg
│   └── utils
│       └── helper.go  # Utility functions for basic arithmetic operations
├── go.mod             # Module definition file
└── README.md          # Project documentation
```

## Getting Started

### Prerequisites

Make sure you have Go installed on your machine. You can download it from the official Go website.

### Building the Application

To build the application, navigate to the project directory and run:

```
go build -o my-go-app ./cmd
```

This will create an executable named `my-go-app` in the project directory.

### Running the Application

To run the application, use the following command:

```
./my-go-app
```

### Usage

The application currently supports basic arithmetic operations through the utility functions defined in the `pkg/utils/helper.go` file. You can modify the `main.go` file to implement specific functionalities as needed.

## Contributing

Feel free to submit issues or pull requests if you would like to contribute to this project.