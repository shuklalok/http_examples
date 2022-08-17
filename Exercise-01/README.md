## Reference
[Youtube Reference](https://www.youtube.com/watch?v=VzBGi_n65iU&list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_&index=1)

## Run server
```bash
go run main.go
```
## Test with curl
```bash
curl localhost:7777/
```

```bash
curl localhost:7777/bye
```

```bash
curl localhost:7777/ -d Alok
```
```bash
curl localhost:7777/bye -d Alok
````

## Test with ```REST Client```
- Install ```REST Client``` VSCODE extension.
- test using the ```test.http``` file.