package main

import (
    "fmt"
    "container/list"
    "strconv"
)

type BallTrack struct {
    Max int
    Min int
    Name string
    Groove *list.List
}

func (bt *BallTrack) IsTilted() bool {
    return (bt.Groove.Len() >= bt.Max)
}

func (bt *BallTrack) IsClear() bool {
    return (bt.Groove.Len() <= bt.Min)
}

func (bt *BallTrack) Pop() Ball {
    element := bt.Groove.Front()
    var ball Ball
    if element != nil {
        ball = element.Value.(Ball)
        bt.Groove.Remove(element)
    }
    return ball
}

func (bt *BallTrack) Push(b Ball) {
    bt.Groove.PushFront(b)
}

func (bt *BallTrack) Enqueue(b Ball) {
    bt.Groove.PushBack(b)
}

func (bt *BallTrack) String() string {
    s := "{Name:" + bt.Name + ", Max:" + strconv.Itoa(bt.Max) + ", Min: " + strconv.Itoa(bt.Min)
    s += ", Groove: {"

    e := bt.Groove.Front()
    
    for i := 1; e != nil && i <= bt.Groove.Len(); i++ {
        b := e.Value.(Ball)
        if i > 1 {
            s += ", "
        }
        s += b.String()
        e = e.Next()
    }

    s += "}"
    s += "}"
   
    return s
}



type Ball struct {
    Id int
}

func (b *Ball) String() string {
    return "{Id:" + strconv.Itoa(b.Id) + "}"
}

type BallClock struct {
    
}


func main() {
    fmt.Println("Hello, World. I'm the beginning of the Rolling Ball Clock.")

    minuteTrack := BallTrack{Max:5, Name:"Minute Track", Groove:list.New()}
    for i := 1; i <= 5; i++ {
        minuteTrack.Push( Ball{Id:i} )
    }

    fmt.Println(minuteTrack.String())
    ball := minuteTrack.Pop()

    fmt.Println("Front ball Id: ", ball.Id)
    fmt.Println(minuteTrack.String())


}

