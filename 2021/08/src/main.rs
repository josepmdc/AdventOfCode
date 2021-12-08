struct Note {
    signals: Vec<String>,
    outputs: Vec<String>,
}

fn main() {
    let input: Vec<Note> = include_str!("../input.test")
        .lines()
        .map(|line| {
            line.split_once(" | ")
                .map(|part| Note {
                    signals: part.0.split_whitespace().map(|s| s.to_string()).collect(),
                    outputs: part.1.split_whitespace().map(|s| s.to_string()).collect(),
                })
                .unwrap()
        })
        .collect();

    println!("Part 1: {}", part1(&input));
}

fn part1(input: &Vec<Note>) -> i32 {
    let digits = [2, 3, 4, 7];

    input.iter().fold(0, |sum, note| {
        sum + note
            .outputs
            .iter()
            .filter(|x| {
                let len = x.len() as i32;
                digits.contains(&len)
            })
            .count() as i32
    })
}
