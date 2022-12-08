tree_map = [[int(tree) for tree in row] for row in open("input").read().splitlines()]

total = len(tree_map) * 2 + (len(tree_map[0]) - 2) * 2

for row in range(1, len(tree_map) - 1):
    for column in range(1, len(tree_map[0]) - 1):
        tallest_neighbours = [
            max(tree_map[row][0:column]),
            max(tree_map[row][column + 1 :]),
            max([row[column] for row in tree_map[0:row]]),
            max([row[column] for row in tree_map[row + 1 :]]),
        ]

        if not all(neighbour >= tree_map[row][column] for neighbour in tallest_neighbours):
            total += 1

print("part 1:", total)

scenic_scores = []
for row in range(0, len(tree_map)):
    for column in range(0, len(tree_map[0])):
        tree = tree_map[row][column]

        scenic_score = 1

        for i, r in enumerate(reversed(tree_map[row][0:column])):
            if r >= tree:
                scenic_score *= i+1
                break
        else:
            scenic_score *= len(tree_map[row][0:column])

        for i, r in enumerate(tree_map[row][column + 1 :]):
            if r >= tree:
                scenic_score *= i+1
                break
        else:
            scenic_score *= len(tree_map[row][column + 1 :])

        for i, r in enumerate(reversed([row[column] for row in tree_map[0:row]])):
            if r >= tree:
                scenic_score *= i+1
                break
        else:
            scenic_score *= len(tree_map[0:row])

        for i, r in enumerate([row[column] for row in tree_map[row + 1 :]]):
            if r >= tree:
                scenic_score *= i+1
                break
        else:
            scenic_score *= len(tree_map[row + 1 :])

        scenic_scores.append(scenic_score)

print("part 2:", max(scenic_scores))
