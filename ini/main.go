package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
func main() {
	/////////////////////////////////////
	// cfg, err := ini.Load("config.ini")
	// if err != nil {
	// 	fmt.Printf("Fail to read file: %v", err)
	// 	os.Exit(1)
	// }
	// fmt.Println("XNBNamePos:", cfg.Section("SETTING").Key("XNBNamePos").String())
	// fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())
	/////////////////////////////////////
	dat, _ := ioutil.ReadFile("config.ini")
	str := strings.NewReader(string(dat))
	tr := transform.NewReader(str, charmap.Windows1251.NewDecoder())
	buf, err := ioutil.ReadAll(tr)
	if err != err {
		// обработка ошибки
	}

	s := string(buf) // строка в UTF-8
	fmt.Print(s)
	s = s + "Просто text"
	err = WriteToFile("config.ini", s)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Print(string(dat))
	// Let's do some candidate value limitation
	// fmt.Println("Server Protocol:",
	// 	cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
	// // Value read that is not in candidates will be discarded and fall back to given default value
	// fmt.Println("Email Protocol:",
	// 	cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))

	// Try out auto-type conversion
	// fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
	// fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))

	// Now, make some changes and save it
	// cfg.Section("").Key("app_mode").SetValue("production")
	// cfg.SaveTo("my.ini.local")
}
