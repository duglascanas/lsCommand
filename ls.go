package main

import (
	"time"

	"github.com/fatih/color"
)

// windows os system
const Windows = "windows"

// file type
const (
	fileRegular int = iota
	fileDirectory
	fileExecutable
	fileCompress
	fileImage
	fileLink
)

// file extension
const (
	exe = ".exe"
	deb = ".deb"
	zip = ".zip"
	gz  = ".gz"
	tar = ".tar"
	rar = ".rar"
	png = ".png"
	jpg = ".jpg"
	gif = ".gif"
)

// struct
type file struct {
	name             string
	fileType         int
	isDir            bool
	isHidden         bool
	userName         string
	groupName        string
	size             int64
	modificationTime time.Time
	mode             string
}

type styleFileType struct {
	icon   string
	color  color.Attribute
	symbol string
}

var mapStyleByFileType = map[int]styleFileType{
	fileRegular:    {icon: "📘"},
	fileDirectory:  {icon: "📁", color: color.FgBlue, symbol: "/"},
	fileExecutable: {icon: "⛳", color: color.FgGreen, symbol: "*"},
	fileCompress:   {icon: "🗜", color: color.FgRed},
	fileImage:      {icon: "📷", color: color.FgMagenta},
	fileLink:       {icon: "📎", color: color.FgCyan},
}

var (
	blue    = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
	green   = color.New(color.FgGreen).Add(color.Bold).SprintfFunc()
	red     = color.New(color.FgRed).Add(color.Bold).SprintfFunc()
	magenta = color.New(color.FgMagenta).Add(color.Bold).SprintfFunc()
	cyan    = color.New(color.FgCyan).Add(color.Bold).SprintfFunc()
	yellow  = color.New(color.FgYellow).SprintfFunc()
)
