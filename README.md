# PGN-Database-Extractor
Parse and extract chess games to CSV format, based on time control, skill group (ELO), and maximum ELO difference from large PGN files.

[![CodeFactor Grade](https://img.shields.io/codefactor/grade/github/Greenest-Guy/PGN-Database-Extractor?style=for-the-badge)](https://www.codefactor.io/repository/github/greenest-guy/PGN-Database-Extractor)

## How it Works

## Installation & Usage
```bash
git clone https://github.com/Greenest-Guy/PGN-Database-Extractor
```
1. Clone the repository
2. Edit the ```.env``` file setting ```PGN_PATH='/your/path/here'``` and ```OUTPUT_PATH='/your/path/here```.
3. Edit the parameters in ```/criteria/criteria.go```.
4. To change the definitions for skill groups, edit ```/skillgroups/skillgroups.go```.

## CPU Profile (10-Minute)
![profile](https://github.com/Greenest-Guy/PGN-Database-Extractor/blob/main/10min-CPU-profile-prod.png)
