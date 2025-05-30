package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"stats/internal/apiserver"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

var configPath = "config/config.toml"

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	//fmt.Println(config)
	if err != nil {
		log.Fatal(err)
	}
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}

}
