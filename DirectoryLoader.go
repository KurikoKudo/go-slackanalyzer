package slackanalyzer

import (
	"path/filepath"
	"os"
)

var messageList []Message

//DirectoryLoader is loading directory function
func DirectoryLoader (dirPath string) error{

	err := filepath.Walk(dirPath, func (path string, info os.FileInfo, err error) error {

		if !info.IsDir(){
			fileloader(path)
		}


		return nil
	})

	return err
}

func fileloader(path string){

	//pathのjsonファイルをopenする

	//読み込みにはjsonのパッケージ使えそう

	//Messageに値をセットする

	//appendする
	
}