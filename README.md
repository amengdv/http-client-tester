# http-client-tester

## Supported Input Field
- `method:` string
- `url:` string
- `header:` object
- `body:` any

## Supported Input URL Format
- `protocol://domain/path`
- `domain:port/path`

## Supported Expected Output Field
- `is_true:` boolean
- `is_header_contains:` boolean
- `contains:` any
- `json_equal:` object
- `content_type_equal:` string
