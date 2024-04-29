package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	// "net/http"
	// "os"
	"time"
	"strings"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	
	
)


var id int //идентификатор в бд

func main() {
	
	to_database("a b c")
	// // подключение к mqtt брокеру Rightech
	// opts := MQTT.NewClientOptions()
	// opts.AddBroker("tcp://dev.rightech.io:1883")
	// opts.SetClientID("mqtt_db_go")
	   
	// client := MQTT.NewClient(opts)
	// if token := client.Connect(); token.Wait() && token.Error() != nil {
	// 	fmt.Println(token.Error())
	// 	os.Exit(1)
	// }
	// fmt.Println("Connected to Rightech MQTT server") 
	// // подписываемся на топик  
	// client.Subscribe("base/state/user", 0, onMessageReceived)
	   
	// for {
	// 	time.Sleep(1 * time.Second)
	// }
}



// если получены данные от брокера, вызываеися функция для записи их в бд
func onMessageReceived(client MQTT.Client, message MQTT.Message) {
 fmt.Printf("Received message on topic: %s\n", message.Topic())
 fmt.Printf("Message: %s\n", message.Payload())
 to_database(string(message.Payload()))
}




// инициализация базы данных Firebase Realtime Database
func InitializeAppWithServiceAccount() *firebase.App {
	opt := option.WithCredentialsFile("db-go-44965-firebase-adminsdk-86osd-cbd0a5b355.json")
	config := &firebase.Config{
		DatabaseURL: "https://db-go-44965-default-rtdb.europe-west1.firebasedatabase.app/",
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("Ошибка при инициализации приложения Firebase: %v\n", err)
	}
	return app
}

// проверка текущего id в базе данных 
func Check_id() {

	app := InitializeAppWithServiceAccount()
	client, err := app.Database(context.Background())
	if err != nil {
		log.Fatalf("Ошибка при инициализации Realtime Database клиента: %v\n", err)
	}
	ref := client.NewRef("id")

	if err := ref.Get(context.Background(), &id); err != nil {
		log.Fatalln("Error reading value:", err)
	}
	// меняем id на следующий (+1)
	if err := ref.Set(context.Background(), id+1); err != nil {
		log.Fatalf("Ошибка при записи id в бд: %v\n", err)
	}else{
		fmt.Println("id записан в бд")
	}
	// при удалении данных обновлять id на следующий
}



// запись в бд
func to_database(data string) {
	data_mas := strings.Split(data, " ")
	tel, complect_num := data_mas[1], data_mas[0]
	// Получаем текущую дату и время
	currentTime := time.Now()

	// Выводим текущую дату и время в формате по умолчанию
	fmt.Println("Текущая дата и время:", currentTime)
   
	// Определяем отдельно текущую дату и время
	currentDate := currentTime.Format("2006-01-02")
	currentTimeOfDay := currentTime.Format("15:04:05")
   
	fmt.Println("Текущая дата:", currentDate)
	fmt.Println("Текущее время:", currentTimeOfDay)
	date:= currentDate+" "+currentTimeOfDay
	app := InitializeAppWithServiceAccount()
	client, err := app.Database(context.Background())
	if err != nil {
		log.Fatalf("Ошибка при инициализации Realtime Database клиента: %v\n", err)
	}

	// проверка id по которому записываем
	Check_id()
	log.Println(id)
	ref := client.NewRef(strconv.Itoa(id))
	
	datalist := map[string]interface{}{
		"telephone": tel,
		"complect_num":  complect_num,
		"date":date,
	}
	// if err := ref.Get(context.Background(), &datalist); err != nil {
	// 	return
	// }
	fmt.Printf("%v", datalist["telephone"])
	fmt.Printf("%v", datalist["complect_num"])
	fmt.Printf("%v", datalist["date"])

	
	if err := ref.Set(context.Background(), datalist); err != nil {
		log.Fatalf("Ошибка при записи данных в бд: %v\n", err)
	}else{
		fmt.Println("Данные записаны в бд")
	}
	
}


   
   
