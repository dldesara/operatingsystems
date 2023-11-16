package main

import(
		"fmt"
		"syscall"

)

func openFile(){

		fmt.Println("vim-go")
		filename := "bytes.html"
		// open in read only mode
		flags := syscall.O_RDONLY
		mode := uint32(0666) // file permissions
		fd, err := syscall.Open(filename, flags, mode)
		if err != nil {
				fmt.Printf("Error !!!: %s\n", err)
		}
		defer syscall.Close(fd)
		fmt.Printf("Opened file: %s\n", filename)


}

//func getSysinfo(){

	//var info syscall.Sysinfo_t
	//if err := syscall.Sysinfo(&info); err != nil {
		//panic(err)
	//}
	//fmt.Printf("Total RAM: %v bytes\n", info.Totalram)
	//fmt.Printf("Free RAM: %v bytes\n", info.Freeram)
	//fmt.Printf("Uptime: %v bytes\n", info.Uptime)
	
//}

func createProcess() {
	err := syscall.Exec("/bin/ls", []string{"ls", "-la"}, syscall.Environ())
	if err != nil{
	fmt.Println("Error !!!: ", err)
	return
	
	}
	
	fmt.Println("Process create successfully: ")

}

func main(){
	//getSysinfo()
	createProcess()
}