fn part1(input: &Vec<i32>) -> i32 {
    let mut prev = i32::MAX;
    let mut result = 0;

    for value in input {
        if *value > prev {
            result += 1;
        }
        prev = *value;
    }
    result
}

fn part1_fold(input: &Vec<i32>) -> i32 {
    input
        .iter()
        .fold((0, i32::MAX), |(sum, prev), current| {
            (if current > &prev { sum + 1 } else { sum }, *current)
        })
        .0
}

fn part2(input: &Vec<i32>) -> i32 {
    let mut prev = i32::MAX;
    let mut result = 0;

    for value in input.windows(3) {
        let curr: i32 = value.iter().sum();
        if curr > prev {
            result += 1;
        }
        prev = curr;
    }
    result
}

fn part2_fold(input: &Vec<i32>) -> i32 {
    input
        .windows(3)
        .map(|v| v.iter().sum())
        .fold((0, i32::MAX), |(sum, prev), current| {
            (if current > prev { sum + 1 } else { sum }, current)
        })
        .0
}

fn main() {
    let input: Vec<i32> = include_str!("../input")
        .lines()
        .map(|line| line.parse().unwrap())
        .collect();

    println!("Part 1:        {:?}", part1(&input));
    println!("Part 1 (fold): {:?}", part1_fold(&input));
    println!("Part 2:        {:?}", part2(&input));
    println!("Part 2 (fold): {:?}", part2_fold(&input));
}
