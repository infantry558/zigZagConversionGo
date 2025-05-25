package main
import ("fmt"; "strings")

func main(){
	//var integer int = 32;
	//fmt.Println(integer)
	//var string = "Hello \nWorld!"
	//fmt.Println("The string is " + string)
	fmt.Println(convert("PAYPALISHIRING", 4))
}


func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    res := make([]string, numRows)
    row := 0
    dir := 1

    for i:=0; i<len(s); i++ {
    	res[row] += string(s[i])
    	if row == numRows - 1 {
    		dir = -1
    	}else if row == 0 && dir == -1 {
    		dir = 1
    	}
    	row = row + dir
    }


    return strings.Join(res[:], "")

}
