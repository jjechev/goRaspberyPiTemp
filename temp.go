package main

import "fmt"
import "os/exec"

func main() {
	fmt.Println(getTempCPU())
	fmt.Println(getTempGPU())
}

func getTempCPU() string {
	tempcmd, _ := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp").Output()
	temp := string(tempcmd)
	temp = temp[:2] + "." + temp[2:3]
	temp = "CPU temp=" + temp + "'C"
	return temp
}

func getTempGPU() string {
	tempcmd, _ := exec.Command("/opt/vc/bin/vcgencmd", "measure_temp").Output()
	temp := string(tempcmd)
	temp = "GPU " + temp
	return temp
}
