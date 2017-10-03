package main

import (
    std "fmt"
    "io/ioutil"
    "path/filepath"
    "strings"
    "os"
    "os/exec"
    "gopkg.in/yaml.v2"
)

type Config struct {
    Orders map[string]Order
}

type Order struct {
  Alias string `yaml:"alias"`
  Command []string `yaml:"command"`
  Variable string `yaml:"variable"`
  Value string `yaml:"value"`
}

func main() {
    filename, _ := filepath.Abs("./dobby.yml")
    yamlFile, err := ioutil.ReadFile(filename)

    if err != nil {
        panic(err)
    }

    var config Config

    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        panic(err)
    }

    var order = config.Orders[strings.Join(os.Args[1:]," ")]

    //resolve alias if any
    for ;order.Alias != ""; {
      order = config.Orders[order.Alias]
    }

    if order.Command != nil {
      out, _ := exec.Command(order.Command[0], order.Command[1:]...).Output()
      std.Printf(string(out));
    }
}
