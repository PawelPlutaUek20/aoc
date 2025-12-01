use std::fs;

pub fn solve() {
    part_1();
}

fn part_1() {
    let content = fs::read_to_string("src/inputs/day01.txt").unwrap();

    let mut part1 = 0;
    let mut part2 = 0;
    let mut pointing = 50;

    for line in content.lines() {
        let dir = &line[0..1];

        let count = &line[1..];
        let count = count.parse::<i32>().unwrap();

        let dir = match dir {
            "L" => -1,
            "R" => 1,
            _ => unreachable!(),
        };

        for _ in 0..count {
            pointing = ((pointing + 1 * dir) as i32).rem_euclid(100);
            if pointing == 0 {
                part2 += 1;
            }
        }
        if pointing == 0 {
            part1 += 1;
        }
    }

    println!("part 1: {}", part1);
    println!("part 2: {}", part2);
}
