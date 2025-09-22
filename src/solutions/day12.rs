pub fn solve() {
    let input = std::fs::read_to_string("src/inputs/day12.txt").unwrap();

    part_1(&input);
    part_2(&input);
}

fn part_1(input: &str) {
    let instructions = Instruction::parse(input);

    let mut cpu = CPU {
        registers: Registers {
            a: 0,
            b: 0,
            c: 0,
            d: 0,
        },
        pc: 0,
        instructions: instructions,
    };

    while (cpu.pc as usize) < cpu.instructions.len() {
        cpu.step();
    }

    println!("part1: {}", cpu.registers.a);
}

fn part_2(input: &str) {
    let instructions = Instruction::parse(input);

    let mut cpu = CPU {
        registers: Registers {
            a: 0,
            b: 0,
            c: 1,
            d: 0,
        },
        pc: 0,
        instructions: instructions,
    };

    while (cpu.pc as usize) < cpu.instructions.len() {
        cpu.step();
    }

    println!("part2: {}", cpu.registers.a);
}

#[derive(Copy, Clone)]
enum Register {
    A,
    B,
    C,
    D,
}

#[derive(Copy, Clone)]
enum Arg {
    Reg(Register),
    Val(isize),
}

#[derive(Copy, Clone)]
enum Instruction {
    Cpy(Arg, Register),
    Inc(Register),
    Dec(Register),
    Jnz(Arg, isize),
}

impl Instruction {
    fn is_register(str: &str) -> bool {
        return str == "a" || str == "b" || str == "c" || str == "d";
    }

    fn parse_register(str: &str) -> Register {
        debug_assert!(Self::is_register(str));
        match str {
            "a" => Register::A,
            "b" => Register::B,
            "c" => Register::C,
            "d" => Register::D,
            _ => unreachable!("unknwon register {str}"),
        }
    }

    fn parse_num(str: &str) -> isize {
        return str.parse::<isize>().unwrap();
    }

    fn parse_arg(str: &str) -> Arg {
        match Self::is_register(str) {
            true => return Arg::Reg(Self::parse_register(str)),
            false => return Arg::Val(Self::parse_num(str)),
        }
    }

    fn parse(input: &str) -> Vec<Instruction> {
        return input
            .lines()
            .map(|line| {
                let parts: Vec<&str> = line.split_whitespace().collect();

                match parts.as_slice() {
                    ["inc", x] => Instruction::Inc(Self::parse_register(x)),
                    ["dec", x] => Instruction::Dec(Self::parse_register(x)),
                    ["cpy", x, y] => Instruction::Cpy(Self::parse_arg(x), Self::parse_register(y)),
                    ["jnz", x, y] => Instruction::Jnz(Self::parse_arg(x), Self::parse_num(y)),
                    _ => unreachable!("unknown instruction {line}"),
                }
            })
            .collect();
    }
}

struct Registers {
    a: isize,
    b: isize,
    c: isize,
    d: isize,
}

struct CPU {
    registers: Registers,
    pc: isize,
    instructions: Vec<Instruction>,
}

impl CPU {
    fn arg_value(&self, arg: Arg) -> isize {
        match arg {
            Arg::Reg(register) => match register {
                Register::A => self.registers.a,
                Register::B => self.registers.b,
                Register::C => self.registers.c,
                Register::D => self.registers.d,
            },
            Arg::Val(value) => value,
        }
    }

    fn execute(&mut self, instruction: Instruction) -> isize {
        match instruction {
            Instruction::Cpy(x, y) => {
                let value = self.arg_value(x);
                match y {
                    Register::A => self.registers.a = value,
                    Register::B => self.registers.b = value,
                    Register::C => self.registers.c = value,
                    Register::D => self.registers.d = value,
                };
                return self.pc + 1;
            }
            Instruction::Inc(x) => {
                match x {
                    Register::A => self.registers.a += 1,
                    Register::B => self.registers.b += 1,
                    Register::C => self.registers.c += 1,
                    Register::D => self.registers.d += 1,
                };
                return self.pc + 1;
            }
            Instruction::Dec(x) => {
                match x {
                    Register::A => self.registers.a -= 1,
                    Register::B => self.registers.b -= 1,
                    Register::C => self.registers.c -= 1,
                    Register::D => self.registers.d -= 1,
                };
                return self.pc + 1;
            }
            Instruction::Jnz(x, y) => {
                let value = self.arg_value(x);
                match value {
                    0 => self.pc + 1,
                    _ => self.pc + y,
                }
            }
        }
    }

    fn step(&mut self) {
        let instruction = self.instructions[self.pc as usize];
        let next_pc = self.execute(instruction);
        self.pc = next_pc;
    }
}
