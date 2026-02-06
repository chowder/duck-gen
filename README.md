# 🦆 `duck-gen`

Generate [**DuckDuckGo Email Protection**](https://spreadprivacy.com/protect-your-inbox-with-duckduckgo-email-protection/) addresses from the command-line.

```console
foo@bar:~$ duck-gen
Enter your Duck address: johnny@duck.com
Enter the one-time passphrase sent to your email: heftiness slate captivate cornmeal 
bzpkash3@duck.com

foo@bar:~$ duck-gen
embargo-frisk-rack@duck.com
```

`duck-gen` will initiate a login flow the first time you run it, but will save your authentication tokens at `~/.duck_token` for future runs. 

## Installation 

Download pre-built binaries from the [Releases](https://github.com/chowder/duck-gen/releases) page

### Building from Source

1. **Install Go** (requires Go 1.23.2 or later):
   - Download from https://go.dev/dl/

2. **Clone and build**:
   ```bash
   git clone https://github.com/chowder/duck-gen.git
   cd duck-gen
   go mod vendor
   go build -mod=vendor
   ```

   This will create a `duck-gen` binary in the current directory.

3. **Install to your PATH** (optional):
   ```bash
   go install
   ```
   This installs the binary to `$GOPATH/bin` or `$HOME/go/bin` (add to your PATH if needed).

   > Alternatively, you can move the binary to anywhere you like and put an alias to it in your `.bash_aliases`, `.bashrc`, `.zsh_aliases` or `.zshrc`.