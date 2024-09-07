# tURL - HTTP Client For Automated Testing

tURL is a streamlined command-line tool for efficiently testing HTTP servers and APIs.
Designed for **simplicity and flexibility**, tURL lets you define test cases in `JSON` format,
specifying endpoints, methods, headers, bodies, and expected responses.
While **not intended** for complex testing, tURL is ideal for quickly and easily running basic HTTP requests and verifying responses.
It's particularly useful when you find yourself frequently using `curl` but need a more organized and reusable approach to testing.

## Installation

### Linux & MacOS

1. Download the latest releases from Github

```bash
# For Linux
wget https://github.com/amengdv/http-client-tester/releases/download/v1.0.0/turl-linux-amd64 -O turl

# For macOS (Intel)
wget https://github.com/amengdv/http-client-tester/releases/download/v1.0.0/turl-darwin-amd64 -O turl

# For macOS (Apple Silicon)
wget https://github.com/yourusername/turl/releases/download/v1.0.0/turl-darwin-arm64 -O turl
```

2. Run this command

```bash
# Give it permission to execute
chmod +x turl
# Move to $PATH
sudo mv turl /usr/local/bin
```

3. Verify that it is installed

```bash
turl --version
```

### Windows

1. Download the latest binary releases from Github. Use Powershell.

```bash
Invoke-WebRequest -Uri "https://github.com/amengdv/http-client-tester/releases/download/v1.0.0/turl-windows-amd64.exe" -OutFile "C:\path\to\your\folder\turl.exe"
```
> [!IMPORTANT]
> Replace `path/to/your/folder` accordingly 

2. Add the binary to your PATH. WINDOWS STYLE

3. Verify the installation

```bash
turl --version
```

## Supported Input Field
- `name`: string
- `method`: string
- `url`: string, **required**
- `header`: object
- `input_data`: any

## Supported Input URL Format
- `protocol://hostname/path`

## Supported Method
- `GET`
- `POST`
- `PUT`
- `PATCH`
- `DELETE`

> [!NOTE]
> `method` field is default to GET

## Supported Expected Output Field
- `status_code_equal`: int
- `json_body_equal`: object
- `body_equal`: string
- `body_contains`: string
- `header_contain_key`: string
- `header_contain_value`: string

If a field is not provided, the parser will ignored it.

## Usage

To test for one file
```bash
turl <filename>
```

To test for multiple choice of file
```bash
turl <filename> <filename2>
```

To test for all test file in current working directory.
**Prefix** your file's name with `turl_`.
```bash
turl .
```
