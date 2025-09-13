use std::fs;
use std::str::FromStr;

#[derive(Clone, Copy)]
enum Direction {
    L,
    R,
}

#[derive(Clone, Copy)]
enum Orientation {
    N,
    E,
    S,
    W,
}

impl Orientation {
    fn to_int(self) -> usize {
        match self {
            Orientation::N => 0,
            Orientation::E => 1,
            Orientation::S => 2,
            Orientation::W => 3,
        }
    }

    fn from_int(index: usize) -> Self {
        match index % 4 {
            0 => Orientation::N,
            1 => Orientation::E,
            2 => Orientation::S,
            3 => Orientation::W,
            _ => unreachable!(),
        }
    }

    fn turn(self, direction: Direction) -> Self {
        let current = self.to_int() as i32;
        let delta = match direction {
            Direction::R => 1,
            Direction::L => -1,
        };
        let next = (current + delta).rem_euclid(4);
        Self::from_int(next as usize)
    }
}

impl FromStr for Direction {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "L" => Ok(Direction::L),
            "R" => Ok(Direction::R),
            _ => unreachable!(),
        }
    }
}

struct Instruction {
    dir: Direction,
    steps: isize,
}

struct Position {
    x: isize,
    y: isize,
    orientation: Orientation,
}

pub fn solve() {
    part_1();
    part_2();
}

fn part_1() {
    let content = fs::read_to_string("src/inputs/day01.txt").unwrap();

    let instructions = content.trim().split(", ").map(|e| {
        let (dir_str, steps_str) = e.split_at(1);

        let dir = Direction::from_str(dir_str).unwrap();
        let steps = steps_str.parse().unwrap();
        return Instruction { dir, steps };
    });

    let initial_position = Position {
        x: 0,
        y: 0,
        orientation: Orientation::N,
    };

    let target_position: Position = instructions.fold(initial_position, |acc, e| {
        let next_orientation: Orientation = acc.orientation.turn(e.dir);

        let (dx, dy) = match next_orientation {
            Orientation::N => (0, 1),
            Orientation::S => (0, -1),
            Orientation::E => (1, 0),
            Orientation::W => (-1, 0),
        };

        return Position {
            x: acc.x + (dx * e.steps),
            y: acc.y + (dy * e.steps),
            orientation: next_orientation,
        };
    });

    let result = target_position.x.abs() + target_position.y.abs();
    println!("part1: {result}");
}

fn part_2() {
    let content = fs::read_to_string("src/inputs/day01.txt").unwrap();

    let instructions = content.trim().split(", ").map(|e| {
        let (dir_str, steps_str) = e.split_at(1);

        let dir = Direction::from_str(dir_str).unwrap();
        let steps = steps_str.parse().unwrap();
        return Instruction { dir, steps };
    });

    let mut current_position = Position {
        x: 0,
        y: 0,
        orientation: Orientation::N,
    };

    let mut visited = vec![];
    visited.push((current_position.x, current_position.y));

    let mut maybe_target_position: Option<Position> = None;

    'outer: for e in instructions {
        let next_orientation = current_position.orientation.turn(e.dir);

        for _ in 0..e.steps {
            let (dx, dy) = match next_orientation {
                Orientation::N => (0, 1),
                Orientation::S => (0, -1),
                Orientation::E => (1, 0),
                Orientation::W => (-1, 0),
            };

            let next = Position {
                x: current_position.x + dx,
                y: current_position.y + dy,
                orientation: next_orientation,
            };

            if visited.contains(&(next.x, next.y)) {
                maybe_target_position = Some(next);
                break 'outer;
            }

            visited.push((next.x, next.y));
            current_position = next;
        }
    }

    let target_position = maybe_target_position.unwrap();
    let result = target_position.x.abs() + target_position.y.abs();
    println!("part2: {result}")
}
