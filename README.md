# 🦆 `duck-gen`

Generate [**DuckDuckGo Email Protection**](https://spreadprivacy.com/protect-your-inbox-with-duckduckgo-email-protection/) addresses from the command-line.

```console
foo@bar:~$ duck-gen
Enter your Duck address: johnny@duck.com
Enter the one-time passphrase sent to your email: heftiness slate captivate cornmeal 
bzpkash3@duck.com

foo@bar:~$ duck-gen
b2lo513ds@duck.com 
```

`duck-gen` will initiate a login flow the first time you run it, but will save your authentication tokens at `~/.duck_token` for future runs. 

## Installation 

Download pre-built binaries from the [Releases](https://github.com/chowder/duck-gen/releases) page
