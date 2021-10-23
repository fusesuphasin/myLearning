package main

func h() {
	defer func(){
		r := recover(); r != nil{ // r != nil -> เกิด panic
			fmt.Println("recover")
		}
	}
}

func main(){
	h()
}