package handler

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

type TemplateContext struct {
	Error   string
	Data    any
	Message string
}

var context *TemplateContext = &TemplateContext{
	Error:   "",
	Data:    nil,
	Message: "",
}

func clearContext() {
	context.Error = ""
	context.Message = ""
}

func render(w http.ResponseWriter, templateName string) {
	t, err := template.ParseFiles("template/base.html", fmt.Sprintf("template/%s.html", templateName))
	check(err)
	err = t.Execute(w, context)
	check(err)
}

func renderError(err error, w http.ResponseWriter, templateName string) {
	context.Error = err.Error()
	render(w, templateName)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func validatePhoneNumber(phoneNumber string) error {
	match, err := regexp.Match(`^\+[1-9]\d{1,14}$`, []byte(phoneNumber))

	if err != nil {
		return err
	}
	if !match {
		return errors.New("phone number must be in E.164 format")
	}
	return nil
}
