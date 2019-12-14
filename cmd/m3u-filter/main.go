package main

import (
	"flag"
	"github.com/mitchellh/go-homedir"
	"github.com/otaconix/m3ufilter/config"
	"github.com/otaconix/m3ufilter/logger"
	"github.com/otaconix/m3ufilter/m3u"
	"github.com/otaconix/m3ufilter/server"
	"github.com/otaconix/m3ufilter/writer"
	"os"
)

func main() {
	configFile := flag.String("config", "~/.m3u.conf", "Config file location")
	playlistOutput := flag.String("playlist", "", "Where to output the playlist data. Ignored when using -server flag. Defaults to stdout")
	logOutput := flag.String("log", "", "Where to output logs. Defaults to stderr")
	versionFlag := flag.Bool("version", false, "show version and exit")
	flag.Parse()

	if *versionFlag {
		config.ShowVersion()
	}

	path, e := homedir.Expand(*configFile)
	if e != nil {
		panic(e)
	}

	run(path, fd(*playlistOutput, false), fd(*logOutput, true))
}

func run(configFilename string, stdout *os.File, stderr *os.File) {
	log := logger.Get()
	log.SetOutput(stderr)

	conf := config.New(configFilename)
	if conf.Core.ServerListen != "" {
		server.Serve(conf)
	} else {
		playlist, _ := m3u.GetPlaylist(conf)
		writer.WriteOutput(conf.Core.Output, stdout, playlist)
	}
}

func fd(filename string, defaultStderr bool) *os.File {
	if filename == "-" {
		return os.Stdout
	}
	if filename == "" && defaultStderr {
		return os.Stderr
	}
	if filename == "" {
		return os.Stdout
	}

	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return fd
}
