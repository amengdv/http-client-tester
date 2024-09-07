# tURL - HTTP Client For Automated Testing

tURL is a streamlined command-line tool for efficiently testing HTTP servers and APIs.
Designed for **simplicity and flexibility**, tURL lets you define test cases in `JSON` format,
specifying endpoints, methods, headers, bodies, and expected responses.
While **not intended** for complex testing, tURL is ideal for quickly and easily running basic HTTP requests and verifying responses.
It's particularly useful when you find yourself frequently using `curl` but need a more organized and reusable approach to testing.

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
