package main

// DO NOT EDIT, this file is generated by hover at compile-time for the fast_rsa plugin.

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	fast_rsa "github.com/jerson/flutter-rsa/go"
)

func init() {
	// Only the init function can be tweaked by plugin maker.
	options = append(options, flutter.AddPlugin(&fast_rsa.Plugin{}))
}