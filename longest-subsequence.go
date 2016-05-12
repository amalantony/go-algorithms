/* 
  1) Consider the input strings “AGGTAB” and “GXTXAYB”. Last characters match for the strings. So length of LCS can be written as:

    L(“AGGTAB”, “GXTXAYB”) = 1 + L(“AGGTA”, “GXTXAY”)

  2) Consider the input strings “ABCDGH” and “AEDFHR. Last characters do not match for the strings. So length of LCS can be written as:
    L(“ABCDGH”, “AEDFHR”) = MAX ( L(“ABCDG”, “AEDFHR”), L(“ABCDGH”, “AEDFH”) )
*/


package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func max (x int, y int) (int) {
    if x > y {
        return x
    } 
    return y
}

func longestSubSequence (s1 string, s2 string) (int) {
    m := len(s1)
    n := len(s2)
    
    V := make([][]int, m+1)
        
    for i := 0; i <= m; i++ {
        V[i] = make([]int, n+1)
        for j := 0; j <= n; j++ {
            if i == 0 || j == 0 {
                V[i][j] = 0
            } else if s1[i - 1] == s2[j - 1] {
                V[i][j] = 1 + V[i-1][j-1]
            } else {
                V[i][j] = max(V[i-1][j], V[i][j-1])
            }
        }
    } 
    return V[m][n]
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    str1, _ := reader.ReadString('\n')
    str2, _ := reader.ReadString('\n')
    fmt.Println(longestSubSequence(strings.Trim(str1, " \n"), strings.Trim(str2, "\n ")))   
}