# Pwds Server

Making Pwds Password Manager into a web service it needed a server-side script
that accepts requests from the client.



## Getting Started

If you want to host Pwds on your own server, you are absolutely welcome to do
so! 


### Prerequisites

To compile from source, you need to have **Go** installed on your machine! You
can try installing it through your package manager of choice like this:

```bash
apt-get install golang-go
```

Alternatively, you can download [Go binary distributions][bin], go through the [installation process][install], and don't forget to [set the `$GOPATH`
environment variable][GOPATH]!

[bin]: https://golang.org/dl/
[install]: https://golang.org/doc/install
[GOPATH]: https://github.com/golang/go/wiki/SettingGOPATH


### Config & Run

As soon as you have **Go** installed and running on your machine, you can do the
following:

```bash
go get https://github.com/sharpvik/pwds-backend
```

This command will fetch the whole GitHub repo and put it into a specific place
on your computer. For those, who are new to Go programming, I'll give a hack.

```bash
go env GOPATH
# Prints something like /home/username/go on UNIX-based systems.
```

The above command prints out the absolute path to the folder where all
Go-related things are supposed to be stored. I don't know what that path is for 
you, so here I'll just call it `$GOPATH`. Knowing what your `$GOPATH` is, you
can now easily locate the newly installed `pwds-backend` package as follows:

```bash
cd $GOPATH/src/github.com/sharpvik/pwds-backend

# Remember, $GOPATH (here and below) is not actually an environment variable,
# it's just a placeholder for your own Go folder and is supposed to be
# substituted with the path printed by `go env GOPATH`.
```

You can also try ...

```bash
cd $(go env GOPATH)/src/github.com/sharpvik/pwds-backend
```

... instead of copy-pasting the output from `go env GOPATH` by hand.

Once you are in that folder, you'll need to change the `config.go` file a notch.
Edit the `RootFolder` constant in `pwds-backend/config/config.go` so that it
reflects the actual path to the `pwds-backend` folder on your machine.

On my machine it looks like this:

```go
RootFolder = "/home/sharpvik/go/src/github.com/sharpvik/pwds-backend"

// My $GOPATH is set to /home/sharpvik/go so the string in RootFolder
// reflects the exact location of the pwds-backend project folder on my machine.
// You need to change this string to be
//     RootFolder = "$YOUR_GOPATH/src/github.com/sharpvik/pwds-backend"
```

Run the following command from the project's root folder to start your server.

```bash
go run main.go
```

It will immediately start serving at `localhost:8000`. To change the port, stop
the server with `CTRL+C`, edit the `Port` constant in
`pwds-backend/config/config.go`, restart the server.


### Build & Install

Go is actually a compiled language, however the `go run` command doesn't produce
any visible executable files. To compile `pwds-backend`, you can

```bash
cd $GOPATH/src/github.com/sharpvik/pwds-backend

go build # puts binary `pwds-backend` file into the project folder

# or alternatively, use

go install # creates binary file at $GOPATH/bin/pwds-backend
```



## Running the tests

If you want to test some module `X` you can do the following:

```bash
# Assuming you are now in the project's root folder -- `pwds-backend`

cd X
go test

# The `go test` command will print out something like this:
#
#     PASS
#     ok  	github.com/username/pwds-backend/X	0.001s
#
# In case some of the tests fail, output will look like this:
#
#     --- FAIL: TestSomeFunc (0.00s)
#         x_test.go:11: Function `SomeFunc` works incorrectly.
#     FAIL
#     exit status 1
#     FAIL	github.com/username/pwds-backend/X	0.001s
#
```



## Contributing

All contributions are welcome and appreciated! I'd be glad to accept your help
with this project. `CONTRIBUTING.md` will be added it future commits.



## Authors

- **Viktor A. Rozenko Voitenko** - *Initial work* - [sharpvik]

[sharpvik]: https://github.com/sharpvik



## License

This project is licensed under the **Mozilla Public License Version 2.0** --
see the [LICENSE](LICENSE) file for details.

Please note that this project is distributred as is,
**with absolutely no warranty of any kind** to those who are going to deploy
and/or use it. None of the authors and contributors are responsible (liable)
for **any damage**, including but not limited to, loss of sensitive data and
server machine malfunction.



## Acknowledgments

- Hat tip to [Billie Thompson] for the great [README template].
- This project has been greatly improved by all those who commented under
[this Reddit Post][Reddit post] of mine and gave valuable advice.

[Billie Thompson]: https://gist.github.com/PurpleBooth
[README template]: https://gist.github.com/PurpleBooth/109311bb0361f32d87a2

[Reddit post]: https://www.reddit.com/r/Python/comments/egtgjq/password_manager_in_python3/
