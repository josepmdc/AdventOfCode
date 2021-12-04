fn main() {
    let input = include_str!("../input");
    let numbers: Vec<i32> = input
        .lines()
        .nth(0)
        .unwrap()
        .split(',')
        .map(|num| num.parse().unwrap())
        .collect();

    let boards: Vec<Vec<Vec<(i32, bool)>>> =
        input.lines().skip(1).fold(Vec::new(), |mut boards, line| {
            if line.trim().is_empty() {
                boards.push(Vec::new());
            } else {
                boards.last_mut().unwrap().push(
                    line.split_whitespace()
                        .map(|num| (num.parse().unwrap(), false))
                        .collect(),
                );
            }
            boards
        });

    println!("Part 1 Bingo! => {}", part1(boards.clone(), numbers.clone()));
    println!("Part 2 Bingo! => {}", part2(boards, numbers));
}

fn part1(mut boards: Vec<Vec<Vec<(i32, bool)>>>, numbers: Vec<i32>) -> i32 {
    for drawn_number in numbers {
        for i in 0..boards.len() {
            for row in 0..boards[i].len() {
                let mut row_bingo = true;
                for col in 0..boards[i][row].len() {
                    if boards[i][row][col].0 == drawn_number {
                        boards[i][row][col].1 = true;
                    }
                    if !boards[i][row][col].1 {
                        row_bingo = false;
                    }
                }
                if row_bingo {
                    return bingo_score(&boards[i], drawn_number);
                }
            }
            if column_bingo(&boards[i]) {
                return bingo_score(&boards[i], drawn_number);
            }
        }
    }
    0
}

fn part2(mut boards: Vec<Vec<Vec<(i32, bool)>>>, numbers: Vec<i32>) -> i32 {
    let mut score = 0;
    let mut completed_boards = vec![false; boards.len()];
    for drawn_number in numbers {
        for i in 0..boards.len() {
            if completed_boards[i] {
                if completed_boards.iter().all(|v| *v == true) {
                    return score;
                }
                continue;
            }

            for row in 0..boards[i].len() {
                let mut row_bingo = true;
                for col in 0..boards[i][row].len() {
                    if boards[i][row][col].0 == drawn_number {
                        boards[i][row][col].1 = true;
                    }
                    if !boards[i][row][col].1 {
                        row_bingo = false;
                    }
                }
                if row_bingo {
                    completed_boards[i] = true;
                    score = bingo_score(&boards[i], drawn_number);
                }
            }
            if column_bingo(&boards[i]) {
                completed_boards[i] = true;
                score = bingo_score(&boards[i], drawn_number);
            }
        }
    }
    score
}

fn bingo_score(board: &Vec<Vec<(i32, bool)>>, drawn_number: i32) -> i32 {
    let mut sum = 0;
    for row in 0..board.len() {
        for col in 0..board[row].len() {
            let num = board[row][col];
            if !num.1 {
                sum += num.0;
            }
        }
    }
    sum * drawn_number
}

fn column_bingo(board: &Vec<Vec<(i32, bool)>>) -> bool {
    for col in 0..board[0].len() {
        let mut col_bingo = true;
        for row in 0..board.len() {
            if !board[row][col].1 {
                col_bingo = false;
                break;
            }
        }
        if col_bingo {
            return true;
        }
    }
    false
}
