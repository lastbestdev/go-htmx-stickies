# stickies
a minimalist tool for adding sticky notes to a board.

## tech
htmx + go (using `templ` for components) 

## features
- single user support
    - create board (user names the board)
    - create stickies (user enters note content and chooses a color)
    - delete stickies

## working on
- [x] feature: add database 
- [x] styling: menu and boards pages 
- [ ] feature: form validation
- [ ] styling: create board and sticky note forms
- [ ] bug: single form submit creates multiple boards/stickies (expected: single submit should create single sticky/board)
- [ ] bug: no navigation after delete board (expected: should navigate back to menu)

## demo (updated 10/17/25)
![best basketball player of all time](stickies-demo-2.gif)