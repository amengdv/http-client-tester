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
- `header_contain_key:` object
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
