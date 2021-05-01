package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/yagoazedias/dog-escaper/infraestruture"
	"github.com/yagoazedias/dog-escaper/respository"
	"strconv"
)

func sub(client mqtt.Client) {
	topic := infraestruture.Config["MQTT_TOPIC"]
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s\n", topic)
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n\n", msg.Payload(), msg.Topic())
	payload := string(msg.Payload()[:])
	newStatus, err := strconv.ParseBool(payload)

	if err != nil {
		fmt.Printf("Could not parse mensage to boolean, message value was: %s", msg.Payload())
	}

	repo := respository.PortRepository{}
	_, err = repo.UpdateLastStatus(newStatus)

	if err != nil {
		fmt.Printf("Could not save new port status on database: %s", err.Error())
		return
	}

	fmt.Printf("Successfully saved new port status")
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v\n", err)
}

func ConfigureMQQT() {
	var broker = infraestruture.Config["MQTT_HOST"]
	var port = infraestruture.Config["MQTT_PORT"]

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", broker, port))
	opts.SetClientID("dog_escaper")
	opts.SetUsername("dog_escaper")
	opts.SetPassword("testing")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
}
