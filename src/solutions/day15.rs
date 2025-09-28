pub fn solve() {
    let input = std::fs::read_to_string("src/inputs/day15.txt").unwrap();
    part_1(&input);
    part_2(&input);
}

struct Disc {
    total_positions: usize,
    starting_position: usize,
}

impl Disc {
    fn parse(input: &str) -> Self {
        let input = input.strip_suffix('.').unwrap();
        let parts: Vec<&str> = input.split_whitespace().collect();
        match parts.as_slice() {
            [_, _, _, total, _, _, _, _, _, _, _, starting] => {
                let starting_position = starting.parse::<usize>().unwrap();
                let total_positions = total.parse::<usize>().unwrap();
                return Disc {
                    starting_position: starting_position,
                    total_positions: total_positions,
                };
            }
            _ => unreachable!(),
        }
    }
}

struct Sculpture {
    discs: Vec<Disc>,
}

impl Sculpture {
    fn parse(input: &str) -> Self {
        let mut discs: Vec<Disc> = Vec::new();

        for line in input.lines() {
            let disc = Disc::parse(line);
            discs.push(disc)
        }

        return Sculpture { discs };
    }
}

// TODO: I could have used the chinese reminder theorem to make this better in terms of time
// complexity, but this is much faster to code.
fn part_1(input: &str) {
    let mut time = 0;
    let sculpture = Sculpture::parse(input);

    'outer: loop {
        for (disc_number, disc) in sculpture.discs.iter().enumerate() {
            let current_position = time + disc.starting_position + disc_number + 1;

            if current_position.rem_euclid(disc.total_positions) != 0 {
                time += 1;
                continue 'outer;
            }
        }

        println!("part1: {}", time);
        break;
    }
}

fn part_2(input: &str) {
    let mut time = 0;
    let mut sculpture = Sculpture::parse(input);

    sculpture.discs.push(Disc {
        starting_position: 0,
        total_positions: 11,
    });

    'outer: loop {
        for (disc_number, disc) in sculpture.discs.iter().enumerate() {
            let current_position = time + disc.starting_position + disc_number + 1;

            if current_position.rem_euclid(disc.total_positions) != 0 {
                time += 1;
                continue 'outer;
            }
        }

        println!("part2: {}", time);
        break;
    }
}
