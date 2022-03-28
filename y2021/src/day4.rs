use std::fs::File;
use std::io::{BufRead, BufReader};
use std::str::Split;

struct Cell {
    value: i32,
    found: bool,
}

pub fn day4_step1() {
    let reader = BufReader::new(File::open("inputs/d4").expect("cannot open 'input' file"));
    let mut runs: Split<&str>;
    let mut games = vec![];
    let mut rounds = vec![];
    for (x, line) in reader.lines().enumerate() {
        let l = line.unwrap();
        if x == 0 {
            runs = l.split(",");
            for run in runs {
                if run != "" {
                    rounds.push(run.parse::<i32>().unwrap());
                }
            }
        } else {
            let mut game = vec![];
            if l == "" {
                games.push(game);
            } else {
                let mut line_game = vec![];
                for number in l.split(" ") {
                    if number != "" {
                        let cell = Cell {
                            value: number.parse::<i32>().unwrap(),
                            found: false,
                        };
                        line_game.push(cell);
                    }
                }
                println!("push here {}", line_game.len());
                game.push(line_game);
                println!("push :: {}", game[game.len() - 1].len());
            }
            println!("push > {}", game[game.len() - 1].len());
        }
    }

    for game in games {
        // println!("game ! {}", game.len());
        for line in game {
            println!("super ! {}", line.len());
            if line.len() != 5 {
                println!("game line error");
            }
            for c in line {
                println!("{} > {}", c.found, c.value);
            }
        }
    }
    // for (x, round) in rounds.iter().enumerate() {
    //     let continue_game = true;
    //     update_games(*round, &mut games);
    //     let (game_id, found) = check_games(&mut games);
    //     println!("winner game id : {}, found {}", game_id, found);
    //     // for game in games {
    //     //     update_game(*round, game);
    //     // }
    //     // let continue_game = run_round(*round, &mut *games);
    //     if !continue_game {
    //         println!("game complete at round {}", x);
    //         break;
    //     }
    // }

    println!("d4s1 | not found")
}

// fn update_games(round: i32, games: &mut Vec<Vec<Vec<Cell>>>) {
//     println!("games len {}", games.len());
//     for game in games {
//         for line in game {
//             for x in 0..line.len() {
//                 if line[x].value == round {
//                     line[x].found = true;
//                 }
//             }
//         }
//     }
// }
// fn check_games(games: &mut Vec<Vec<Vec<Cell>>>) -> (usize, bool) {
//     let width = 5;
//     let height = 5;

//     for (id, game) in games.iter().enumerate() {
//         println!("Testing game {} | {}", id, game.len());
//         let mut col_ok = vec![true, true, true, true, true];
//         for x in 0..height {
//             let mut row_ok = true;
//             for y in 0..width {
//                 if !game[x][y].found {
//                     col_ok[y] = false;
//                     row_ok = false;
//                 }
//             }
//             if row_ok {
//                 return (id, true);
//             }
//         }
//         for c in col_ok {
//             if c {
//                 return (id, true);
//             }
//         }
//     }
//     return (0, false);
// }
pub fn day4_step2() {
    println!("d4s2 | not found")
}
