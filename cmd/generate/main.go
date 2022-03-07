package main

import (
	"os"
	"os/exec"
	"errors"
	"strings"
    "html/template"
    "github.com/nlittlepoole/horcrux"
    b64 "encoding/base64"
)

func MinifyFile(inputFilename, outputFilename string) (err error) {
	cmd := exec.Command("minify", []string{"-o", outputFilename, inputFilename}...)
	var cmdOut []byte
	if cmdOut, err = cmd.CombinedOutput(); err != nil {
		err = errors.New(strings.Join([]string{err.Error(), string(cmdOut)}, ", "))
	}
	return
}

func Base64File(inputFilename, outputFilename string) (err error) {
	var dat []byte
	if dat, err = os.ReadFile(inputFilename); err == nil {
		sEnc := b64.StdEncoding.EncodeToString(dat)
		err = os.WriteFile(outputFilename, []byte("data:text/html;base64," + sEnc), 0644)
	}
	return
}

func main() {
	inname := os.Args[1]
	parsedTemplate, _ := template.ParseFiles(inname)
	outname := strings.ReplaceAll(strings.ReplaceAll(inname, "gohtml", "html"), "templates", "build")
	if f, err := os.Create(outname); err == nil {
	    if err = parsedTemplate.Execute(f, horcrux.GetEnvironmentHorcrux()); err == nil {
	    	f.Close()
	    	minname := strings.ReplaceAll(outname, ".html", "-minified.html")
	    	if err = MinifyFile(outname, minname); err == nil {
	    		base64Name := strings.ReplaceAll(outname, ".html", ".txt")
	    		err = Base64File(minname, base64Name)
	    	}
	    } else {
	    	f.Close()
	    }
		if err != nil {
		    panic(err)
		}
	} else {
		panic(err)
	}
}