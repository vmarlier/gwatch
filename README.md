# gwatch (go-watch)
Rewriting of the original `watch` linux command in order to add some new feature to it:
- Better view.
- Execute 2 or 3 commands simultaneously.

![Screenshot](./assets/screenshot.png)

## How to use it ?

```sh
# basic execution
gwatch "command"

# interval selection (in seconds)
gwatch "command" -i 5

# multiple commands
gwatch -c "command1" -c "command2"
```

## How to install ?

Clone the repository then:
```sh
$ make build
```
