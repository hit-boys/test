package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
    "errors"
)

var m = map[uint8]int{'I' : 1, 'V' : 5, 'X' : 10, 'L' : 50, 'C' : 100}

func convert(inst string) (int, bool, error)  {
    x, _ := strconv.Atoi(strings.TrimSpace(inst))
    what := false
 
     if(x == 0){
       
        tmp :=  m[inst[0]]
        for i := 0; i < len(strings.TrimSpace(inst)); i++{
            now, chek := m[inst[i]]
          
            if(!chek){
                err := errors.New("invalid number")
                return 0, false, err
            }
          
 
            x += now
            if(tmp < now){
                x -= tmp * 2
            }
 
            tmp = now
        }
 
        what = true
    }
    if(x < 1 || x >10){
       err := errors.New("invalid number")
       return 0, false, err
     }
     
	return x, what, nil
}

func Print(ans int, chek bool) error {
    if(chek){
        if(ans <= 0){
            err := errors.New("there is no number in the Roman system")
            return err
        }
        var mp = map[int]string{100 : "C", 90 : "XC", 50: "L", 40 : "XL", 10 : "X", 9 : "IX", 5 : "V", 4 : "IV", 1 : "I"}
        vec := [9]int{1, 4, 5, 9, 10, 40, 50, 90, 100}
        result := ""
        tmp := len(vec) - 1
        for(ans > 0){
            if (ans / vec[tmp] > 0){
                for i := 0; i <  ans / vec[tmp]; i++{
                    result += mp[vec[tmp]] 
                    ans -= vec[tmp]
                }
            }else {
                tmp--
            }
        }
        fmt.Println(result)
        return nil
    }
    fmt.Println(ans)
    return nil
}

func main() {
    reader := bufio.NewReader(os.Stdin) 
    
    for{
    
    val, _ := reader.ReadString('\n')
    vector := strings.Split(val, " ")
    
	if(len(vector) != 3){
	    err := errors.New("incorrect input")
	    fmt.Println(err)
	    return
	}

    num1, num1_chek, err :=  convert(vector[0])
    if (err != nil){
        fmt.Println(err)
	    return
    }
    num2, num2_chek, err :=  convert(vector[2])
    if (err != nil){
        fmt.Println(err)
	    return
    }
    
    if(num1_chek != num2_chek){
        err := errors.New("different number systems")
        fmt.Println(err)
	    return
    }
    
    switch vector[1]{
        case "+":
            err := Print(num1 + num2, num1_chek)
            if (err != nil){
                fmt.Println(err)
               return
            }
        case "-":
            err := Print(num1 - num2, num1_chek)
            if (err != nil){
                fmt.Println(err)
               return
            }
        case "/":
            err := Print(num1 / num2, num1_chek)
            if (err != nil){
                fmt.Println(err)
               return
            }
        case "*":
            err := Print(num1 * num2, num1_chek)
            if (err != nil){
                fmt.Println(err)
               return
            }
        default:
            err := errors.New("incorrect sign")
            fmt.Println(err)
	        return

    }
    
    }
}
