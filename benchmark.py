from pgn_analyze import PGNAnalyzer
import cpuinfo
import time
import os


class benchmark:
    def __init__(self, game1, game2):
        self.game1 = game1
        self.game2 = game2

        directory = os.path.dirname(__file__)
        file_path = os.path.join(directory, "first_million.pgn")

        self.analyzer = PGNAnalyzer(file_path)

    def run_benchmark(self):
        start_time = time.perf_counter()

        obj = self.analyzer.get_games_in_range(self.game1, self.game2)

        end_time = time.perf_counter()

        elapsed_time = end_time - start_time

        print(
            f"Took {elapsed_time} seconds from game 1 to game {self.game2}")


if __name__ == '__main__':
    print(cpuinfo.get_cpu_info())
    game_ranges = [(1, 1000), (1, 10000), (1, 100000)]

    for i in game_ranges:
        obj = benchmark(i[0], i[1])
        obj.run_benchmark()


"""
{'python_version': '3.12.1.final.0 (64 bit)', 'cpuinfo_version': [9, 0, 0], 'cpuinfo_version_string': '9.0.0', 'arch': 'ARM_8', 'bits': 64, 'count': 8, 'arch_string_raw': 'arm64', 'brand_raw': 'Apple M1'}
Took 0.9084110420080833 seconds from game 1 to game 1000
Took 9.603569707949646 seconds from game 1 to game 10000
Took 114.05601770797512 seconds from game 1 to game 100000
"""
