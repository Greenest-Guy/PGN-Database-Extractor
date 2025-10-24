from pgn_analyze import PGNAnalyzer
import time
import os

directory = os.path.dirname(__file__)
file_path = os.path.join(directory, "first_million.pgn")

analyzer = PGNAnalyzer(file_path)


start_time = time.perf_counter()

obj = analyzer.get_games_in_range(20000, 21000)

end_time = time.perf_counter()

elapsed_time = end_time - start_time


print(elapsed_time)
