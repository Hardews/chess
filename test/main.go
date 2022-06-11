package main

import (
	"chess/chessServer"
	"fmt"
)

func main() {
	client1, client2 := chessServer.NewClient()
	for i := 0; i < 100; i++ {
		switch {
		case i%2 == 0:
			cli(client1, client2)
		case i%2 != 0:
			cli2(client2, client1)
		}
	}
}

func cli(client *chessServer.Client, gre *chessServer.Client) {
	client.Sync(gre)
	fmt.Println("red do")
	client.Print()
	var z, a, b, c int
	fmt.Scan(&z, &a, &b, &c)
	switch z {
	case 1:
		err := client.GunMove(a, b, c)
		if err != nil {
			fmt.Println(err)
		}
	case 2:
		err := client.HorseMove(a, b)
		if err != nil {
			fmt.Println(err)
		}
	case 3:
		err := client.GuardMove(a, b)
		if err != nil {
			fmt.Println(err)
		}
	case 4:
		err := client.VehicleMove(a, b, c)
		if err != nil {
			fmt.Println(err)
		}
	case 5:
		err := client.MinisterMove(a, b)
		if err != nil {
			fmt.Println(err)
		}
	case 6:
		err := client.CommanderMove(b)
		if err != nil {
			fmt.Println(err)
		}
	case 7:
		err := client.SoldierMove(a, b)
		if err != nil {
			fmt.Println(err)
		}
	}
	client.Print()
}

func cli2(client *chessServer.Client, red *chessServer.Client) {
	client.Sync(red)
	fmt.Println("gre do")
	client.Print()
	var z, a, b, c int
	fmt.Scan(&z, &a, &b, &c)
	switch z {
	case 1:
		err := client.GunMove(a, b, c)
		if err != nil {
			fmt.Println(err)
		}
	case 2:
		err := client.HorseMove(a, b)
		if err != nil {
			fmt.Println(err)
		}
	case 3:
		err := client.GuardMove(a, b)
		if err != nil {
			fmt.Println(err)
		}
	case 4:
		err := client.VehicleMove(a, b, c)
		if err != nil {
			fmt.Println(err)
		}
	case 5:
		err := client.MinisterMove(a, b)
		if err != nil {
			fmt.Println(err)
		}
	case 6:
		err := client.CommanderMove(b)
		if err != nil {
			fmt.Println(err)
		}
	case 7:
		err := client.SoldierMove(a, b)
		if err != nil {
			fmt.Println(err)
		}
	}
	client.Print()

}
