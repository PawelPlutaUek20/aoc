pub fn solve() {
    let input = std::fs::read_to_string("src/inputs/day18.txt").unwrap();
    println!("part1: {}", count_safe_tiles(&input.trim(), 40));
    println!("part2: {}", count_safe_tiles(&input.trim(), 400000));
}

fn count_safe_tiles(input: &str, rows: usize) -> usize {
    let mut curr = Tile::parse(input);
    let mut safe_count = count_safe_in_row(&curr);
    for _ in 1..rows {
        curr = next_row(curr);
        safe_count += count_safe_in_row(&curr);
    }
    return safe_count;
}

#[derive(Debug)]
enum Tile {
    Trap,
    Safe,
}

impl Tile {
    fn parse(input: &str) -> Vec<Self> {
        return input
            .chars()
            .map(|c| match c {
                '.' => Tile::Safe,
                '^' => Tile::Trap,
                _ => unreachable!(),
            })
            .collect();
    }

    fn is_safe(&self) -> bool {
        match self {
            Tile::Safe => true,
            Tile::Trap => false,
        }
    }
}

fn get_tile(left: Option<&Tile>, center: Option<&Tile>, right: Option<&Tile>) -> Tile {
    let tiles = (left, center, right);
    match tiles {
        (None, Some(Tile::Trap), Some(Tile::Trap)) => Tile::Trap,
        (None, Some(Tile::Safe), Some(Tile::Trap)) => Tile::Trap,
        (Some(Tile::Trap), Some(Tile::Trap), None) => Tile::Trap,
        (Some(Tile::Trap), Some(Tile::Safe), None) => Tile::Trap,
        (Some(Tile::Trap), Some(Tile::Trap), Some(Tile::Safe)) => Tile::Trap,
        (Some(Tile::Trap), Some(Tile::Safe), Some(Tile::Safe)) => Tile::Trap,
        (Some(Tile::Safe), Some(Tile::Trap), Some(Tile::Trap)) => Tile::Trap,
        (Some(Tile::Safe), Some(Tile::Safe), Some(Tile::Trap)) => Tile::Trap,
        _ => Tile::Safe,
    }
}

fn next_row(prev_row: Vec<Tile>) -> Vec<Tile> {
    return prev_row
        .iter()
        .enumerate()
        .map(|(idx, _)| {
            let left = if idx > 0 { prev_row.get(idx - 1) } else { None };
            let center = prev_row.get(idx);
            let right = prev_row.get(idx + 1);
            return get_tile(left, center, right);
        })
        .collect();
}

fn count_safe_in_row(tiles: &Vec<Tile>) -> usize {
    return tiles.iter().filter(|tile| tile.is_safe()).count();
}
