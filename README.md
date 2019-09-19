# formautomator

Create HTML Forms Automatically

JSON metadata

```json
{
    "method": "POST",
    "action": "/addr",
    "fields": [
        {
            "name": "address",
            "label": "Address",
            "class": "form-control",
            "type": "text",
            "placeholder": "Sunset Blvd, 38"
        }
    ]
}
```

## Usage examples

```console
cat examples/example.json| go run cmd/fa/main.go -t ./templates
```

```console
go run cmd/fa/main.go -f examples/example.json -t ./templates
```
