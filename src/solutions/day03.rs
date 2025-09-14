use std::fs;

pub fn solve() {
    part_1();
    part_2();
}

fn part_1() {
    let content = fs::read_to_string("src/inputs/day03.txt").unwrap();

    let result = content
        .lines()
        .map(|line| {
            let side_lengths: Vec<i32> = line
                .trim()
                .split_whitespace()
                .map(|s| s.parse().unwrap())
                .collect();
            return side_lengths;
        })
        .filter(|side_lengths| is_triangle_possible(side_lengths))
        .count();

    println!("part1: {result}")
}

fn part_2() {
    let content = fs::read_to_string("src/inputs/day03.txt").unwrap();

    let rows: Vec<Vec<i32>> = content
        .lines()
        .map(|line| {
            line.split_whitespace()
                .map(|n| n.parse().unwrap())
                .collect()
        })
        .collect();

    let cols = rows[0].len();

    let mut result = 0;

    for col in 0..cols {
        let column: Vec<i32> = rows.iter().map(|row| row[col]).collect();
        for chunk in column.chunks(3) {
            if is_triangle_possible(chunk) {
                result += 1;
            }
        }
    }

    println!("part2: {result}")
}

fn is_triangle_possible(sides: &[i32]) -> bool {
    assert_eq!(sides.len(), 3);

    return sides[0] + sides[1] > sides[2]
        && sides[0] + sides[2] > sides[1]
        && sides[1] + sides[2] > sides[0];
}
