rounds = open("input").read().splitlines()

result_points = {
        "A Z": 0,
        "B X": 0,
        "C Y": 0,
        "A X": 3,
        "B Y": 3,
        "C Z": 3,
        "A Y": 6,
        "B Z": 6,
        "C X": 6,
}

move_points = {
    'X': 1,
    'Y': 2,
    'Z': 3,
}

# Part 1
total = sum(result_points[round] + move_points[round[-1]] for round in rounds) 
print("part 1: ", total)

# Part 2
opponent_moves = ['A', 'B', 'C']
moves = ['X', 'Y', 'Z']

total = 0
for round in rounds:
    (opponent, outcome) = round.split()

    if outcome == 'X':   # lose
        move = moves[opponent_moves.index(opponent) - 1]
        points = 0
    elif outcome == 'Y': # draw
        move = moves[opponent_moves.index(opponent)]
        points = 3
    else:                # win
        points = 6
        move = moves[(opponent_moves.index(opponent) + 1) % 3]

    total += points + move_points[move]

print("part 2: ", total)
