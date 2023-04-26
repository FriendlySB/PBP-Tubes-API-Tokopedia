package controller

import (
	"PBP-Tubes-API-Tokopedia/model"
	"bytes"
	"fmt"
	"text/template"

	gm "gopkg.in/gomail.v2"
)

func sendMailRegis(user model.User) {
	mail := gm.NewMessage()

	template := "bin/template/mailRegis.html"

	result, _ := parseTemplate(template, user)

	mail.SetHeader("From", "lamabunta@gmail.com")
	mail.SetHeader("To", user.Email)
	mail.SetHeader("Subject", "Notifications")
	mail.SetBody("text/html", result)

	sender := gm.NewDialer("smtp.gmail.com", 25, "lamabunta@gmail.com", "gnkglansnfmbshty")

	if err := sender.DialAndSend(mail); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Email sent to: ", user.Email)
	}
}

func sendMailLogin(user model.User) {
	mail := gm.NewMessage()

	template := "bin/template/mailLogin.html"

	result, _ := parseTemplate(template, user)

	mail.SetHeader("From", "lamabunta@gmail.com")
	mail.SetHeader("To", user.Email)
	mail.SetHeader("Subject", "Notifications")
	mail.SetBody("text/html", result)

	sender := gm.NewDialer("smtp.gmail.com", 25, "lamabunta@gmail.com", "gnkglansnfmbshty")

	if err := sender.DialAndSend(mail); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Email sent to: ", user.Email)
	}
}

func sendMailBanUser(user model.User) {
	mail := gm.NewMessage()

	template := "bin/template/mailBanUser.html"

	result, _ := parseTemplate(template, user)

	mail.SetHeader("From", "lamabunta@gmail.com")
	mail.SetHeader("To", user.Email)
	mail.SetHeader("Subject", "Notifications")
	mail.SetBody("text/html", result)

	sender := gm.NewDialer("smtp.gmail.com", 25, "lamabunta@gmail.com", "gnkglansnfmbshty")

	if err := sender.DialAndSend(mail); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Email sent to: ", user.Email)
	}
}

func sendMailRegisShop(user model.User, shop string) {
	// Fungsi untuk mengirim email
	sendEmail := func(to, template string) {
		mail := gm.NewMessage()
		result, _ := parseTemplate(template, user)
		mail.SetHeader("From", "lamabunta@gmail.com")
		mail.SetHeader("To", to)
		mail.SetHeader("Subject", "Notifications")
		mail.SetBody("text/html", result)
		sender := gm.NewDialer("smtp.gmail.com", 25, "lamabunta@gmail.com", "gnkglansnfmbshty")
		if err := sender.DialAndSend(mail); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Email sent to: ", to)
		}
	}

	template := "bin/template/mailRegisShop.html"
	// Mengirim email pertama ke user
	sendEmail(user.Email, template)

	// Mengirim email kedua ke shop
	sendEmail(shop, template)
}

func sendMailInsertAdmin(user model.User) {
	mail := gm.NewMessage()

	template := "bin/template/mailInsertAdmin.html"

	result, _ := parseTemplate(template, user)

	mail.SetHeader("From", "lamabunta@gmail.com")
	mail.SetHeader("To", user.Email)
	mail.SetHeader("Subject", "Notifications")
	mail.SetBody("text/html", result)

	sender := gm.NewDialer("smtp.gmail.com", 25, "lamabunta@gmail.com", "gnkglansnfmbshty")

	if err := sender.DialAndSend(mail); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Email sent to: ", user.Email)
	}
}

func sendMailBanShop(shop model.Shop) {
	mail := gm.NewMessage()

	template := "bin/template/mailInsertAdmin.html"

	result, _ := parseTemplate(template, shop)

	mail.SetHeader("From", "lamabunta@gmail.com")
	mail.SetHeader("To", shop.Email)
	mail.SetHeader("Subject", "Notifications")
	mail.SetBody("text/html", result)

	sender := gm.NewDialer("smtp.gmail.com", 25, "lamabunta@gmail.com", "gnkglansnfmbshty")

	if err := sender.DialAndSend(mail); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Email sent to: ", shop.Email)
	}
}

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}

	var buff bytes.Buffer
	if err := t.Execute(&buff, data); err != nil {
		return "", err
	}

	return buff.String(), nil
}
