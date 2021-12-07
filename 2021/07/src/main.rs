fn main() {
    let input: Vec<i32> = include_str!("../input")
        .replace("\n", "")
        .split(',')
        .map(|num| num.parse().unwrap())
        .collect();

    println!("Part 1: {}", part1(&input));
    println!("Part 2: {}", part2(&input));
}

fn part1(input: &Vec<i32>) -> i32 {
    let mut result = i32::MAX;
    let min = *input.iter().min().unwrap();
    let max = *input.iter().max().unwrap();
    for target in min..max {
        let sum = input.iter().map(|x| (x - target).abs()).sum();
        result = std::cmp::min(result, sum);
    }
    result
}

fn part2(input: &Vec<i32>) -> i32 {
    let mut result = i32::MAX;
    let min = *input.iter().min().unwrap();
    let max = *input.iter().max().unwrap();
    for target in min..max {
        let sum = input
            .iter()
            .map(|x| {
                let distance = (x - target).abs();
                distance * (distance + 1) / 2
            })
            .sum();
        result = std::cmp::min(result, sum);
    }
    result
}
