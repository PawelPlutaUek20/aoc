use std::{fs, str::FromStr};

pub fn solve() {
    part_1();
    part_2();
}

enum Direction {
    U,
    R,
    D,
    L,
}

impl FromStr for Direction {
    type Err = ();
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "U" => Ok(Direction::U),
            "R" => Ok(Direction::R),
            "D" => Ok(Direction::D),
            "L" => Ok(Direction::L),
            _ => unreachable!(),
        }
    }
}

struct Position(isize, isize);

fn part_1() {
    let content = fs::read_to_string("src/inputs/day02.txt").unwrap();

    let keypad = vec![
        vec![Some("1"), Some("2"), Some("3")],
        vec![Some("4"), Some("5"), Some("6")],
        vec![Some("7"), Some("8"), Some("9")],
    ];

    let mut position = Position(1, 1);
    let mut code: Vec<String> = vec![];

    content.lines().enumerate().for_each(|(i, line)| {
        code.resize(i + 1, String::from("0"));
        let chars: Vec<_> = line.trim().chars().collect();
        for c in chars {
            let direction = Direction::from_str(&c.to_string()).unwrap();
            let (dx, dy) = match direction {
                Direction::U => (0, -1),
                Direction::R => (1, 0),
                Direction::D => (0, 1),
                Direction::L => (-1, 0),
            };

            let next_position = Position(position.0 + dx, position.1 + dy);
            match get_key(&keypad, next_position.0, next_position.1) {
                Some(key) => {
                    code[i] = key.to_string();
                    position = next_position;
                }
                None => {}
            };
        }
    });

    let result = code.join("");
    println!("part1: {result}");
}

fn part_2() {
    let content = fs::read_to_string("src/inputs/day02.txt").unwrap();

    let keypad = vec![
        vec![None, None, Some("1"), None, None],
        vec![None, Some("2"), Some("3"), Some("4"), None],
        vec![Some("5"), Some("6"), Some("7"), Some("8"), Some("9")],
        vec![None, Some("A"), Some("B"), Some("C"), None],
        vec![None, None, Some("D"), None, None],
    ];

    let mut position = Position(0, 3);
    let mut code: Vec<String> = vec![];

    content.lines().enumerate().for_each(|(i, line)| {
        code.resize(i + 1, String::from("0"));
        let chars: Vec<_> = line.trim().chars().collect();
        for c in chars {
            let direction = Direction::from_str(&c.to_string()).unwrap();
            let (dx, dy) = match direction {
                Direction::U => (0, -1),
                Direction::R => (1, 0),
                Direction::D => (0, 1),
                Direction::L => (-1, 0),
            };

            let next_position = Position(position.0 + dx, position.1 + dy);
            match get_key(&keypad, next_position.0, next_position.1) {
                Some(key) => {
                    code[i] = key.to_string();
                    position = next_position;
                }
                None => {}
            };
        }
    });

    let result = code.join("");
    println!("part2: {result}");
}

fn get_key(keypad: &Vec<Vec<Option<&'static str>>>, x: isize, y: isize) -> Option<&'static str> {
    if y >= 0 && x >= 0 {
        keypad.get(y as usize)?.get(x as usize)?.as_ref().copied()
    } else {
        None
    }
}
