import chess.pgn
import os


class PGNAnalyzer:
    def __init__(self, file_path):
        self.file_path = file_path

    def get_games_in_range(self, start_game, end_game, start_offset=0):
        games = []
        current_game_num = 0

        with open(self.file_path, "r", encoding="utf-8") as file:
            file.seek(start_offset)

            while True:
                offset = file.tell()
                game = chess.pgn.read_game(file)
                if game is None:
                    break

                current_game_num += 1

                if current_game_num < start_game:
                    continue
                if current_game_num > end_game:
                    break

                games.append((current_game_num, offset, game))

        return games
