package fuzzyhipster

import (
  "bytes"
  "os"
  "io"                               
  "io/ioutil"
  "strings"
  "log"
  //"fmt"
)

// [ ] ADD TIMING
// [ ] ADD Caching of output
// [ ] ADD fmt.StringS to simplify the export


func bundle() (string) {
  
  
  buf := bytes.NewBuffer(nil)
  
  header, _ := os.Open("src/views/header.html")
  io.Copy(buf, header) 
  header.Close()
  
  // get the templates
  //filenames := []string {"src/views/body.html"}
  //for _, filename := range filenames {
  //  f, _ := os.Open(filename)
  //  io.Copy(buf, f) 
  //  f.Close()
 // }
  
  templates, _ := ioutil.ReadDir("src/views/templates")
  for _, template := range templates {
    if(!template.IsDir()) {
      path := "src/views/templates/" + template.Name()
      log.Println(path)
      t, _ := os.Open(path)
      name := strings.Replace(template.Name(), ".html", "", -1)
      io.WriteString(buf, "<script type='text/x-handlebars' data-template-name='")
      io.WriteString(buf,  name)
      io.WriteString(buf,  "'>")
      io.Copy(buf, t) 
      io.WriteString(buf,  "</script>")
      t.Close()
    } else {
      // recursivly search the folders
      bundleFolder("src/views/templates", template.Name(), buf)
    }
  }
  
  // add the component templates
  components, _ := ioutil.ReadDir("src/views/components")
  for _, component := range components {
    if(!component.IsDir()) {
      path := "src/views/components/" + component.Name()
      t, _ := os.Open(path)
      name := strings.Replace(component.Name(), ".html", "", -1)
      io.WriteString(buf, "<script type='text/x-handlebars' data-template-name='components/")
      io.WriteString(buf,  name)
      io.WriteString(buf,  "'>")
      io.Copy(buf, t) 
      io.WriteString(buf,  "</script>")
      t.Close()
    } 
  }
  
  footer, _ := os.Open("src/views/footer.html")
  io.Copy(buf, footer)          
  footer.Close()
  
  s := string(buf.Bytes())
  
  return s
}

func bundleFolder(root string, folder string, buf *bytes.Buffer) error {
	templateFolder := root + "/" + folder
	templates, _ := ioutil.ReadDir(templateFolder)
	 for _, template := range templates {
	 	 if(!template.IsDir()) {
      path := root + "/" + folder + "/"+ template.Name()
      t, _ := os.Open(path)
      name := folder + "/" + strings.Replace(template.Name(), ".html", "", -1)
      io.WriteString(buf, "<script type='text/x-handlebars' data-template-name='")
      io.WriteString(buf,  name)
      io.WriteString(buf,  "'>")
      io.Copy(buf, t) 
      io.WriteString(buf,  "</script>")
      t.Close()
    } else {
      // recursivly search the folders
      	bundleFolder("root", folder + "/" + template.Name(), buf)
    }
	 }
	 return nil
}

func bundleJavascript() (string) {

  buf := bytes.NewBuffer(nil)
    
  bundleJavascriptFolder("src/app", buf)
  
  s := string(buf.Bytes())
  
  return s
}

func bundleJavascriptFolder(root string, buf *bytes.Buffer) error {
	scripts, _ := ioutil.ReadDir(root)
  for _, script := range scripts {
    if(!script.IsDir()) {
      path := root + "/" + script.Name()
      log.Println(path)
      t, _ := os.Open(path)
      io.Copy(buf, t) 
      t.Close()
    } else {
      // recursivly search the folders
      bundleJavascriptFolder(root + "/" +  script.Name(), buf)
    }
  }
  return nil
}

