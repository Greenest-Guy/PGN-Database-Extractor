# PGN-Database-Extractor
Parse and extract chess games to CSV format, based on time control, skill group (ELO), and maximum ELO difference from large PGN files.

[![CodeFactor Grade](https://img.shields.io/codefactor/grade/github/Greenest-Guy/PGN-Database-Extractor?style=for-the-badge)](https://www.codefactor.io/repository/github/greenest-guy/PGN-Database-Extractor)

## DISCLAIMER
This project was developed as part of a research project and is not intended to serve as a universal or production-ready tool. Its primary purpose is to support the specific goals and
methodology of that research. Anyone is welcome to use, modify, or learn from this code.

## Installation & Usage
```
git clone https://github.com/Greenest-Guy/PGN-Database-Extractor
```
1. Clone the repository
2. Create a ```.env``` file setting ```PGN_PATH='/your/path/here'``` and ```OUTPUT_PATH='/your/path/here```.
3. Edit the parameters in ```/criteria/criteria.go```.
4. To change the definitions for skill group and time control, edit ```/skillgroups/skillgroups.go``` and ```/timecontrols/timecontrols.go``` respectively.

## CPU Profile (10-Minute)
![profile](https://github.com/Greenest-Guy/PGN-Database-Extractor/blob/main/10min-CPU-profile-prod.png)
