use std::fs::File;
use std::io::{BufRead, BufReader};

#[derive(Copy, Clone)]
struct Cell {
    value: i32,
    found: bool,
}

fn parse_file() -> (Vec<i32>, Vec<Vec<Vec<Cell>>>) {
    let reader = BufReader::new(File::open("inputs/d4").expect("cannot open 'input' file"));
    let mut boards: Vec<Vec<Vec<Cell>>> = Vec::new();
    let mut board: Vec<Vec<Cell>> = Vec::new();
    let mut runs: Vec<i32> = Vec::new();
    let mut cell: Cell;

    for (x, line) in reader.lines().enumerate() {
        let l: String = line.unwrap();
        if x == 0 {
            for run in l.split(",") {
                if run != "" {
                    runs.push(run.parse::<i32>().unwrap());
                }
            }
        } else {
            if l == "" {
                if board.len() > 0 {
                    boards.push(board.to_vec());
                    board = Vec::new();
                }
            } else {
                let mut line_board = Vec::new();
                for number in l.split_whitespace() {
                    if number != "" {
                        cell = Cell {
                            value: number.parse::<i32>().unwrap(),
                            found: false,
                        };
                        line_board.push(cell);
                    }
                }
                board.push(line_board.to_vec());
            }
        }
    }
    return (runs, boards);
}

fn check_win(board: &mut Vec<Vec<Cell>>) -> (i32, bool) {
    let mut total: i32 = 0;
    let mut found_by_column: Vec<i32> = vec![0, 0, 0, 0, 0];
    let mut found_line: bool = false;
    let mut found_column: bool = false;
    for (_, line_board) in board.iter().enumerate() {
        let mut found = 0;
        for (y, cell) in line_board.iter().enumerate() {
            if cell.found {
                found += 1;
                found_by_column[y] += 1;
            }
        }
        if found == 5 {
            found_line = true;
        }
    }
    for column in found_by_column {
        if column == 5 {
            found_column = true;
        }
    }
    if found_line || found_column {
        for (_, line) in board.iter().enumerate() {
            for (_, cell) in line.iter().enumerate() {
                if !cell.found {
                    total += cell.value;
                }
            }
        }
        return (total, true);
    }
    return (total, false);
}

fn print_board(board: Vec<Vec<Cell>>) {
    for line_board in board.iter() {
        for cell in line_board.iter() {
            print!("{}", if cell.found { "✅" } else { "❌" });
        }
        println!();
    }
}

fn mark_number(number: i32, board: &mut Vec<Vec<Cell>>)  {
    for line_board in board.iter_mut() {
        for cell in line_board.iter_mut() {
            if cell.value == number {
                (*cell).found = true;
            }
        }
    }
    print_board(board.to_vec());
}

pub fn day4_step1() {
    let  (runs, mut boards) = parse_file();
    let mut mut_boards = boards.iter_mut();
    for run in runs {
        // println!("{}", run);
        for mut board in mut_boards {
            mark_number(run, &mut board);
            let (value, found) = check_win(board);
            // println!("{}>{}", value, found);
            if found {
                println!("d04s1 |{}", value * run);
                return;
            }
        }
    }
}

pub fn day4_step2() {
    println!("d4s2 | not found")
}
