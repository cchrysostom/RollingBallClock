package main

import (
    "fmt"
    "container/list"
    "encoding/json"
    "strconv"
    "os"
)

type BallTrack struct {
    Max int
    Name string
    Groove *list.List
}

func (bt *BallTrack) IsTilted() bool {
    return (bt.Groove.Len() >= bt.Max)
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

func (bt BallTrack) String() string {
    s := "{\"Name\":\"" + bt.Name + "\", \"Max\":" + strconv.Itoa(bt.Max)
    s += ", \"Groove\": ["

    e := bt.Groove.Front()
    
    for i := 1; e != nil && i <= bt.Groove.Len(); i++ {
        b := e.Value.(Ball)
        if i > 1 {
            s += ", "
        }
        s += b.String()
        e = e.Next()
    }

    s += "]"
    s += "}"
   
    return s
}

// TODO: MarshalJSON does not get invoked. Look for correct solution.
func (bt *BallTrack) MarshalJSON() ([]byte, error) {
    s := "{" + "\"max\":\"" + strconv.Itoa(bt.Max) + "\""
    s += ",\"name\":\"" + bt.Name + "\""
    s += ",\"groove\":" 

    e := bt.Groove.Front()
    s += "["
    firstPass := true
    for e != nil {
        b := e.Value.(Ball)
        if !firstPass {
            s += ","
        } else {
            firstPass = false
        }
        s += "\"id\":" + strconv.Itoa(b.Id)
        e = e.Next()
    }
    s += "]"

    s += "}"
    
    return []byte(s), nil
}

type Ball struct {
    Id int
}

func (b *Ball) String() string {
    return "{\"Id\":" + strconv.Itoa(b.Id) + "}"
}

type BallClock struct {
    BallCount int
    MinuteTrack BallTrack
    FiveMinuteTrack BallTrack
    HourTrack BallTrack
    ReturnTrack BallTrack
}

func (bc *BallClock) DisplayTracks() string {
    s := "BallCount:" + strconv.Itoa(bc.BallCount) + "\n"
    s += bc.MinuteTrack.String() + "\n"
    s += bc.FiveMinuteTrack.String() + "\n"
    s += bc.HourTrack.String() + "\n"
    s += bc.ReturnTrack.String() + "\n"
    return s
}

func (bc *BallClock) CycleBall() {
    ball := bc.ReturnTrack.Pop()
    if &ball == nil {
        // Error condition, throw exception or go equivalant
        fmt.Println("ERROR: ReturnTrack is empty. Create rolling ball clock with higher ball count.")
        os.Exit(1)
    }

    bc.MinuteTrack.Push(ball)
}

func (bc *BallClock) TrackAction() {
    if bc.MinuteTrack.IsTilted() {
        tiltBall := bc.MinuteTrack.Pop()
        bc.FiveMinuteTrack.Push(tiltBall)
        
        for ball := bc.MinuteTrack.Pop(); ball.Id != 0; ball = bc.MinuteTrack.Pop() {
            bc.ReturnTrack.Enqueue(ball)
        }
    }

    if bc.FiveMinuteTrack.IsTilted() {
        tiltBall := bc.FiveMinuteTrack.Pop()
        bc.HourTrack.Push(tiltBall)

        for ball := bc.FiveMinuteTrack.Pop(); ball.Id != 0; ball = bc.FiveMinuteTrack.Pop() {
            bc.ReturnTrack.Enqueue(ball)
        }

    }

    if bc.HourTrack.IsTilted() {
        tiltBall := bc.HourTrack.Pop()

        for ball := bc.HourTrack.Pop(); ball.Id != 0; ball = bc.HourTrack.Pop() {
            bc.ReturnTrack.Enqueue(ball)
        }
        
        bc.ReturnTrack.Enqueue(tiltBall)
    }
}

func (bc *BallClock) AdvanceMinute() {
    bc.CycleBall()
    bc.TrackAction()
}

func (bc *BallClock) RunForMinutes(minutes int) {
    for i := 1; i <= minutes; i++ {
        bc.AdvanceMinute()
    }

}

func (bc *BallClock) ReturnTrackMatchesOriginal() bool {
    matchesOriginal := true
    
    if bc.ReturnTrack.Groove.Len() != bc.BallCount {
        return false
    }

    element := bc.ReturnTrack.Groove.Front()
    for i := 1; i <= bc.BallCount; i++ {
        var ball Ball    
		if element != nil {
		    ball = element.Value.(Ball)
		}
        if (ball.Id != i) {
            matchesOriginal = false
        }
        element = element.Next()
    }
    return matchesOriginal    
}

func (bc *BallClock) MinutesToBallCycle() int {
    
    count := 0
    clockCycled := false
    for !clockCycled {
        bc.AdvanceMinute()
        count++
        if bc.ReturnTrackMatchesOriginal() {
            clockCycled = true
        }
    }
    return count;
}

func CreateBallClock(ballCount int) BallClock {
    var clock BallClock
    clock.BallCount = ballCount
    clock.MinuteTrack = BallTrack{Max:5, Name:"Minute Track", Groove:list.New()}
    clock.FiveMinuteTrack = BallTrack{Max:12, Name:"Five Minute Track", Groove:list.New()}
    clock.HourTrack = BallTrack{Max:12, Name:"Hour Track", Groove:list.New()}
    clock.ReturnTrack = BallTrack{Max:ballCount, Name:"Return Track", Groove:list.New()}

    for i := 1; i <= ballCount; i++ {
        clock.ReturnTrack.Enqueue(Ball{Id:i})
    }

    return clock
}

func BallOrderRepeat(totalBallCount int) {
    clock := CreateBallClock(totalBallCount)
    fmt.Println(clock.DisplayTracks())
    minsToCycle := clock.MinutesToBallCycle()
    daysToCycle := (float64)(minsToCycle) / (float64)(24 * 60)
    fmt.Print("Minutes to cycle: " + strconv.Itoa(minsToCycle))
    fmt.Println(", Days to cycle: " + strconv.FormatFloat(daysToCycle, 'f', 4, 64))
    fmt.Println(clock.DisplayTracks())    
}

func DisplayClockForMinutes(totalBallCount int, minutes int) {
    clock := CreateBallClock(totalBallCount)
    fmt.Println(clock.DisplayTracks())

    for i := 1; i <= minutes; i++ {
        clock.AdvanceMinute()
        fmt.Println("========== Minute " + strconv.Itoa(i) + " ===========")
        fmt.Println(clock.DisplayTracks())
    }   
}

func DisplayClockAfterMinutes(totalBallCount int, minutes int) {
    clock := CreateBallClock(totalBallCount)
    fmt.Println("========== Clock starting state ==========")
    fmt.Println(clock.DisplayTracks())

    for i := 1; i <= minutes; i++ {
        clock.AdvanceMinute()
    }
 
    fmt.Println(fmt.Sprintf("========== Clock ending state at minute, %d ===========", minutes))
    fmt.Println(clock.DisplayTracks())
}

// TODO: Custom MarshalJSON doesn't work, find better solution
func ClockStateAfterMinutes(totalBallCount int, minutes int) []byte {
    clock := CreateBallClock(totalBallCount)

    for i := 1; i <= minutes; i++ {
        clock.AdvanceMinute()
    }

    b, err := json.Marshal(clock)
    if err != nil {
        fmt.Println("error:", err)
    } 
    return b
}

func main() {
    fmt.Println("Rolling ball clock simulator.")
    BallOrderRepeat(27)
    BallOrderRepeat(45)
    BallOrderRepeat(127)
    DisplayClockAfterMinutes(30, 325)
}

