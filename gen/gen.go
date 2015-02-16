package gen

import (
	"fmt"
	"github.com/xebia/microgen/spec"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GenerateApplication(application spec.Application, baseDir string) error {
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

func validate(application spec.Application) error {
	// TODO detect collisions with golang
	// TODO detect duplicate names etc
	return nil
}

func generateEvents(application spec.Application, baseDir string) error {
	src := fmt.Sprintf("%s/gen/event.go.tmpl", baseDir)
	target := fmt.Sprintf("%s/%s/events/events.go", baseDir, application.Name)

	err := generateFileFromTemplate(application, src, target)
	if err != nil {
		log.Fatalf("Error generating for events (%s)", err)
		return err
	}
	return nil
}

func generateServices(application spec.Application, baseDir string) error {
	for _, service := range application.Services {
		src := fmt.Sprintf("%s/gen/service-interface.go.tmpl", baseDir)
		target := fmt.Sprintf("%s/%s/%s/interface.go", baseDir, application.Name, strings.ToLower(service.Name))

		err := generateFileFromTemplate(service, src, target)
		if err != nil {
			log.Fatalf("Error generating for service-commands %s (%s)", service.Name, err)
			return err
		}
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
