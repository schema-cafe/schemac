# schemac
A CLI app that will download and compile types from schema.cafe into your project

## Install

```bash
go install github.com/schema-cafe/schemac@latest
```

## Usage examples
Write Go types to the directory `pkg/types`:
```bash
schemac go pkg/types
```

Write TypeScript types to the directory `types`:
```bash
schemac ts types
```
