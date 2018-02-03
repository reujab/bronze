# bronze
A cross-shell customizable powerline-like prompt heavily inspired by [Agnoster](https://github.com/agnoster/agnoster-zsh-theme).<br/>
![](./sleep.png)

## How does it work?
Unlike pretty much every other shell prompt, bronze is not written in shell script, but entirely in Go, so all prompt segments are loaded asynchronously for a speed boost.

When `bronze init` is run, it outputs shell code that sets your prompt to run `bronze prompt`, which outputs the actual prompt. The `bronze prompt` command relies on environment variables for configuration.

## Getting started
Since bronze is not written in shell script, it should theoretically be compatible with any shell, but the three supported shells are Bash, Zsh, and fish.

### Icons
To be able to use the custom icons (which are enabled by default), you must patch your font or install a pre-patched font from [Nerd Fonts](https://github.com/ryanoasis/nerd-fonts).

### Installation
#### From source
* install and setup [Go](https://golang.org/)
* install [libgit2](https://libgit2.github.com/) development packages
* run `go get github.com/reujab/bronze`

#### From pre-compiled binary
* download a binary on the [releases page](https://github.com/reujab/bronze/releases)
* add binary to `PATH` environment variable

#### macOS
On macOS, you will have to do a bit more:
* install [Homebrew](https://brew.sh/)
* run `brew install coreutils`
* add `alias date="gdate"` to your shell rc

### Configuration
Now that you have bronze installed, you need to configure it. To have your prompt look like the one in the screenshot above, add this to your `~/.bashrc`/`~/.zshrc`:
```sh
BRONZE=(status:black:white dir:blue:black git:green:black cmdtime:magenta:black)
export BRONZE_SHELL=$SHELL # bash, zsh, or fish
```

Or add the following to your `~/.config/fish/config.fish`:
```fish
set BRONZE status:black:white dir:blue:black git:green:black cmdtime:magenta:black
set -x BRONZE_SHELL fish
```

Now that bronze is configured, you need to evaluate its bootstrap code.

`~/.bashrc`/`~/.zshrc`:
```sh
eval "$(bronze init)"
```

`~/.config/fish/config.fish`:
```fish
eval (bronze init)
```

## Documentation
Documentation is available on [the wiki](https://github.com/reujab/bronze/wiki).

## Project structure
* [`packagesd/`](packagesd)
	* `main.go`
		* source code for the [`packages`](https://github.com/reujab/bronze/wiki/Packages) module daemon
	* `packagesd.service`
		* a [systemd service file](https://www.freedesktop.org/software/systemd/man/systemd.service.html)
		* can be placed in `/usr/lib/systemd/system/`
		* can be enabled with `systemctl enable packagesd`
		* can be started with `systemctl start packagesd`
* [`static/`](static)
	* `ab0x.go`
		* a file automatically generated with `go generate`
		* contains the shell script files that are printed when running [`bronze init`](https://github.com/reujab/bronze/wiki#init)
* [`types/`](types)
	* `main.go`
		* exports types that are useful when create a [custom segment](https://github.com/reujab/bronze/wiki/Plugin)
* `cmdtime.go`
	* source code for the [`cmdtime`](https://github.com/reujab/bronze/wiki/Command-Time) module
* `dir.go`
	* source code for the [`dir`](https://github.com/reujab/bronze/wiki/Directory) module
* `env.go`
	* source code for the [`env`](https://github.com/reujab/bronze/wiki/Environment-Variable) module
* `git.go`
	* source code for the [`git`](https://github.com/reujab/bronze/wiki/Git) module
* `icons.go`
	* initializes every icon based on [`BRONZE_ICONS`](https://github.com/reujab/bronze/wiki#bronze_icons) and other environment variables
* `init.bash`
	* bootstraps prompt for Bash
* `init.fish`
	* bootstraps prompt for fish
* `init.go`
	* source code for the [`init`](https://github.com/reujab/bronze/wiki#init) subcommand
* `init.zsh`
	* bootstraps prompt for Zsh
* `main.go`
	* parses command line arguments
* `modules.go`
	* contains logic that determines which module function to call
* `os.go`
	* source code for the [`os`](https://github.com/reujab/bronze/wiki/OS) module
* `packages.go`
	* source code for the [`packages`](https://github.com/reujab/bronze/wiki/Packages) module
* `plugin.go`
	* source code for the [`plugin`](https://github.com/reujab/bronze/wiki/Plugin) module
	* invokes code from dynamic libraries (`.so` files)
* `print.go`
	* source code for the [`print`](https://github.com/reujab/bronze/wiki#print) subcommand
* `rss.go`
	* source code for the [`rss`](https://github.com/reujab/bronze/wiki/RSS) module
* `sh.go`
	* contains shell-specific code
* `status.go`
	* source code for the [`status`](https://github.com/reujab/bronze/wiki/Status) module
* `time.go`
	* source code for the [`time`](https://github.com/reujab/bronze/wiki/Time) module
* `user.go`
	* source code for the [`user`](https://github.com/reujab/bronze/wiki/User) module

## Similar projects
[powerline-go](https://github.com/justjanne/powerline-go)
