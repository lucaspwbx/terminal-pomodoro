# Terminal Pomodoro

## Installing and executing
```go
go get github.com/lucasweiblen/terminal-pomodoro
go install github.com/lucasweiblen/terminal-pomodoro
terminal-pomodoro
```

## How does it work?
Run the binary and set the flags (time for activity time and rest for rest time).
Default is 25 minutes for work time and 5 minutes for rest time.
In the following configuration, you would have pomodoros with 20 minutes of activity time and 2 minutes to rest:

```shell
terminal-pomodoro -time=20 -rest=2 
```

## User Notifications only work for OSX 10.8 and higher.
If the notifications do not appear it may be a problem with Tmux.
Please refer to [this page](https://github.com/julienXX/terminal-notifier/issues/115)
