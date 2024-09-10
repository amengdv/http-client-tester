# tURL - HTTP Client For Automated Testing

tURL is a streamlined command-line tool for efficiently testing HTTP servers and APIs.
Designed for **simplicity and flexibility**, tURL lets you define test cases in `JSON` format,
specifying endpoints, methods, headers, bodies, and expected responses.
While **not intended** for complex testing, tURL is ideal for quickly and easily running basic HTTP requests and verifying responses.


It's particularly useful when you find yourself frequently using `curl` 
but need a more organized and reusable approach to testing. This tool is built because
I want to test my endpoint but found myself going back and forth using `curl` and it became hard
to edit text and keep track of my previous command. And I don't want to
use other HTTP client or GUI Client because I just want something simple and quick.

## Installation

### Linux & MacOS

1. Download the latest releases from Github

**LINUX**
```bash
# For Linux
curl -L -o turl https://github.com/amengdv/http-client-tester/releases/latest/download/turl-linux-amd64
# Give it permission to execute
chmod +x turl
# Move to $PATH
sudo mv turl /usr/local/bin
```

**MacOS (Intel)**
```bash
# For macOS (Intel)
curl -L -o turl https://github.com/amengdv/http-client-tester/releases/latest/download/turl-darwin-amd64
# Give it permission to execute
chmod +x turl
# Move to $PATH
sudo mv turl /usr/local/bin
```

**MacOS (Apple Silicon)**
```bash
# For macOS (Apple Silicon)
curl -L -o turl https://github.com/yourusername/turl/releases/latest/download/turl-darwin-arm64
# Give it permission to execute
chmod +x turl
# Move to $PATH
sudo mv turl /usr/local/bin
```

2. Verify that it is installed

```bash
turl --version
```

### Windows

1. Download the latest binary releases from Github. Use Powershell. It is easier
to just download the executable on the release pages. But if you're feeling extra here
is the command

```bash
Invoke-WebRequest -Uri "https://github.com/amengdv/http-client-tester/releases/latest/download/turl-windows-amd64.exe" -OutFile "C:\path\to\your\folder\turl.exe"
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
- `header`: object of lists (look at example usage)
- `input_data`: any
- `show_body`: boolean default=true

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

### Example test file

**turl_example.json**
```json
{
    "tests": 
    [
        {
            "name": "test",
            "method": "get",
            "header": {
                "Authorization": [
                    "Bearer yourtoken"
                ]
            },
            "url": "http://localhost:8080/",
            "body_equal": "OK",
            "status_code_equal": 200
        }
    ]
}
```
> [!IMPORTANT]
> To begin writing the test case, start with a field `"tests"`. It is a list of
objects i.e your test cases

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

> [!IMPORTANT]
> It is recommended to add the test file in `.gitignore` so you don't expose
any sensitive information just in case.

## Contributing 🤝

1. Clone the repo
```bash
git clone https://github.com/amengdv/http-client-tester.git
cd http-client-tester
```

2. Build the project. You need `go` CLI for this
```bash
go build -o turl
```

3. Run
```bash
./turl --version
```

> [!NOTE]
> If you want to contribute, fork the repository and open a pull request to the
`main` branch
