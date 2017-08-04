package src

import (
	"fmt"
	"github.com/bndr/gojenkins"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func init() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	jenkins := gojenkins.CreateJenkins(
		"http://10.10.129.37:8080/",
		"admin", "bd46e1d7c7384fb9b0c57014f8854a37")
	_, err := jenkins.Init()
	if err != nil {
		print(err)
		log.Error("init failed.")
	}

	rs, err := jenkins.GetAllJobs()
	if err != nil {
		print(err)
		log.Error("init failed.")
	}
	job := rs[0]

	println(job.GetName())
	println(job.GetDescription())

	buildIDS, err := job.GetAllBuildIds()
	builds := make([]*gojenkins.Build, 0)

	log.Info("get all builds ...")
	for _, v := range buildIDS {
		build_num := v.Number
		build, err := job.GetBuild(build_num)
		if err != nil {
			println(err)
		}
		timestamp := build.GetTimestamp()
		log.Info(fmt.Sprintf("get build num: %v,%v", build_num, timestamp))
		if timestamp.Month() == time.June {
			builds = append(builds, build)
			//log.Info(fmt.Sprintf("%v", timestamp.Month()))
		}
	}

	log.Infoln(len(builds))
	log.Infoln(builds[0].GetConsoleOutput())

}
