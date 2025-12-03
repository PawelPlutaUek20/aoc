use std::fs;

pub fn solve() {
    let content = fs::read_to_string("src/inputs/day03.txt").unwrap();

    let mut part1 = 0;
    let mut part2 = 0;

    for line in content.lines() {
        part1 += output_joltage(line, 2);
        part2 += output_joltage(line, 12);
    }

    println!("part1: {}", part1);
    println!("part2: {}", part2);
}

fn output_joltage(bank: &str, batteries: u32) -> u64 {
    if batteries == 0 {
        return 0;
    }

    let digits_count = batteries - 1;

    let (idx, digit) = bank
        .chars()
        .take(bank.len() - digits_count as usize)
        .map(|char| char.to_digit(10).unwrap() as u64)
        .enumerate()
        .max_by(|(i1, a), (i2, b)| a.cmp(b).then(i2.cmp(i1)))
        .unwrap();

    return digit * 10_u64.pow(digits_count) + output_joltage(&bank[idx + 1..], batteries - 1);
}
