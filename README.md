

## Default curl request
```
curl -d '{"source":"<button>Hello</button>", "viewport": {"width": null, "height": null}}' -H "Content-Type: application/json" -X POST http://localhost:3001/chrome/screenshot/ >> example.png

curl -d '{"source":"<button>Hello</button>", "viewport": {"width": 100, "height": 100}}' -H "Content-Type: application/json" -X POST http://localhost:3001/chrome/screenshot/ >> example.png

```
## Curl request with file

- Install NodeJS >= 16
- `cd template`
- `npm run genarate`

```
curl -d '@template/request.json' -H "Content-Type: application/json" -X POST http://localhost:3001/chrome/screenshot/ >> example.png
```