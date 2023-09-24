package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Config struct {
	App        *App        `mapstructure:"app"`
	Source     *Source     `mapstructure:"source"`
	Cloudinary *Cloudinary `mapstructure:"cloudinary"`
	Wasabi     *Wasabi     `mapstructure:"wasabi"`
	Postgres   *Postgres   `mapstructure:"postgres"`
}

type App struct {
	Mode    string        `mapstructure:"mode"`
	Port    string        `mapstructure:"port"`
	Host    string        `mapstructure:"host"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type Source struct {
	NewsAPI    string        `mapstructure:"news_api"`
	NewsAPIKey string        `mapstructure:"news_api_key"`
	Timeout    time.Duration `mapstructure:"timeout"`
}

type Cloudinary struct {
	CloudName string `mapstructure:"cloud_name"`
	APIKey    string `mapstructure:"api_key"`
	APISecret string `mapstructure:"api_secret"`
}

type Wasabi struct {
	Endpoint    string        `mapstructure:"endpoint"`
	AccessKeyID string        `mapstructure:"access_key_id"`
	SecretKey   string        `mapstructure:"secret_key"`
	Region      string        `mapstructure:"region"`
	Bucket      string        `mapstructure:"bucket"`
	Path        string        `mapstructure:"path"`
	Timeout     time.Duration `mapstructure:"timeout"`
}

type Postgres struct {
	Host        string        `mapstructure:"hostname"`
	Port        string        `mapstructure:"port"`
	User        string        `mapstructure:"user"`
	Password    string        `mapstructure:"password"`
	DataBase    string        `mapstructure:"database"`
	UserTable   string        `mapstructure:"user_table"`
	SurveyTable string        `mapstructure:"survey_table"`
	PhotoTable  string        `mapstructure:"photo_info"`
	Timeout     time.Duration `mapstructure:"timeout"`
}

func (c *Postgres) MakeConnectionURL() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.Host, c.User, c.Password, c.DataBase, c.Port)
}

type file struct {
	path   string `mapstructure:"path"`
	name   string `mapstructure:"name"`
	format string `mapstructure:"format"`
}

func newFile() *file {
	return &file{
		path:   os.Getenv("CONFIG_FILE_PATH"),
		name:   os.Getenv("CONFIG_FILE_NAME"),
		format: os.Getenv("CONFIG_FILE_FORMAT"),
	}
}

func NewConfig() (*Config, error) {
	configFile := newFile()
	err := loadConfig(configFile.path, configFile.name, configFile.format)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = viper.Unmarshal(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func loadConfig(path, fileName, fileFormat string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileFormat)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func PrepareEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
}
