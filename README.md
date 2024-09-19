## GSC

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/BRAVO68WEB/git-selective-clone?style=for-the-badge&logo=Go)
![GitHub](https://img.shields.io/github/license/BRAVO68WEB/git-selective-clone?style=for-the-badge)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/bravo68web/git-selective-clone/build.yaml?style=for-the-badge&logo=Github)


Git Selective Clone (GSC) is a tool that allows you to clone only the folder you need from a git repository. This is useful when you need to clone a specific folder from a repository that contains a lot of folders and files.

## Installation

```bash
go install github.com/BRAVO68WEB/git-selection-clone@latest
mv $GOPATH/bin/git-selection-clone $GOPATH/bin/gsc
```

## Usage

```bash
gsc <repository-url> <folder-path>
```

## Example

```bash
gsc https://github.com/BRAVO68WEB/pacsearch/tree/main/apps/pacsearch-app/components pacsearch_components
```

## License

[MIT](./LICENSE)