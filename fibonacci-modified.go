 
/* A series is defined in the following manner:

Given the nth and (n+1)th terms, the (n+2)th can be computed by the following relation 
Tn+2 = (Tn+1)2 + Tn

So, if the first two terms of the series are 0 and 1: 
the third term = 12 + 0 = 1 
fourth term = 12 + 1 = 2 
fifth term = 22 + 1 = 5 
... And so on.

Given three integers A, B and N, such that the first two terms of the series (1st and 2nd terms) are A and B respectively, compute the Nth term of the series. 
*/

package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "math/big"
)

func fibModified (a *big.Int, b *big.Int, n *big.Int) (* big.Int) {
    t1 := a
    t2 := b
    t3 := big.NewInt(0)
    i := big.NewInt(3)
    for {
        cmp := i.Cmp(n)
        if (cmp == -1 || cmp == 0) {
            t3 = big.NewInt(0)
            t3.Mul(t2, t2)
            t3.Add(t3, t1)
            t1 = t2
            t2 = t3
            i = i.Add(i, big.NewInt(1))
        } else {
            break   
        }
    }
    
    return t3
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    a, _ := reader.ReadString(' ')
    b, _ := reader.ReadString(' ')
    n, _ := reader.ReadString('\n')
    
    a = strings.Trim(a, " \n")
    b = strings.Trim(b, " \n")
    n = strings.Trim(n, " \n")
    
    a1, _ := strconv.ParseInt(a, 0, 64)
    b1, _ := strconv.ParseInt(b, 0, 64)
    n1, _ := strconv.ParseInt(n, 0, 64)
    
    
    res := fibModified(big.NewInt(a1), big.NewInt(b1), big.NewInt(n1))
    
    fmt.Println(res.String())
}
