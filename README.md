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
sound: /share/local/somefile.mp3 #example path
interval: 10s #use 10s, 10m, 10h format
```

## Contributing

If your interested in contributing then fork the repo and do a PR.
You guys can work on these if you are interested

- linux compatible (I'm not sure notification and sound works in all linux distros)
- windows compatible
- mac compatible
