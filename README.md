# Rolling Ball Clock
Simulation of the rolling ball clock as seen in https://youtu.be/NgcjJySoeTw. The
purpose is to provide a project that can be used to learn the basics of Go Programming 
Language. The goal is to provide RESTful web services that can allow clients to run
the rolling ball clock simulation in different ways. One may use may to calculate
how days before the balls return to their orginal order. Another use may be to run
an animation that displays in a web browser with the help of HTML5.

## Operation of the rolling ball clock

Every minute, the least recently used ball is removed from the queue of balls at the bottom
of the clock, elevated, then deposited on the minute indicator track, which is able to hold
four balls. When a fifth ball rolls on to the minute indicator track, its weight causes the 
track to tilt. The four balls already on the track run back down to join the queue of balls 
waiting at the bottom in reverse order of their original addition to the minutes track. The 
fifth ball, which caused the tilt, rolls on down to the five-minute indicator track. This 
track holds eleven balls. The twelfth ball carried over from the minutes causes the five-minute 
track to tilt, returning the eleven balls to the queue, again in reverse order of their addition.
The twelfth ball rolls down to the hour indicator. The hour indicator also holds eleven balls, 
but has one extra fixed ball which is always present so that counting the balls in the hour 
indicator will yield an hour in the range one to twelve. The twelfth ball carried over from 
the five-minute indicator causes the hour indicator to tilt, returning the eleven free balls 
to the queue, in reverse order, before the twelfth ball itself also returns to the queue.

