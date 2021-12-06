fn main() {
    let input: Vec<i64> = include_str!("../input")
        .replace("\n", "")
        .split(",")
        .map(|num| num.parse().unwrap())
        .collect();

    println!("Part 1: {}", part1(&input));
    println!("Part 2: {}", part2(&input));
}

fn part1(input: &Vec<i64>) -> i64 {
    procreate(input, 80)
}

fn part2(input: &Vec<i64>) -> i64 {
    procreate(input, 256)
}

fn procreate(input: &Vec<i64>, days: i32) -> i64 {
    let mut fish = vec![0; 9];

    for i in input {
        fish[*i as usize] += 1;
    }

    for _ in 0..days {
        fish.rotate_left(1);
        fish[6] += fish[8];
    }

    fish.iter().sum()
}
