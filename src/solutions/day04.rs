use std::{collections::HashMap, fs};

struct Room {
    encrypted_name: String,
    sector_id: u32,
    checksum: String,
}

impl Room {
    fn parse(input: &str) -> Option<Self> {
        let bracket_start = input.find('[')?;
        let bracket_end = input.find(']')?;
        let last_dash = input[..bracket_start].rfind('-')?;

        Some(Room {
            encrypted_name: input[..last_dash].to_string(),
            sector_id: input[last_dash + 1..bracket_start].parse().ok()?,
            checksum: input[bracket_start + 1..bracket_end].to_string(),
        })
    }

    fn is_real(&self) -> bool {
        let mut letter_counts = HashMap::new();
        for ch in self.encrypted_name.chars() {
            if ch != '-' {
                *letter_counts.entry(ch).or_insert(0) += 1
            }
        }

        let mut letters: Vec<char> = letter_counts.keys().copied().collect();
        letters.sort_by(|&a, &b| {
            let count_a = letter_counts[&a];
            let count_b = letter_counts[&b];

            return count_b.cmp(&count_a).then(a.cmp(&b));
        });

        let expected_checksum: String = letters.into_iter().take(5).collect();
        self.checksum == expected_checksum
    }

    fn decrypt_name(&self) -> String {
        return self
            .encrypted_name
            .chars()
            .map(|ch| {
                if ch == '-' {
                    return ' ';
                }

                let start = 'a' as u32;
                let end = 'z' as u32;
                let count = end - start + 1;

                let pos_old = ch as u32;
                let pos_new = (pos_old - start + self.sector_id).rem_euclid(count) + start;
                return pos_new as u8 as char;
            })
            .collect();
    }
}

pub fn solve() {
    part_1();
    part_2();
}

fn part_1() {
    let content = fs::read_to_string("src/inputs/day04.txt").unwrap();

    let result: u32 = content
        .lines()
        .map(|line| Room::parse(line).unwrap())
        .filter(|room| room.is_real())
        .map(|room| room.sector_id)
        .sum();

    println!("part1: {result}")
}

fn part_2() {
    let content = fs::read_to_string("src/inputs/day04.txt").unwrap();

    let result = content
        .lines()
        .map(|line| Room::parse(line).unwrap())
        .filter(|room| room.is_real())
        .find(|room| room.decrypt_name() == "northpole object storage")
        .map(|room| room.sector_id)
        .unwrap();

    println!("part2: {result}")
}
