package common

import (
    "encoding/json"
    "os"
    "log"
)

type Configuration struct {
    TelegramAPI string
}

func LoadConfiguration() *Configuration {
    file, _ := os.Open("config/config.json")
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        log.Printf("error:", err)
    }
    log.Printf("the key is %s", configuration.TelegramAPI)
    return &configuration
}
