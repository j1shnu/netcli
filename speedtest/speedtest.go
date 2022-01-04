package speedtest

import (
	"encoding/json"
	"fmt"
	"log"
	"net-cli/helpers"
	"time"
)

type fullOutput struct {
	Timestamp outputTime `json:"timestamp"`
	UserInfo  *User      `json:"user_info"`
	Servers   Servers    `json:"servers"`
}
type outputTime time.Time

func SpeedTest(savingMode bool, jsonOutput bool) {
	user, err := FetchUserInfo()
	if err != nil {
		fmt.Println("Warning: Cannot fetch user information. http://www.speedtest.net/speedtest-config.php is temporarily unavailable.")
	}

	if !jsonOutput {
		showUser(user)
	}

	serverList, err := FetchServerList(user)
	helpers.ErrorHandler(err)

	targets, err := serverList.FindServer([]int{})
	helpers.ErrorHandler(err)

	startTest(targets, savingMode, jsonOutput)

	if jsonOutput {
		jsonBytes, err := json.Marshal(
			fullOutput{
				Timestamp: outputTime(time.Now()),
				UserInfo:  user,
				Servers:   targets,
			},
		)
		helpers.ErrorHandler(err)

		fmt.Println(string(jsonBytes))
	}
}

func startTest(servers Servers, savingMode bool, jsonOutput bool) {
	for _, s := range servers {
		if !jsonOutput {
			showServer(s)
		}

		err := s.PingTest()
		checkError(err)

		if jsonOutput {
			err := s.DownloadTest(savingMode)
			checkError(err)

			err = s.UploadTest(savingMode)
			checkError(err)

			continue
		}

		showLatencyResult(s)

		err = testDownload(s, savingMode)
		checkError(err)
		err = testUpload(s, savingMode)
		checkError(err)

		showServerResult(s)
	}

	if !jsonOutput && len(servers) > 1 {
		showAverageServerResult(servers)
	}
}

func testDownload(server *Server, savingMode bool) error {
	quit := make(chan bool)
	fmt.Printf("Download Test : ")
	go dots(quit, "//")
	err := server.DownloadTest(savingMode)
	quit <- true
	if err != nil {
		return err
	}
	fmt.Println()
	return err
}

func testUpload(server *Server, savingMode bool) error {
	quit := make(chan bool)
	fmt.Printf("Upload Test   : ")
	go dots(quit, "\\\\")
	err := server.UploadTest(savingMode)
	quit <- true
	if err != nil {
		return err
	}
	fmt.Println()
	return nil
}

func dots(quit chan bool, bar string) {
	for {
		select {
		case <-quit:
			return
		default:
			time.Sleep(time.Second)
			fmt.Print(bar)
		}
	}
}

func showUser(user *User) {
	if user.IP != "" {
		fmt.Printf("Testing From IP: %s\n", user.String())
	}
}

func showServer(s *Server) {
	fmt.Printf(" \n")
	fmt.Printf("Target Server: [%4s] %8.2fkm ", s.ID, s.Distance)
	fmt.Printf(s.Name + " (" + s.Country + ") by " + s.Sponsor + "\n")
}

func showLatencyResult(server *Server) {
	fmt.Println("Latency:", server.Latency)
}

// ShowResult : show testing result
func showServerResult(server *Server) {
	fmt.Printf(" \n")

	fmt.Printf("Download: %5.2f Mbit/s\n", server.DLSpeed)
	fmt.Printf("Upload: %5.2f Mbit/s\n\n", server.ULSpeed)
	valid := server.CheckResultValid()
	if !valid {
		fmt.Println("Warning: Result seems to be wrong. Please speedtest again.")
	}
}

func showAverageServerResult(servers Servers) {
	avgDL := 0.0
	avgUL := 0.0
	for _, s := range servers {
		avgDL = avgDL + s.DLSpeed
		avgUL = avgUL + s.ULSpeed
	}
	fmt.Printf("Download Avg: %5.2f Mbit/s\n", avgDL/float64(len(servers)))
	fmt.Printf("Upload Avg: %5.2f Mbit/s\n", avgUL/float64(len(servers)))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (t outputTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05.000"))
	return []byte(stamp), nil
}
