# Go Lambda App

This project is a Go application designed to run on AWS Lambda. It receives alerts from CloudWatch regarding the CPU usage of RDS instances and sends transformed data to a specified Slack channel.

## Project Structure

```
go-lambda-app
├── src
│   ├── main.go              # Entry point of the application
│   ├── handlers             # Contains alert handling logic
│   │   └── alert_handler.go  # Processes incoming CloudWatch alerts
│   ├── services             # Contains service logic
│   │   └── slack_service.go  # Sends data to Slack
│   └── utils                # Contains utility functions
│       └── transform.go      # Transforms alert data
├── go.mod                   # Module definition file
├── go.sum                   # Dependency checksums
└── README.md                # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone https://github.com/marina-zhelo-smeetz/go-lambda-app.git
   cd go-lambda-app
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Deploy to AWS Lambda:**

1. **Build the Go executable:**
   ```sh
   GOOS=linux GOARCH=amd64 go build -o main main.go

## Usage

- The application listens for CloudWatch alerts. When an alert is received, the `HandleAlert` function in `alert_handler.go` processes the alert and extracts CPU data.
- The extracted data is then transformed using the `TransformData` function in `transform.go`.
- Finally, the transformed data is sent to a Slack channel using the `SendToSlack` function in `slack_service.go`.

## License

This project is licensed under the MIT License. See the LICENSE file for details.