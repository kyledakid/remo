package main

import "fmt"

var (
	version    = "0.0"
	branch     = "master"
	build_date = "now"
	commit_id  = "000"
	tag        = "0.0"
)

func main() {
	intro := "Welcome to Remo :)\n\n" +
		"Remo is an app to demo Github releases\n" +
		"Author: Kyle\n" +
		"Date: June 2020\n" +
		"Email: nvietthu@gmail.com\n" +
		"Github: github.com/kyledakid\n"
	rel_info := "Release Info: \n" +
		"Version: " + version + "\n" +
		"Branch: " + branch + "\n" +
		"Build Date: " + build_date + "\n" +
		"Commit ID: " + commit_id + "\n" +
		"Tag: " + tag + "\n"

	fmt.Println(intro)
	fmt.Println(rel_info)
}
