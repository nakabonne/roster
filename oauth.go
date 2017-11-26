package main

var scopes = []string{
	"https://www.googleapis.com/auth/spreadsheets",
	"https://www.googleapis.com/auth/drive",
	"https://www.googleapis.com/auth/drive.file"}

func authorize() {
	for i, scope := range scopes {
		config, err := google.ConfigFromJSON(b, scope)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}
		giveScope(config)
	}
}
