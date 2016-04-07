package main

import "gopkg.in/alecthomas/kingpin.v2"

var (
	debug    = kingpin.Flag("debug", "Turn on the debug-mode").Bool()
	userName = kingpin.Flag("username", "Username").Short('u').Required().String()
	password = kingpin.Flag("password", "Password").Short('p').Required().String()
)
