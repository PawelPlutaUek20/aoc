pub fn solve() {
    let input = "11100010111110100";
    println!("part1: {}", get_checksum(input, 272));
    println!("part2: {}", get_checksum(input, 35651584));
}

fn get_checksum(input: &str, length: usize) -> String {
    let mut initial_state = input.to_string();

    while initial_state.len() < length {
        initial_state = step(initial_state);
    }

    let final_state: String = initial_state.chars().take(length).collect();

    let mut checksum = compute_checksum(final_state);
    while checksum.len() & 1 == 0 {
        checksum = compute_checksum(checksum)
    }

    return checksum;
}

fn step(a: String) -> String {
    let b = a
        .chars()
        .rev()
        .map(|c| if c == '0' { '1' } else { '0' })
        .collect::<String>();
    let b = &b;

    return a + "0" + b;
}

fn compute_checksum(a: String) -> String {
    a.chars()
        .collect::<Vec<char>>()
        .chunks(2)
        .map(|chunk| {
            if chunk[0] == chunk[1] {
                return '1';
            } else {
                return '0';
            }
        })
        .collect()
}
