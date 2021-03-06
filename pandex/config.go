package main

import (
	"flag"
)

type Config struct {
	Update   bool
	LogLevel string
	CacheDir string
	ExtDir   string

	InputIndex  string
	OutputIndex string

	Flatten bool
}

var config *Config

func configure() {
	update := false
	flag.BoolVar(&update, "update", false, "Update the cached packages file with new packages")

	loglevel := "INFO"
	flag.StringVar(&loglevel, "loglevel", "INFO", "Log level")

	cachedir := ".gopancache"
	flag.StringVar(&cachedir, "cachedir", ".gopancache", "GoPAN cache directory")

	extdir := ".gopancache"
	flag.StringVar(&extdir, "extdir", ".gopancache", "Temporary directory for extraction")

	inputidx := "index"
	flag.StringVar(&inputidx, "input", "index", "Input index file")

	outputidx := "packages"
	flag.StringVar(&outputidx, "output", "packages", "Output index file")

	flatten := false
	flag.BoolVar(&flatten, "flatten", false, "Only flatten index, don't update anything")

	flag.Parse()

	config = &Config{
		Update:      update,
		LogLevel:    loglevel,
		CacheDir:    cachedir,
		ExtDir:      extdir,
		InputIndex:  inputidx,
		OutputIndex: outputidx,
		Flatten:     flatten,
	}
}
