package main

import (
        "fmt"
)

const MAXCOLOR = 6
const MAXNUM   = 10
const MAXCARD  = 60

type Card struct {
        Num     int
        Color   int
}

type Board struct {
        Board [3][9]int
}

type Player struct {
        Cards   [8]Card
}

type Bag struct {
        Cards   [60]Card
}

func (b *Bag) init() {
        for k:=0; k<MAXCARD; k++ {
                b.Cards[k].Color = k / MAXCOLOR
                b.Cards[k].Num   = k % MAXCOLOR  // check
        }
}

func (b *Bag) Prnt() {
        for j:=0; j<MAXCOLOR; j++ {
                for i:=0 ; i<MAXNUM; i++ {
                        fmt.Printf("[%d %d]", b.Cards[j*i+i].Color, b.Cards[j*i+i].Num) // check
                }
                fmt.Printf("\n")
        }
}

func isStraight(a, b, c Card) bool {
        if (a.Num + 1 == b.Num) && (b.Num + 1 == c.Num) {
                return true
        } else {
                return false
        }
}

func isSameColor(a, b, c Card) bool {
        if (a.Color  == b.Color) && (b.Color == c.Color) {
                 return true
        } else {
                return false
        }
}

func isSameNum(a, b, c Card) bool {
        if (a.Num == b.Num) && (b.Num == c.Num) {
                return true
        } else {
                return false
        }
}

func getValue(a, b, c Card) int {
        value := a.Num + b.Num + c.Num

        if isStraight(a,b,c) {
                value = value + 100
        }

        if isSameColor(a,b,c) {
                value = value + 1000
        }

        if isSameNum(a,b,c) {
                value = value + 10000
        }

        if isStraight(a,b,c) && isSameColor(a,b,c) {
                value = value + 100000
        }

        return value
}

func main() {
        b := Bag{}
        b.init()
        b.Prnt()
        fmt.Printf("Battle Line\n")
}
