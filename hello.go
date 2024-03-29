package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 5
const delay = 5

func main() {
	intro()
	for{
		fmt.Println("")
		showMenu()
		command := inputReader()
		switch command {
			case 1:
				startMonitoring()		
			case 2:
				fmt.Println("Exibir logs")
				printLogs()
			case 0:
				fmt.Println("Saindo do programa")
				os.Exit(0)
			default:
				fmt.Println("Comando invalido")
				os.Exit(-1)
			}
		}
	}
	
func intro() {
	var name = "Caio"
	var version float32 = 1.1	
	fmt.Println("Olá sr.", name)
	fmt.Println("Esse código está na versão ", version)
}

func showMenu() {
	fmt.Println("-----------------------------------")
	fmt.Println("1- Iniciar Monitoramento ")
	fmt.Println("2- Exibir Logs ")
	fmt.Println("0- Sair do programa ")
	fmt.Println("-----------------------------------")
}

func inputReader() int{
	var command int
	fmt.Scan(&command)
	fmt.Println("O Comando escolhido foi :", command)
	return command
}

func printLogs(){
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um error: ", err)
	}

	fmt.Println(string(file))
}

func startMonitoring(){
		fmt.Println("Monitorando...")
		sites := openFile()
		for i:=0; i< monitoring; i++ {
			for i,site := range sites {
				fmt.Println("Testando site",i+1,":", site)
				siteTest(site)
			}
			time.Sleep(delay * time.Second)
			fmt.Println("")
		}
		fmt.Println("")
	}
	
	func openFile() [] string{
		var output [] string
		file,err :=os.Open("sites.txt");

		if err!= nil{
			fmt.Println("Error: ",err)
		}

		reader := bufio.NewReader(file)
		
		for{
			line, err := reader.ReadString('\n')
			line = strings.TrimSpace(line)
			
			if err == io.EOF{
				break
			}

			output=append(output,line)
		}

		file.Close()
		return output
	}

	func siteTest(site string){
		resp, err := http.Get(site)
		
		if err != nil {
			fmt.Println("Ocorreu um erro :",err)
		}

		if resp.StatusCode==200{
			fmt.Println("O site ", site, " está em pleno funcionamento!")
			logRegistry(site,true)
		}else{
			fmt.Println("O site", site, " não funciona adequadamente! Status Code:", resp.StatusCode)
			logRegistry(site,false)
		}	
	}


	

	func logRegistry (site string, status bool){
		file, err := os.OpenFile("log.txt",os.O_RDWR | os.O_CREATE | os.O_APPEND,0666);

		if err != nil {
			fmt.Println("Ocorreu um erro: ",err)
		}

		file.WriteString(site + " - online: "+ strconv.FormatBool(status) + " horário: "+ time.Now().Format("02/01/2006 15:04:05") + "\n")
		file.Close()
	}