package config

import (
	"cv/pkg/terminal"
	"cv/pkg/utils"
	"github.com/spf13/viper"
	"os"
)

func Setup() {
	if home := os.Getenv("HOME"); home != "" {
		viper.AddConfigPath(home)
	}
	viper.SetConfigName(FileName)
	viper.SetConfigType(FileType)

	_ = Exists()
}

func Exists() bool {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return false
		}
	}

	return true
}

func Create() string {
	home := getHome()
	f, err := os.Create(home + "/" + FileName + "." + FileType)
	if err != nil {
		terminal.ErrorMessage("Failed to create %s.%s (%s)", FileName, FileType, err.Error())
	}

	defer utils.CloseFileHandler(f, FileName+"."+FileType)

	return f.Name()
}

func Read() []byte {
	home := getHome()
	if data, err := os.ReadFile(home + "/" + FileName + "." + FileType); err != nil {
		terminal.ErrorMessage("Failed to read config from %s.%s (%s)", FileName, FileType, err.Error())
		return nil
	} else {
		return data
	}
}

func Write(data []byte) {
	home := getHome()
	if err := os.WriteFile(home+"/"+FileName+"."+FileType, data, 0); err != nil {
		terminal.ErrorMessage("Failed to write config to %s.%s (%s)", FileName, FileType, err.Error())
	}
}

func fileExists(file string) (bool, string) {
	var info os.FileInfo
	var err error
	home := getHome()
	info, err = os.Stat(home + "/" + file)

	if os.IsNotExist(err) {
		return false, home
	}

	return !info.IsDir(), home
}

func getHome() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	} else {
		return "~"
	}
}
