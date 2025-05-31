## Hoot

Hoot is Promodoro like timer app which remind you to look away from screen. for now this is only compatible with linux systems. 

## Installation
```
go install github.com/codeshaine/hoot@latest

```

## Usage

```bash
hoot start --interval 10s 
hoot stop
hoot status

```
## Configuration (optional)
hoot make use of config file for custom notification sound which reside in your
```
$HOME/.config/hoot/config.yml

```
Current prop supported
```yml
sound: 
interval: 10s

```



