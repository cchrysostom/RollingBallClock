// rbclock_test
package rbclock

import (
	"fmt"
	"testing"
)

func TestBallOrderRepeat(t *testing.T) {
	totalBallCount := 27
	expectedMinutes := 33120
	
    clock := CreateBallClock(totalBallCount)
    fmt.Println(clock.DisplayTracks())
    minsToCycle := clock.MinutesToBallCycle()
	
	if minsToCycle != expectedMinutes {
		t.Errorf("Expected minutes to cycle to be %d. Actual was %d.", expectedMinutes, minsToCycle)
	}
}

func TestDisplayClockAfterMinutes(t *testing.T) {
	totalBallCount := 30
	minutes := 325
	
	expected := []struct{
		Id int
	}{
		{7},
		{3},
		{25},
		{13},
		{22},
	}
	
    clock := CreateBallClock(totalBallCount)

    for i := 1; i <= minutes; i++ {
        clock.AdvanceMinute()
    }
 
    if clock.FiveMinuteTrack.Groove.Len() != 5 {
		t.Errorf("Expected 5 Minute Track to have length of %d, found length %d.", 5, clock.FiveMinuteTrack.Groove.Len())
	}
	
	ballIndex := 0
    for ball := clock.FiveMinuteTrack.Groove.Front(); ball != nil; ball = ball.Next() {
		expectedValue := expected[ballIndex].Id
		actualValue := ball.Value.(Ball).Id
		if expectedValue != actualValue {
			t.Errorf("Expected ball value, %d. Found %d.", expectedValue, actualValue)
		}
		ballIndex++
	}

}