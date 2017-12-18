package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/hunterhug/parrot/util"
	"github.com/hunterhug/parrot/util/gomail"
)

const (
	Config  = "config.txt"
	Emails  = "email.txt"
	Content = "subject.txt"
)

var Symol = "\n"

func main() {
	fmt.Println("See https://github.com/hunterhug/pigeon\n")
	osname := runtime.GOOS
	if strings.Contains(osname, "window") {
		Symol = "\r\n"
	}

	if strings.Contains(osname, "darwin") {
		Symol = "\r"
	}

	rawconfig, err := util.ReadfromFile(Config)
	exiterror(err)
	config := cutbyrow(rawconfig)
	if len(config) < 4 {
		exiterror(errors.New("config.txt error"))
	}
	fmt.Printf("Config: %v\n", config)
	fmt.Println("-----Content--------")

	rawemails, err := util.ReadfromFile(Emails)
	exiterror(err)
	emails := cutbyrow(rawemails)
	if len(emails) == 0 {
		exiterror(errors.New("email.txt empty"))
	}

	rawcontent, err := util.ReadfromFile(Content)
	exiterror(err)
	ctx := strings.Split(string(rawcontent), Symol+Symol)
	if len(ctx) < 2 {
		exiterror(errors.New("subject.txt wrong format"))
	}
	subject := ctx[0]
	paper := strings.Join(ctx[1:len(ctx)], Symol)
	fmt.Printf("%v\n", subject)
	fmt.Printf("%v\n", paper)

	auth := gomail.MailAuth{}
	auth.Host = config[0]
	auth.Port, _ = util.SI(config[1])
	auth.UserName = config[2]
	auth.Password = config[3]

	mconfig := gomail.MailConfig{}
	mconfig.From = auth.UserName
	mconfig.Subject = subject
	mconfig.Body = []byte(paper)
	mconfig.BodyType = ""

	bccs := []gomail.MailAddress{{"459527502@qq.com", ""}}
	for _, email := range emails {
		bccs = append(bccs, gomail.MailAddress{email, ""})
	}

	mconfig.Bcc = bccs

	files, err := util.WalkDir("attach", "")
	exiterror(err)

	if len(files) > 0 {
		mconfig.Attach = files
		fmt.Println("-----Attach--------")
		fmt.Printf("prepare attach:%v\n", files)
		fmt.Println("-----Sending--------")
	}

	err = gomail.SendMail(auth, mconfig)
	if err != nil {
		fmt.Println("Error:" + err.Error())
		fmt.Println("Your Email config error")
	} else {
		fmt.Printf("Send email: %#v\n", emails)
		fmt.Println("-----Success Send----")
	}

	fmt.Println("----Please Close it-----")
	util.Sleep(10)
}

func cutbyrow(raw []byte) []string {
	// Linux Unix \n
	// Windows \r\n
	// Mac \r
	temp := strings.Split(string(raw), Symol)
	change := []string{}
	for _, i := range temp {
		xx := strings.TrimSpace(i)
		if xx == "" {
			continue
		}
		change = append(change, xx)
	}
	return change
}

func exiterror(err error) {
	if err != nil {
		fmt.Println(err.Error())
		util.Sleep(10)
		os.Exit(1)
	}
	return
}
