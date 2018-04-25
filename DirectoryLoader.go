package slackanalyzer

import (
	"path/filepath"
	"os"
)

//DirectoryLoader is load directory function
func DirectoryLoader (dirPath string) error{

	err := filepath.Walk(dirPath, func (path string, info os.FileInfo, err error) error {

		if !info.IsDir(){
			fileloder(path)
		}


		return nil
	})

	return err
}

func fileloder(path string){
	
}