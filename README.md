# SSL Checker

SSL Checker is a simple Go (Golang) application for checking the validity of SSL certificates for a given URL and port. It utilizes the Fiber web framework and the crypto/tls package to perform SSL certificate checks.

## Features

- Verify SSL certificate for a given URL and port.
- Return the certificate's expiry date and organization information.

## Getting Started

### Prerequisites

- Go (Golang) installed on your system.

### Installation

1. Clone this repository:

 ```bash
 git clone https://github.com/yourusername/ssl-checker.git
 ```

2. Change the working directory:

 ```bash
 cd ssl-checker
 ```

3. Build the app

 ```bash
 go build
 ```

4. Run 

 ```bash
 ./ssl-checker
 ```

### Usage

You can send a JSON POST request to the /api/ssl-check endpoint with the URL and PORT you want to check.

Example request:

```json
{
  "url": "google.com",
  "port": "443"
}
```

Example response:

```json
{
  "organization": "Google Trust Services LLC",
  "ssl_expiry": "Monday, 01-Jan-24 08:04:02 UTC"
}
```
