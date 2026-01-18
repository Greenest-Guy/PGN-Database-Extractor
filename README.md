# PGN-Database-Extractor
Parse and extract chess games into a CSV based on time control, skill group (ELO), and maximum ELO difference from large PGN files.

[![CodeFactor Grade](https://img.shields.io/codefactor/grade/github/Greenest-Guy/PGN-Database-Extractor?style=for-the-badge)](https://www.codefactor.io/repository/github/greenest-guy/PGN-Database-Extractor)

## Parser Design & Guarantees

- CSV formatted as ```WhiteElo,BlackElo,TimeControl,Moves```, where ```TimeControl``` is in the X+Y format, and ```Moves``` are a single string as formated in the PGN.

- No move validation, games are assumed to only contain legal moves.
  
- Time controls are based on their [Lichess definitions](https://lichess.org/faq#time-controls).

- Games are scanned individually for RAM safety.

## Example
### Parameters
```Go
var (
	TimeControl = "Bullet"
	SkillGroup  = "Expert"
	MaxEloDiff  = 200
	FileName    = fmt.Sprintf("%s_%s_%d.csv", TimeControl, SkillGroup, MaxEloDiff)
	NumGames    = 1000000 // Lichess monthly dump - 2025 October 29.9 GB 91,549,148
)
```
### PGN Game
```
[Event "Rated Bullet game"]
[Site "https://lichess.org/2WHJ7Fyz"]
[Date "2020.01.01"]
[Round "-"]
[White "tristanoff"]
[Black "blue-night"]
[Result "0-1"]
[UTCDate "2020.01.01"]
[UTCTime "00:00:00"]
[WhiteElo "2420"]
[BlackElo "2269"]
[WhiteRatingDiff "-9"]
[BlackRatingDiff "+9"]
[ECO "B18"]
[Opening "Caro-Kann Defense: Classical Variation, Main Line"]
[TimeControl "30+0"]
[Termination "Normal"]

1. e4 { [%clk 0:00:30] } c6 { [%clk 0:00:30] } 2. d4 { [%clk 0:00:29] } d5 { [%clk 0:00:30] } 3. Nc3 { [%clk 0:00:29] } dxe4 { [%clk 0:00:29] } 4. Nxe4 { [%clk 0:00:29] } Bf5 { [%clk 0:00:29] } 5. Ng3 { [%clk 0:00:29] } Bg6 { [%clk 0:00:29] } 6. h4 { [%clk 0:00:29] } h6 { [%clk 0:00:28] } 7. h5 { [%clk 0:00:28] } Bh7 { [%clk 0:00:28] } 8. Nf3 { [%clk 0:00:28] } Nd7 { [%clk 0:00:27] } 9. Bd3 { [%clk 0:00:28] } Bxd3 { [%clk 0:00:25] } 10. Qxd3 { [%clk 0:00:28] } Qc7 { [%clk 0:00:25] } 11. Bf4 { [%clk 0:00:27] } Qxf4 { [%clk 0:00:23] } 12. Ne2 { [%clk 0:00:27] } Qc7 { [%clk 0:00:23] } 13. O-O-O { [%clk 0:00:27] } e6 { [%clk 0:00:23] } 14. Qd2 { [%clk 0:00:26] } Ngf6 { [%clk 0:00:23] } 15. Nf4 { [%clk 0:00:24] } Bd6 { [%clk 0:00:22] } 16. Qe3 { [%clk 0:00:24] } Bxf4 { [%clk 0:00:21] } 17. Kb1 { [%clk 0:00:24] } Bxe3 { [%clk 0:00:21] } 18. fxe3 { [%clk 0:00:24] } O-O-O { [%clk 0:00:21] } 19. Rhe1 { [%clk 0:00:24] } e5 { [%clk 0:00:20] } 20. Nd2 { [%clk 0:00:24] } exd4 { [%clk 0:00:19] } 21. exd4 { [%clk 0:00:24] } Rhe8 { [%clk 0:00:19] } 22. Nb3 { [%clk 0:00:24] } Rxe1 { [%clk 0:00:19] } 23. Rxe1 { [%clk 0:00:23] } Re8 { [%clk 0:00:18] } 24. Rd1 { [%clk 0:00:23] } c5 { [%clk 0:00:17] } 25. dxc5 { [%clk 0:00:23] } Nxc5 { [%clk 0:00:16] } 26. Nd4 { [%clk 0:00:22] } Rd8 { [%clk 0:00:15] } 27. Kc1 { [%clk 0:00:21] } Ne6 { [%clk 0:00:15] } 28. Nb5 { [%clk 0:00:19] } Rxd1+ { [%clk 0:00:15] } 29. Kxd1 { [%clk 0:00:19] } Qd7+ { [%clk 0:00:13] } 0-1
```
### CSV Output of Game Above
```CSV
WhiteElo,BlackElo,TimeControl,Moves
2420,2269,30+0,1. e4 { [%clk 0:00:30] } c6 { [%clk 0:00:30] } 2. d4 { [%clk 0:00:29] } d5 { [%clk 0:00:30] } 3. Nc3 { [%clk 0:00:29] } dxe4 { [%clk 0:00:29] } 4. Nxe4 { [%clk 0:00:29] } Bf5 { [%clk 0:00:29] } 5. Ng3 { [%clk 0:00:29] } Bg6 { [%clk 0:00:29] } 6. h4 { [%clk 0:00:29] } h6 { [%clk 0:00:28] } 7. h5 { [%clk 0:00:28] } Bh7 { [%clk 0:00:28] } 8. Nf3 { [%clk 0:00:28] } Nd7 { [%clk 0:00:27] } 9. Bd3 { [%clk 0:00:28] } Bxd3 { [%clk 0:00:25] } 10. Qxd3 { [%clk 0:00:28] } Qc7 { [%clk 0:00:25] } 11. Bf4 { [%clk 0:00:27] } Qxf4 { [%clk 0:00:23] } 12. Ne2 { [%clk 0:00:27] } Qc7 { [%clk 0:00:23] } 13. O-O-O { [%clk 0:00:27] } e6 { [%clk 0:00:23] } 14. Qd2 { [%clk 0:00:26] } Ngf6 { [%clk 0:00:23] } 15. Nf4 { [%clk 0:00:24] } Bd6 { [%clk 0:00:22] } 16. Qe3 { [%clk 0:00:24] } Bxf4 { [%clk 0:00:21] } 17. Kb1 { [%clk 0:00:24] } Bxe3 { [%clk 0:00:21] } 18. fxe3 { [%clk 0:00:24] } O-O-O { [%clk 0:00:21] } 19. Rhe1 { [%clk 0:00:24] } e5 { [%clk 0:00:20] } 20. Nd2 { [%clk 0:00:24] } exd4 { [%clk 0:00:19] } 21. exd4 { [%clk 0:00:24] } Rhe8 { [%clk 0:00:19] } 22. Nb3 { [%clk 0:00:24] } Rxe1 { [%clk 0:00:19] } 23. Rxe1 { [%clk 0:00:23] } Re8 { [%clk 0:00:18] } 24. Rd1 { [%clk 0:00:23] } c5 { [%clk 0:00:17] } 25. dxc5 { [%clk 0:00:23] } Nxc5 { [%clk 0:00:16] } 26. Nd4 { [%clk 0:00:22] } Rd8 { [%clk 0:00:15] } 27. Kc1 { [%clk 0:00:21] } Ne6 { [%clk 0:00:15] } 28. Nb5 { [%clk 0:00:19] } Rxd1+ { [%clk 0:00:15] } 29. Kxd1 { [%clk 0:00:19] } Qd7+ { [%clk 0:00:13] } 0-1
```
## Installation & Usage
```bash
git clone https://github.com/Greenest-Guy/PGN-Database-Extractor
```
1. Clone the repository
2. Edit the ```.env``` file setting ```PGN_PATH='/your/path/here'``` and ```OUTPUT_PATH='/your/path/here```.
3. Edit the parameters in ```/criteria/criteria.go```.
4. To change the definitions for skill groups, edit ```/skillgroups/skillgroups.go```.
5. Run the program with ```go run main.go``` in an integrated terminal.

## CPU Profile (10-Minute)
![profile](https://github.com/Greenest-Guy/PGN-Database-Extractor/blob/main/10min-CPU-profile-prod.png)
