
# look_away

Simple 20-20-20 rule timer app for the terminal written in go.

## Install

```sh
git clone https://github.com/jnsltk/look_away.git && cd look_away
go install ./cmd/look_away
```

Make sure that the Go bin directory ($GOPATH/bin or $HOME/go/bin) is in your system’s PATH. You can check if it’s included by running:

```sh
echo $PATH
```

If it's not in your path, add it to your shell config (`~/.bashrc`, `~/.zshrc`, etc.):

```sh
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.zshrc
source ~/.zshrc  # or source ~/.bashrc
```

## Usage

```sh
look_away
```

Tested only on Mac and Linux.
