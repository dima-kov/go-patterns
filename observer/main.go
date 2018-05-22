package main

func main() {
	shop := Shop{}
	client1 := Client{}
	client2 := Client{}
	shop.Subscribe(&client1)
	shop.Subscribe(&client2)

	shop.NotifySubscribers("We have super-puper new product")
    //	Gotcha
}
