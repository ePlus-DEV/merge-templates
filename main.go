package main

import (
	"fmt"
	"github.com/huantt/merge-templates/handler"
	"log/slog"
	"os"
)

func main() {
	var loggingLevel = new(slog.LevelVar)
	loggingLevel.Set(slog.LevelInfo)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: loggingLevel}))
	slog.SetDefault(logger)

	args := os.Args
	if len(args) < 4 {
		slog.Error("Missing args. Usage: go run main.go $OUT_PATH $MAIN_TEMPLATE [<template_path_1> <template_path_2>...]")
		os.Exit(1)
	}

	outPath := args[1]
	mainTemplatePath := args[2]
	templatePaths := args[3:]
	slog.Info(fmt.Sprintf("MergeTemplates: outPath=%s - mainTemplatePath=%s - templatePaths=%v", outPath, mainTemplatePath, templatePaths))
	err := handler.MergeTemplates(outPath, mainTemplatePath, templatePaths...)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("Done")
}
