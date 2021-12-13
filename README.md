# gwatch (go-watch)
Rewriting of the original `watch` linux command in order to add some new feature to it:
- Better view.
- Execute **2, 3 or 4** commands simultaneously.

## Why ?

I'm working on Kubernetes every single day. And my team and I are using a lot the `watch` tool to monitor if our last change was well applied by Flux for example.
I usually need to open 2, 3 in my term/tmux to be able to monitor everything. With this tool, we will be able to monitor 2 or 3 things at once.

## What does it look like ?

![Screenshot1](./assets/screenshot-1.png)
![Screenshot2](./assets/screenshot-2.png)
![Screenshot3](./assets/screenshot-3.png)
![Screenshot4](./assets/screenshot-4.png)

## How to use it ?

```sh
# basic execution
gwatch "command"

# interval selection (in seconds)
gwatch "command" -i 5

# multiple commands
gwatch "command1" "command2" "command3" "command4"
or
gwatch -c "command1" -c "command2" -c ...
```

## How to install ?

Clone the repository then:
```sh
$ make build
```

## Incoming features ?
- :ballot_box_with_check: Multicommand support.
- :ballot_box_with_check: Change interval support.
- :radio_button: Exec commands asynchronously.
- :radio_button: Highlight differences with -d or --differences flag
