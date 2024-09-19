## GSC

Git Selective Clone (GSC) is a tool that allows you to clone only the folder you need from a git repository. This is useful when you need to clone a specific folder from a repository that contains a lot of folders and files.

## Installation

```bash
go install github.com/BRAVO68WEB/git-selection-clone@latest
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