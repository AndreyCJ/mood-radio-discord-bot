package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

var (
	Token         string
	CommandPrefix string
	config        *configStruct
)

type configStruct struct {
	CommandPrefix string
}

func init() {
	// set Token value from BOT_TOKEN env variable
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json") // ioutil package's ReadFile method which we read config.json and return it's value we will then store it in file variable and if an error ocurrs it will be stored in err .

	//Handling error and printing it using fmt package's Println function and returning it .
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// We are here printing value of file variable by explicitly converting it to string .
	// fmt.Println(string(file))

	// Here we performing a simple task by copying value of file into config variable which we have declared above , and if there any error we are storing it in err . Unmarshal takes second arguments reference remember it .
	err = json.Unmarshal(file, &config)

	//Handling error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// After storing value in config variable we will access it and storing it in our declared variables .
	CommandPrefix = config.CommandPrefix

	return nil
}
