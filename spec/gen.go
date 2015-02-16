package spec

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GenerateApplication(application Application, baseDir string) error {
	err := validate(application)
	if err != nil {
		return err
	}
	err = generateEvents(application, baseDir)
	if err != nil {
		return err
	}
	err = generateServices(application, baseDir)
	if err != nil {
		return err
	}
	return nil
}

func validate(application Application) error {
	// TODO detect collisions with golang
	// TODO detect duplicate names etc
	return nil
}

func generateEvents(application Application, baseDir string) error {
	src := fmt.Sprintf("%s/spec/event.go.tmpl", baseDir)
	target := fmt.Sprintf("%s/%s/events/events.go", baseDir, application.Name)

	err := generateFileFromTemplate(application, src, target)
	if err != nil {
		log.Fatalf("Error generating application-events (%s)", err)
		return err
	}
	return nil
}

func generateServices(application Application, baseDir string) error {
	for _, service := range application.Services {
		src := fmt.Sprintf("%s/spec/service-interface.go.tmpl", baseDir)
		target := fmt.Sprintf("%s/%s/%s/interface.go", baseDir, application.Name, strings.ToLower(service.Name))

		err := serviceGenerateFileFromTemplate(application, service, src, target)
		if err != nil {
			log.Fatalf("Error generating for service-interface %s (%s)", service.Name, err)
			return err
		}
	}
	return nil
}

type ApplicationAndService struct {
	Application Application
	Service     Service
}

func serviceGenerateFileFromTemplate(application Application, service Service, templateFileName string, targetFileName string) error {
	log.Printf("Using template %s to generate service target %s\n", templateFileName, targetFileName)
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(targetFileName), 0777)
	if err != nil {
		return err
	}
	w, err := os.Create(targetFileName)
	if err != nil {
		return err
	}
	defer w.Close()
	as := ApplicationAndService{Application: application, Service: service}
	if err := t.Execute(w, as); err != nil {
		return err
	}
	return nil
}
func generateFileFromTemplate(data interface{}, templateFileName string, targetFileName string) error {
	log.Printf("Using template %s to generate target %s\n", templateFileName, targetFileName)
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(targetFileName), 0777)
	if err != nil {
		return err
	}
	w, err := os.Create(targetFileName)
	if err != nil {
		return err
	}
	defer w.Close()
	if err := t.Execute(w, data); err != nil {
		return err
	}
	return nil
}
