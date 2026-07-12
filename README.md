# unix-shell

A small Unix-like shell written in Go.

This project runs an interactive prompt, executes built-in commands, resolves external programs from `PATH`, and supports basic shell parsing features such as quoting, escaping, tilde expansion, and output redirection.

## What it supports

- Built-ins: `exit`, `echo`, `pwd`, `type`, `cd`
- External command execution through `PATH`
- Single quotes and double quotes
- Backslash escaping
- `~` expansion for the home directory
- Redirection for stdout and stderr:
  - `>`
  - `>>`
  - `1>`
  - `1>>`
  - `2>`
  - `2>>`

## Project layout

```text
app/
  main.go              interactive shell loop
  commands/
    tokens.go          command parsing and redirection detection
    run.go             external command execution
    helpers.go         output helpers for built-ins
```

## Run locally

Requirements:

- Go installed locally

Start the shell:

```sh
./run.sh
```

If you prefer running Go directly:

```sh
go run ./app
```

## Example commands

```sh
pwd
echo hello world
type echo
cd /tmp
ls
echo "hello" > out.txt
cat missing.txt 2> errors.txt
echo again >> out.txt
```

## Notes

- Built-in output redirection is handled inside the shell.
- External commands inherit the current terminal unless redirected.
- On Windows/MSYS environments, append redirection is handled explicitly for better compatibility with child processes.
