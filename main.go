package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/9d4/bzn-mqttserver/config"

	mqtt "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/listeners"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

var (
	EnvConfig *viper.Viper

	MQTTServer *mqtt.Server
)

func init() {
	EnvConfig = viper.New()
	EnvConfig.AddConfigPath(".")
	EnvConfig.SetConfigName(".env")
	EnvConfig.SetConfigType("env")
	EnvConfig.AutomaticEnv()

	// load ENV
	if err := EnvConfig.ReadInConfig(); err != nil {
		jww.ERROR.Fatal(err)
	}

	config.LoadConfig(".")
}

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	fmt.Println("MQTT Server initializing...")

	server := mqtt.NewServer(nil)

	tcp := listeners.NewTCP("t1", ":1883")
	err := server.AddListener(tcp, nil)
	if err != nil {
		log.Fatal(err)
	}

	server.Serve()
	fmt.Println(("  Started!  "))

	<-done
	fmt.Println(("  Caught Signal  "))

	server.Close()
	fmt.Println(("  Finished  "))
}
