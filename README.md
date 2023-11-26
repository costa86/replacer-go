# replacer-go

    Replacer
    Replaces entries in a file based on a placeholder
    Placeholder: {{}}
    Usage: ./replacer <json_dict_file> <file_to_modify>
    Version: 1.0.0

Example:

hello.txt
```txt
name = {{name}}
my_age = 50
name = {{name}}
```
hello.json
```json
{
    "name": "jack",
    "age":50
}
```

Run:

    ./replacer hello.json hello.txt

Result:

hello.txt
```txt
name = jack
my_age = 50
name = jack
```