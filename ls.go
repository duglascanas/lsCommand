package main

import "time"

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

//struct
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
	color  string
	symbol string
}

var mapStyleByFileType = map[int]styleFileType{
	fileRegular:    {icon: "📘"},
	fileDirectory:  {icon: "📁", color: "BLUE", symbol: "/"},
	fileExecutable: {icon: "⛳", color: "GREEN", symbol: "*"},
	fileCompress:   {icon: "🗜", color: "RED"},
	fileImage:      {icon: "📷", color: "MAGNETA"},
	fileLink:       {icon: "📎", color: "CYAN"},
}
