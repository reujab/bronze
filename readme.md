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
Install the bronze binary by compiling it from source or download a pre-compiled binary on the [releases page](https://github.com/reujab/bronze/releases) and add it to your `PATH`.

### Configuration
Now that you have bronze installed, you need to configure it. To have your prompt look like the one in the screenshot above, add this to your `~/.bashrc`/`~/.zshrc`:
```sh
BRONZE=(status:black:white dir:blue:black git:green:black cmdtime:magenta:black)
```

Or add the following to your `~/.config/fish/config.fish`:
```fish
set BRONZE status:black:white dir:blue:black git:green:black cmdtime:magenta:black
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
