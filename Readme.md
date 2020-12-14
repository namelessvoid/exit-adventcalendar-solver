# EXIT Advent Calendar Solver

This small program can solve the EXIT advent calender released by German publisher Kosmos.

Disclaimer: This is just a project just for fun. Do not use it to really solve the calendar as you will miss the fun of the game :)

## About the game

The EXIT advent calendar is not a usual advent calendar. Only the first door is numbered up front. To find the next door, you have to solve a riddle a day. The solution will guide you to the door you have to open on the next day.

The solution to each riddle is a three-digit code with numbers between 0 and 9. These numbers are put into a "decoder board" with three colored stripes. After entering the code, the decoder board is flipped. On the back, you can see an arrow and an icon on each stripe.

By starting from the current door and moving from door to door according to the arrows, you will find the door for the next day. To verify the solution, each door has three icons printed on it. These three icons have to match the icons on the back of the decoder board. If they match, the solution is valid. If not, you ahve to try another three-digit code.

## About the solver

The solver exploits that for each day, only one arrow-and-icon combination is valid (i.e. the solution for each day is unique, obviously). A depth-first search is performed which performs:

```
0. Push the start door to the calendar solution.
1. Iterate through all the doors
   1.1 Put the three icons on the door into the decoder board.
   1.2 Get the arrows from the decoder board
   1.3 Starting from the current door, move along the arrows. If it ends on the current door, this door is the solution for the current day.
      1.3.1 Append the currently examined door to the calendar solution
      1.3.2 Continue with 1.
```

## About the program

The program has to be called with parameter `-day` which is a uint and determines the day up to which you want to solve the calendar. 

Example usage: `go run cmd/solver/* -day 13`

Output (faked to not spoil you):

```
  3  ??   9   5   7  ?? 
 ??   6  ??  ??   4  ?? 
 12  13  ??  10   8  ?? 
  1   2  11  ??  ??  ?? 
```

Doors which are not solved yet are marked with `??`

## About the code

The executable lives in `cmd/solver`. The gameplay related components (calendar, decoder board etc) are in `internal/pkg/calendar`. The depth-first solver is implemented in `internal/pkg/solver`.