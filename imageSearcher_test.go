package main

//121024768

//func TestImageSearcher_SearchImage(t *testing.T) {
//	config := NewImageSearchConfig().
//		SetKey("1f2591f0famsh99bcccece234288p1bb55djsn707a4e017919").
//		SetUrl("https://contextualwebsearch-websearch-v1.p.rapidapi.com/api/Search/ImageSearchAPI")
//	imageSearcher := NewImageSearcher()
//	imageSearcher.setConfig(config)
//	ee, _ := imageSearcher.SearchImage("kingdom")
//	println(ee)
//}
//
//func TestDownloadImage(t *testing.T) {
//	data, err := DownloadImage("https://i0.wp.com/www.droidgamers.com/wp-content/uploads/2017/06/kingdom-new-lands-android.jpg?fit=1920%2C1080&ssl=1")
//	//data, _ := ioutil.ReadFile("tests/image.jpg")
//	b := tgbotapi.FileBytes{Name: "image.jpg", Bytes: data}
//
//	msg := tgbotapi.NewPhotoUpload(121024768, b)
//	msg.Caption = "Test"
//
//	viper.SetConfigFile("./config.toml")
//	if err := viper.ReadInConfig(); err != nil {
//		log.Fatalf("Error reading config file, %s", err)
//	}
//
//	telegramToken := viper.GetString("telegram.token")
//	if telegramToken == "" {
//		log.Fatalf("Telegram token are not set in config.toml")
//	}
//	log.Println(telegramToken)
//	bot, err := tgbotapi.NewBotAPI(telegramToken)
//	if err != nil {
//		log.Panic(err)
//	}
//
//	bot.Debug = true
//
//	log.Printf("Authorized on account %s", bot.Self.UserName)
//
//	u := tgbotapi.NewUpdate(0)
//	u.Timeout = 60
//	//bot.GetChatMember(tgbotapi.ChatConfigWithUser{})
//	_, err = bot.Send(msg)
//
//	if err != nil {
//		t.Error(err)
//		t.Fail()
//	}
//}
