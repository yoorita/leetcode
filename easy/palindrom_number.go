package main

// import "fmt"

// func main(){
// 	fmt.Println(isPalindrome(121))
// 	fmt.Println(isPalindrome(-121))
// 	fmt.Println(isPalindrome(10))
// }

func isPalindrome(x int) bool {
    reversed := 0
	temp := x

	for temp > 0 {
		reversed = (reversed * 10 )+ (temp % 10)
		temp /= 10
	}

	return reversed == x
}