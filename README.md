![Build](https://github.com/DannyHinshaw/earmuffs/workflows/Build/badge.svg)
![Integration Test](https://github.com/DannyHinshaw/earmuffs/workflows/Integration%20Test/badge.svg)

# Earmuffs - GitHub Action

![oldschool earmuffs](https://i.pinimg.com/originals/05/1c/ff/051cff03d56ea5a34e7c42c4c2fb4eb8.jpg)

## What

A github action to check your source code for profanities. 
If it finds foul language it will fail.

## How

**example.yaml**
```yml
name: Build
on: [push, pull_request]
jobs:
  build:
    name: Build
    runs-on: ubuntu-latest
    steps:
      - name: Set up Go 1.13
        uses: actions/setup-go@v1
        with:
          go-version: 1.13
        id: go

      - uses: actions/checkout@v1

      - name: Run Earmuffs
        uses: dannyhinshaw/earmuffs@master
        with:
          exclude_files: "\\.html" # optional regex, default "\\.idea|\\.git|node_modules"

      - name: Build (or whatever)
        run: |
          go get -d -v
          go build -v .
```

## Credit
Bad words sourced from:

```text
https://www.cs.cmu.edu/~biglou/resources/bad-words.txt
```

