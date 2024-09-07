# Shopify Customer Data Export Tool

This is a tool written in Go for exporting customer data from multiple Shopify stores.

## Features
- Export customer data concurrently from multiple Shopify stores
- Save the exported data in JSON format
- Support pagination for handling large amounts of customer data
- Error handling and logging

## Installation
1. Make sure you have Go 1.22.6 or higher installed.
2. Clone this repository:

```
git clone https://github.com/yourusername/shopify-customer-dump.git
```

3. Navigate to the project directory:
```
cd shopify-customer-dump
```
4. Install dependencies:
```
go mod tidy
```

## Configuration

In the project directory, create a file named `config.shopify.yaml` and add the following content:

```yaml
shopify:
  - name: storename
    accessToken: "your_access_token"
```
## Usage

Run the following command to start the program:
```
go run main.go
```

Or use the binary file:
```
./shopify-customer-dump
```

The exported data will be saved in the `.cache` directory, with each store having its own subdirectory.

## Project Structure

- `main.go`: Main entry point of the program
- `internal/`: Internal logic implementation
  - `conf.go`: Configuration file reading
  - `dump.go`: Data export logic
- `config.shopify.yaml`: Configuration file

## Dependencies
- `github.com/bold-commerce/go-shopify/v4`: For interacting with the Shopify API
- `gopkg.in/yaml.v3`: For parsing YAML configuration files

## Performance
Under stable network conditions, a single store can handle at least 200,000 records per hour. Actual processing speed may vary depending on network conditions, device specifications, API limitations, etc. This is not a guaranteed result, but a reference.

## Notes

- Ensure that your Shopify access token has the necessary permissions to read customer data.
- Exporting large amounts of data may take a considerable amount of time, please be patient.
- Errors and progress information during the export process will be logged in the `logs.txt` file.

## Contributing

Feel free to submit issues and suggestions for improvements.
