Language agnostic library API search engine. Inspired by [Hoogle](https://hoogle.haskell.org/). Utilizes [Tree-sitter](https://tree-sitter.github.io/tree-sitter/) for parsing.

> [!Important]
> Work in Progress

## Getting Started

Download [stb_image.h](https://github.com/nothings/stb/blob/master/stb_image.h) for testing.

### GNU/Linux

```shell
mkdir -p dump data
curl -L -o dump/stb_image.h https://raw.githubusercontent.com/nothings/stb/refs/heads/master/stb_image.h
```

### Windows

```powershell
New-Item -ItemType Directory -Force -Path dump, data
curl.exe -L -o dump/stb_image.h https://raw.githubusercontent.com/nothings/stb/master/stb_image.h
```

## Test

```shell
go run .\cmd\cli\main.go "stbi__context *s -> stbi_uc* out"
```

> Searches for all the functions with sigature containing the query.

Output:

```console
stbi__tga_read_rgb16 :: stbi__context *s -> stbi_uc* out -> void
```

## References

1. Go bindings for Tree-sitter: <https://github.com/tree-sitter/go-tree-sitter>
2. Go Tree-sitter documentation: <https://pkg.go.dev/github.com/tree-sitter/go-tree-sitter>
