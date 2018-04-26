package slackanalyzer

import (
	"path/filepath"
	"os"
)

var messageList []Message

//DirectoryLoader is loading directory function
func DirectoryLoader (rootPath string) error{

	err := filepath.Walk(rootPath, walkFunc)
	
	return err
}

func walkFunc(filePath string, info os.FileInfo, err error) error {

	if !info.IsDir(){
		message := fileLoader(filePath)
		messageList = append(messageList,message)
	}

	return nil
}

func fileLoader(filePath string) Message {

	var message Message

	//pathのjsonファイルをopenする

	//とりあえず
	_ = filePath

	//読み込みにはjsonのパッケージ使えそう

	//Messageに値をセットする

	//appendする

	return message
}