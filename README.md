<p align="center">
<img style="align: center; padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="238px" height="238px" src="pacman.png" />
</p>

<h2 align="center">Go wrapper for arch package manager</h2>

![Generic badge](https://img.shields.io/badge/status-alpha-red.svg)

This library aims to provide concurrent, stable and extensible interface to interact with arch package manager - pacman.

Some of the default options for functions can contain sudo, if you don't need it you can manually disable it.

Functions:

- `Sync` - syncronize packages

```go
import "fmnx.su/core/pacman"

func main() {
	err := pacman.Sync("nano")
    fmt.Println(err)
}
```

- `Search` - search for packages in pacman databases

```go
import "fmnx.su/core/pacman"

func main() {
	r, err := pacman.Search("vim")
    fmt.Println(r)
    fmt.Println(err)
}
```

- `Upgrade` - install packages from local files

```go
import "fmnx.su/core/pacman"

func main() {
	err := pacman.Upgrade("pkg-1-1-any.pkg.tar.zst")
    fmt.Println(err)
}
```

- `Query` - list installed packages

```go
import "fmnx.su/core/pacman"

func main() {
	r, err := pacman.Query()
    fmt.Println(r)
    fmt.Println(err)
}
```

- `Makepkg` - build package

```go
import "fmnx.su/core/pacman"

func main() {
    err := pacman.Makepkg()
    fmt.Println(err)
}
```

- `Remove` - remove installed packages

```go
import "fmnx.su/core/pacman"

func main() {
    err := pacman.Remove("emacs")
    fmt.Println(err)
}
```
