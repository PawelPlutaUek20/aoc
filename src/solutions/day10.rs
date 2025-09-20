use std::collections::HashMap;

pub fn solve() {
    let input = std::fs::read_to_string("src/inputs/day10.txt").unwrap();

    let commands = Command::parse(&input);
    let mut bots_map = BotsMap::new();

    for command in commands {
        match command {
            Command::AssignReceivers { i, l_bot, h_bot } => {
                bots_map.assign_receivers_to_bot(i, l_bot, h_bot);
            }
            Command::GiveValue { value, i } => {
                bots_map.give_value_to_bot(value, i);
            }
        }
    }

    let part_2 = bots_map
        .bots
        .keys()
        .filter(|key| **key == 1000 || **key == 1001 || **key == 1002)
        .fold(1, |acc, key| acc * bots_map.get_bot(*key).microchips[0]);

    println!("part2: {}", part_2)
}

struct Bot {
    microchips: Vec<u32>,

    low: Option<u32>,
    high: Option<u32>,
}

impl Bot {
    fn new() -> Self {
        return Bot {
            microchips: Vec::with_capacity(2),
            low: None,
            high: None,
        };
    }

    fn clear(&mut self) {
        self.microchips.clear();
        self.low = None;
        self.high = None;
    }
}

struct BotsMap {
    bots: HashMap<u32, Bot>,
}

impl BotsMap {
    fn new() -> Self {
        return BotsMap {
            bots: HashMap::new(),
        };
    }

    fn get_bot(&self, i: u32) -> &Bot {
        return self.bots.get(&i).unwrap();
    }

    fn get_mut_bot(&mut self, i: u32) -> &mut Bot {
        return self.bots.entry(i).or_insert(Bot::new());
    }

    fn give_value_to_bot(&mut self, value: u32, i: u32) {
        let bot = self.get_mut_bot(i);
        bot.microchips.push(value);
        self.update_bot(i);
    }

    fn assign_receivers_to_bot(&mut self, i: u32, l_bot: u32, h_bot: u32) {
        let bot = self.get_mut_bot(i);
        bot.low = Some(l_bot);
        bot.high = Some(h_bot);
        self.update_bot(i);
    }

    fn update_bot(&mut self, i: u32) {
        let bot = self.get_mut_bot(i);

        if bot.microchips.len() != 2 {
            return;
        }

        if bot.high.is_none() || bot.low.is_none() {
            return;
        }

        bot.microchips.sort();

        let low = bot.microchips[0];
        let high = bot.microchips[1];
        let l_bot = bot.low.unwrap();
        let h_bot = bot.high.unwrap();

        if low == 17 && high == 61 {
            println!("part1: {}", i);
        }

        bot.clear();
        self.give_value_to_bot(low, l_bot);
        self.give_value_to_bot(high, h_bot);
    }
}

enum Command {
    AssignReceivers { i: u32, l_bot: u32, h_bot: u32 },
    GiveValue { value: u32, i: u32 },
}

impl Command {
    fn parse_bin(bin: &str, bin_str: &str) -> u32 {
        let i = bin_str.parse::<u32>().unwrap();
        if bin == "output" {
            return i + 1000;
        } else {
            return i;
        }
    }

    fn parse(input: &str) -> Vec<Self> {
        let mut commands = Vec::new();

        for line in input.lines() {
            let parts: Vec<&str> = line.split_whitespace().collect();

            match parts.as_slice() {
                ["bot", i, _, _, _, bin1, bin_l, _, _, _, bin2, bin_h] => {
                    commands.push(Self::AssignReceivers {
                        i: Self::parse_bin("bot", i),
                        l_bot: Self::parse_bin(bin1, bin_l),
                        h_bot: Self::parse_bin(bin2, bin_h),
                    });
                }
                ["value", value, _, _, bin, i] => {
                    commands.push(Self::GiveValue {
                        value: value.parse::<u32>().unwrap(),
                        i: Self::parse_bin(bin, i),
                    });
                }
                _ => unreachable!(),
            }
        }

        return commands;
    }
}
