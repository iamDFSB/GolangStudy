package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar irá retornar a aplicação de linha de comando
func Gerar() *cli.App{
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Buscar IPs e Nomes de Servidor na internet"

	flags := []cli.Flag{
				cli.StringFlag{
					Name: "host", 
					Value: "google.com.br",
				},
			}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca de Ips",
			Flags: flags,
			Action: searchIps,
		},
		{
			Name: "server",
			Usage: "Buscar o nome dos servidores na internet",
			Flags: flags,
			Action: searchServer,
		},
	}

	return app
}


func searchIps(c *cli.Context){
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	
	if erro != nil{
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServer(c *cli.Context){
	host := c.String("host")
	servers, erro := net.LookupNS(host)
		
	if erro != nil{
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server)
	}

}