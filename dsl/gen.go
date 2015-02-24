package spec

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GenerateApplication(application Application, baseDir string) error {
	err := ValidateApplication(application)
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
	err = generateDoumentation(application, baseDir)
	if err != nil {
		return err
	}
	return nil
}

func validateDuplicateEntries(comment string, entries []string) error {
	namesMap := make(map[string]int)
	for _, entry := range entries {
		count, exists := namesMap[entry]
		if exists == false {
			namesMap[entry] = 1
		} else {
			namesMap[entry] = count + 1
		}
	}
	nameList := make([]string, 0, 10)
	for name, counter := range namesMap {
		if counter > 1 {
			nameList = append(nameList, name)
		}
	}
	if len(nameList) > 0 {
		return errors.New(fmt.Sprintf("duplicate %s: %v", comment, nameList))
	}

	return nil
}

func validateNameLength(comment string, entries []string) error {
	for _, entry := range entries {
		if len(entry) <= 1 {
			return errors.New(fmt.Sprintf("%s name too short: %s", comment, entry))
		}
	}
	return nil
}

func ValidateApplication(application Application) error {
	{
		err := validateNameLength("service-name", application.ServiceNames())
		if err != nil {
			return err
		}
		err = validateDuplicateEntries("service-names", application.ServiceNames())
		if err != nil {
			return err
		}
	}
	{
		for _, event := range application.GetUniqueEvents() {
			err := validateNameLength("event-attribute-name", event.AttributeNames())
			if err != nil {
				return err
			}

			err = validateDuplicateEntries("event-attribute-names", event.AttributeNames())
			if err != nil {
				return err
			}
		}
	}
	{
		for _, service := range application.Services {
			err := validateNameLength("command-name", service.CommandNames())
			if err != nil {
				return err
			}
			err = validateDuplicateEntries("command-names", service.CommandNames())
			if err != nil {
				return err
			}
			for _, command := range service.Commands {
				err := validateNameLength("input-attribute-name", command.Input.AttributeNames())
				if err != nil {
					return err
				}
				err = validateDuplicateEntries("input-attribute-names", command.Input.AttributeNames())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func generateEvents(application Application, baseDir string) error {
	src := fmt.Sprintf("%s/dsl/event.go.tmpl", baseDir)
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
		src := fmt.Sprintf("%s/dsl/service-interface.go.tmpl", baseDir)
		target := fmt.Sprintf("%s/%s/%s/interface.go", baseDir, application.NameToFirstLower(), strings.ToLower(service.Name))

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
	//w = template.New("consumingServices").Funcs(defaultfuncs)

	defer w.Close()
	if err := t.Execute(w, data); err != nil {
		return err
	}
	return nil
}

func generateDoumentation(application Application, baseDir string) error {
	src := fmt.Sprintf("%s/dsl/graphviz.dot.tmpl", baseDir)
	target := fmt.Sprintf("%s/%s/doc/graphviz.dot", baseDir, application.Name)

	err := generateFileFromTemplate(application, src, target)
	if err != nil {
		log.Fatalf("Error generating graphviz-doc (%s)", err)
		return err
	}
	return nil
}

var defaultfuncs = map[string]interface{}{
	"consumingServices": func(event Event) []string {
		return []string{"eva", "marc", "pien"}
	},
}
