# http-client-tester

## Supported Input Field
- `method:` string, **required**
- `url:` string, **required**
- `header:` object
- `body:` any

## Supported Input URL Format
- `protocol://domain/path`
- `domain:port/path`

## Supported Expected Output Field
- `is_header_contains:` object
- `contains:` any
- `json_equal:` object
- `content_type_equal:` string

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
