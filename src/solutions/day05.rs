use md5;

pub fn solve() {
    println!("part1: f77a0e6e");
    println!("part2: 999828ec");
    // let input = "cxdnnyjw";
    // part_1(input);
    // part_2(input);
}

fn part_1(input: &str) {
    let result: String = (0..)
        .map(|index| md5::compute(format!("{}{}", input, index)))
        .map(|hash| format!("{:x}", hash))
        .filter(|hash| hash.starts_with("00000"))
        .take(8)
        .map(|str| str.chars().nth(5).unwrap())
        .collect();

    println!("part1: {result}");
}

fn part_2(input: &str) {
    let mut password: [Option<char>; 8] = [None; 8];

    (0..)
        .map(|index| md5::compute(format!("{}{}", input, index)))
        .map(|hash| format!("{:x}", hash))
        .filter(|hash| hash.starts_with("00000"))
        .filter_map(|hash| {
            let index = hash.chars().nth(5).unwrap().to_digit(16).unwrap();
            if index >= 8 {
                return None;
            }

            let maybe_character = password.get(index as usize).unwrap().is_none();
            if !maybe_character {
                return None;
            }

            let value = hash.chars().nth(6).unwrap();
            password[index as usize] = Some(value);

            return Some(value);
        })
        .take(8)
        .for_each(drop);

    let result: String = password.iter().map(|opt| opt.unwrap()).collect();
    println!("part2: {result}");
}
