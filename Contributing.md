
# Easy Distribute and hacking.

Want to change just a couple lines?

The only dependency you need is `podman
- [podman https://podman.io/docs/installation](https://podman.io/docs/installation) On ubuntu just use`sudo apt install podman`

and run the distribute script
```sh
./distribute.sh
```
That will use a podman container to build the entire app and it will put the output
in `./dist`.


# Development

Below are all the dependencies this app needs. but
you can also look at `./resources/Containerfile` to see how to install all
the dependencies

## Deps:

- Download the following dependecies from your system's package manager. On ubuntu use: `sudo apt install pkg-config libchafa-dev`
- Optional: [vscode](https://code.visualstudio.com/) with these recommended extensions:
    - "mesonbuild.mesonbuild",
    - "ms-vscode.cpptools-extension-pack",
    - "golang.go",

### Version map
These are the versions of the tools used to build and run the project:
- chafa 1.18.0



# Running and building


### Most useful tasks

## run

```sh
go run . firefox
```


## clean-all
Remove all build artifacts.
```sh
make clean
```
e, good for local testing or sending to friends

## distribute
Creates an AppImage in a Ubuntu 22.04 podman container. AppImages mostly 
forward-compatible, but are not back-wards compatible, so make them in this
container for compatibility.
