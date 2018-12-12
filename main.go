package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	emailsAndLanguages := os.Args[1:]
	emails, languages := separateEmailsAndLanguages(emailsAndLanguages)
	fmt.Println(emails, languages)
	newsletterHTML := newsletter(languages)

	for _, email := range emails {
		sendNewsletter(newsletterHTML, email)
	}
}

func sendNewsletter(newsletter string, email string) {

}

func separateEmailsAndLanguages(emailsAndLanguages []string) ([]string, []string) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	var emails []string
	var languages []string
	for _, emailOrLanguage := range emailsAndLanguages {
		match := re.MatchString(emailOrLanguage)

		if match {
			emails = append(emails, emailOrLanguage)
		} else {
			languages = append(languages, emailOrLanguage)
		}
	}

	return emails, languages
}
