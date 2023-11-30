# godo
Why leave your nice and cozy terminal just to write tiny lil todos when you can use godo!

> A simple cli todo written in go.

![GitHub release (latest by date)](https://img.shields.io/github/v/release/ritikdhasmana/godo)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/ritikdhasmana/godo/release.yaml)

![godo-gif](https://github.com/ritikdhasmana/godo/assets/54628046/5c1aa1a1-2ff1-474f-a622-3cc2b02b0c99)

By default, it stores all todos as a json file in user's home directory (`.godo.json`).

## Installation

### Homebrew
```shell
brew tap ritikdhasmana/ritikd
brew install godo
```

### Using Go

```shell
go install github.com/ritikdhasmana/godo/cmd/godo@latest
export PATH=$PATH:$HOME/go/bin  #to add path, restart terminal or do `source ~/.xyzrc`
```

