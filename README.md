# GO-Carbon

Convert source code from GH gists into shareable images.<br>
Inspired from the [Carbon](https://github.com/carbon-app/carbon) project.

This project is an effort to build a REST API implimentation of the same in Go. It does not completely work as the original project, and many parameters are yet to be added.

## Build

I use [air](https://github.com/cosmtrek/air) for hot reloading.
Otherwise:

```bash
go build -o server . && ./server
```

## Query Params

Endpoint: `/api`

| Query Param | Description         |
|-------------|---------------------|
| `gistid`    | Gist ID of the file |
| `fsize`     | Font size           |
| `fcolor`    | Font color          |
| `bgcolor`   | Background color    |

## Status Codes

- `404`: Invalid Gist ID
- `400`: Missing query params
