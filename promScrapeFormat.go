package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ScrapeConfigs []ScrapeConfig `yaml:"scrapeconfig"`
}
type ScrapeConfig struct {
	JobName string `yaml:"job_name"`
	// MetricsPath    string        `yaml:"metrics_path"`
	// ScrapeInterval string        `yaml:"scrape_interval"`
	StaticConf StaticConfigs `yaml:"static_configs"`
}
type StaticConfigs struct {
	Targets []string `yaml:"targets"`
}

func isStrStart(str string) bool {
	firstRune := []rune(str)[0]
	if unicode.IsLetter(firstRune) {
		return true
	} else {
		return false
	}
}

func ReadSrcFile(path string) []*ScrapeConfig {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("读取文件%s失败\n", file)
		return nil
	}
	datas := strings.Split(string(file), "\n")
	metrics := make(map[string]*ScrapeConfig)
	var currentConfig *ScrapeConfig
	var job string
	for _, i := range datas {
		if i != "" {
			isletter := isStrStart(i)
			if isletter {
				job = i
				currentConfig = &ScrapeConfig{
					JobName:    i,
					StaticConf: StaticConfigs{Targets: []string{}},
				}

			} else {
				// currentConfig.JobName = job
				if i != "" {
					currentConfig.StaticConf.Targets = append(currentConfig.StaticConf.Targets, i)
				}
				metrics[job] = currentConfig
			}
		}
	}
	confs := []*ScrapeConfig{}
	for _, v := range metrics {
		confs = append(confs, v)
	}
	return confs
}
func main() {
	confs := ReadSrcFile("data")
	fmt.Println(confs)
	ymldata, err := yaml.Marshal(confs)
	if err != nil {
		fmt.Println("数据格式化为yaml失败", err)
	}
	fmt.Println(string(ymldata))
}
