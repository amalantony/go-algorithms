package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

const INFINITY = 1<<32 - 1

func contigSum (a []int) int {
    sum := -INFINITY
    sumArr := make([]int, len(a))
    for index, el := range a {
        if el > sum + el {
            sum = el
        } else {
            sum = sum + el
        }
        sumArr[index] = sum
    }
    largest := -INFINITY
    for _, el := range sumArr {
        if el > largest {
            largest = el
        }
    }
    return largest
}

func nonContigSum (a [] int) int {
    // -1 -2 -3 -4 -5 -6
    // 2 -1 2 3 4 -5
    largestNeg := -INFINITY
    sum := 0
    hasPositive := false
    for _, el := range a {
        if el < 0 {
            if el > largestNeg {
                largestNeg = el
            }
        } else {
            hasPositive = true
            sum += el
        }
    }
    if !hasPositive {
        return largestNeg
    }
    return sum
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    numCasesStr, _ := reader.ReadString('\n')
    numCasesStr = strings.Trim(numCasesStr, " \n")
    numCases, _ := strconv.Atoi(numCasesStr)
    var numArr []int
    output := make([][]int, numCases)
    for i := 0; i < numCases; i++ {
        output[i] = make([]int, 2)
        arrLenStr, _ := reader.ReadString('\n')
        arrLenStr = strings.Trim(arrLenStr, " \n")
        
        numStr, _ := reader.ReadString('\n')
        numStr = strings.Trim(numStr, " \n")
       
        numStrArr := strings.Split(numStr, " ")
        numArr = make([] int, len(numStrArr))
        for index, el := range numStrArr {
           num, _ := strconv.Atoi(el)
           numArr[index] = num
        }
        output[i][0] = contigSum(numArr) 
        output[i][1] = nonContigSum(numArr)
    }
    for i := 0; i < numCases; i++ {
        fmt.Println(output[i][0], output[i][1])    
    }   
}