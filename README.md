# GPM - Games Package Manager
Main repository for the official Games Package Manager CLI's code.

## Summary:
- [[0]](#summary) the summary
- [[1]](#the-repository) the repository
- [[2]](#the-cli) the CLI
- [[3]](#choices-made-for-the-cli) the choices we made for the CLI

## The repository
The CLI is written in [rust](https://www.rust-lang.org/fr) and is separated in 3 projects:
- [`./gpm_cli`](./gpm_cli), it takes user inputs and redirects them to the GPM core library
- [`./gpm_core`](./gpm_core), for all the logic of GPM. Everything done by GPM is achieved by the code in this project
- [`./gpm_vfs`](./gpm_vfs), for the system that dynamically installs your mod at runtime


## The CLI
> There is an ongoing draft about the cli [here](https://github.com/WolvenKit/CP77Wiki/wiki/GPM-(Draft)), it probably move in this repository in the future

## Choices made for the CLI
### Why in Rust
Most of the developers who participated in the project has the most experience in Rust, and also wanted to use Rust in this project because they liked the language.

### Why VFS, and what is VFS
> Anyone with more knowledge should probably explain why VFS (and not symlinks)

### Why a CLI and not a GUI
An independant CLI means anyone can make a GUI on top of this CLI and use the full power of it. We opted for a CLI because it is easier for us, developers to use CLIs than GUI. And as said above a CLI means anyone can build a tool on top of it.
