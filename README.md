## About
This command is used to merge Go template files into one.

## Install
```shell
go install github.com/ePlus-DEV/merge-templates@latest
```

## Usage
```shell
merge-templates $OUT_PATH $MAIN_TEMPLATE [<template_path_1> <template_path_2>...]
```

**Sample**
```shell
merge-templates out.txt main_template.tpl template_1.tpl template_2.tpl
```