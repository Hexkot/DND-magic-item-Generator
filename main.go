package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Выбор случайного предмета
func itemChoice() (string, string, int) {
	// Список предметов
	file, err := os.Open("./items.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	itemsDict, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	itemIndex := rand.Intn(len(itemsDict) + 1) 		// Индекс предмета
	descriptionIndex := rand.Intn(3) + 2 			// Индекс описания
	itemName := itemsDict[itemIndex][0] 			// Название предмета
	itemDescription := itemsDict[itemIndex][descriptionIndex] // Описание предмета
	itemGenus, _ := strconv.Atoi(itemsDict[itemIndex][1]) // Для индексации родовых окончания
	return itemName, itemDescription, itemGenus
}

// Выбор случайного минорного свойства
func minorChoice() (string, string) {
	// список предметов
	file, err := os.Open("./minor.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	minorDict, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	minorIndex := rand.Intn(len(minorDict) + 1)
	//
	minorName := minorDict[minorIndex][0]
	minorDescription := minorDict[minorIndex][1]
	return minorName, minorDescription
}

// Выбор случайного мажорного свойства
func majorChoice() (string, string) {
	file, err := os.Open("./major.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	majorDict, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	majorIndex := rand.Intn(len(majorDict) + 1)
	majorName := majorDict[majorIndex][0]
	majorDescription := majorDict[majorIndex][1]
	return majorName, majorDescription
}

// Выбор случайного специального свойства и подбор родового окончания
func specialChoice(itemGenus int) (string, string) {
	file, err := os.Open("./special.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	specialDict, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	specialIndex := rand.Intn(len(specialDict) + 1)
	declensionIndex, _ := strconv.Atoi(specialDict[specialIndex][1]) // индекс грамматического склонения
	specialGenus := func(itemGenus int) string {                     // выбор родового окончания
		file, err := os.Open("./genus.csv")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		reader := csv.NewReader(file)

		genusDict, err := reader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		genusEnding := genusDict[declensionIndex][itemGenus]
		return genusEnding
	}
	specialName := fmt.Sprint(specialDict[specialIndex][0], specialGenus(itemGenus))
	specialDescription := specialDict[specialIndex][2]
	return specialName, specialDescription
}

// Создание обычного предмета
func createCommon() string {
	itemName, itemDescription, _ := itemChoice() //получение предмета
	minorName, minorDescription := minorChoice() //получение минорного свойства
	itemTxt := fmt.Sprint("*", itemName, " ", minorName, ".* ", itemDescription, " ", minorDescription)
	return itemTxt
}

// Создание необычного предмета
func createUncommon() string {
	itemName, itemDescription, _ := itemChoice()
	majorName, majorDescription := minorChoice()
	itemTxt := fmt.Sprint("*", itemName, " ", majorName, ".* ", itemDescription, " ", majorDescription)
	return itemTxt
}

// Создание редкого предмета
func createRare() string {
	itemName, itemDescription, itemGenus := itemChoice()        //получение предмета
	minorName, minorDescription := minorChoice()                //получение минорного свойства
	specialName, specialDescription := specialChoice(itemGenus) //получение спец. свойства
	// блок описания:
	descriptionBlock := fmt.Sprint(itemDescription, minorDescription,
		strings.Replace(specialDescription, "Пока вы настроены на артефакт,", " Также", 1))
	itemTxt := fmt.Sprint("*", specialName, " ", strings.ToLower(itemName), " ", minorName, ".* ", descriptionBlock)
	return itemTxt
}

// Создание очень редкого предмета
func createVeryRare() string {
	itemName, itemDescription, itemGenus := itemChoice()        //получение предмета
	majorName, majorDescription := minorChoice()                //получение минорного свойства
	specialName, specialDescription := specialChoice(itemGenus) //получение спец. свойства
	// блок описания:
	descriptionBlock := fmt.Sprint(itemDescription, majorDescription,
		strings.Replace(specialDescription, "Пока вы настроены на артефакт,", " Также", 1))
	itemTxt := fmt.Sprint("*", specialName, " ", strings.ToLower(itemName), " ", majorName, ".* ", descriptionBlock)
	return itemTxt
}

func main() {
	bot, err := tgbotapi.NewBotAPI("AwesomeAPI")

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/common":
				var newItem string = createCommon()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, newItem)
				msg.ParseMode = "markdown"
				bot.Send(msg)
			case "/uncommon":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, createUncommon())
				msg.ParseMode = "markdown"
				bot.Send(msg)
			case "/rare":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, createRare())
				msg.ParseMode = "markdown"
				bot.Send(msg)
			case "/very_rare":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, createVeryRare())
				msg.ParseMode = "markdown"
				bot.Send(msg)
			case "/legendary":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Легендарных предметов пока не подвезли.")
				msg.ParseMode = "markdown"
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Undefined command")
				bot.Send(msg)
			}

		}
	}
}
