# http-client-tester (UNDER CONSTRUCTION)

## Supported Input Field
- `method:` string, **required**
- `url:` string, **required**
- `header:` object
- `input_data:` any

## Supported Input URL Format
- `protocol://domain/path`
- `domain:port/path`

## Supported Expected Output Field
- `status_code_equal`: int
- `json_body_equal`: object
- `body_equal`: string
- `body_contains`: string
- `header_contain_key`: string
- `header_contain_value`: string

If a field is not provided, the parser will ignored it

## Usage

To test for one file
```bash
turl <filename>
```

To test for multiple choice of file
```bash
turl <filename> <filename2>
```

To test for all test file in current working directory
```bash
turl .
```
