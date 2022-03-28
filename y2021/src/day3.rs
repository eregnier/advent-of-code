use std::fs::File;
use std::io::{BufRead, BufReader};

pub fn day3_step1() {
    let reader = BufReader::new(File::open("inputs/d3").expect("cannot open 'input' file"));
    let mut ones = vec![0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
    let mut zeroes = vec![0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
    for line in reader.lines() {
        let chars = line.unwrap().chars().collect::<Vec<char>>();
        for i in 0..12 {
            match chars[i] {
                '0' => zeroes[i] += 1,
                '1' => ones[i] += 1,
                _ => print!("Err{}#", chars[i]),
            }
        }
    }
    let mut gamma = Vec::new();
    let mut epsilon = Vec::new();
    for i in 0..12 {
        if ones[i] > zeroes[i] {
            gamma.push('1');
            epsilon.push('0');
        } else {
            gamma.push('0');
            epsilon.push('1');
        }
    }
    let e = &epsilon.iter().cloned().collect::<String>();
    let eval = isize::from_str_radix(e, 2).unwrap();
    let g = &gamma.iter().cloned().collect::<String>();
    let gval = isize::from_str_radix(g, 2).unwrap();
    println!("d3s1 | {} X {} => {}", eval, gval, eval * gval);
}

pub fn day3_step2() {
    println!("d3s2 | not found")
    // let reader = BufReader::new(File::open("inputs/d3").expect("cannot open 'input' file"));
    // let word_size = 12;
    // let mut ones = vec![0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
    // let mut zeroes = vec![0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
    // let mut lcv = vec![];
    // let mut mcv = vec![];
    // let mut oxygen = vec![];
    // let mut scruber = vec![];
    // // let mut lc = 0;
    // for line in reader.lines() {
    //     let line_str = line.unwrap();
    //     println!("{}", line_str);
    //     // lc += 1;
    //     let chars = line_str.chars().collect::<Vec<char>>();
    //     oxygen.push(chars.to_vec());
    //     scruber.push(chars.to_vec());
    //     for i in 0..word_size {
    //         match chars[i] {
    //             '0' => zeroes[i] += 1,
    //             '1' => ones[i] += 1,
    //             _ => print!("Err{}#", chars[i]),
    //         }
    //     }
    //     // if lc == 8 {
    //     //     break;
    //     // }
    // }
    // for i in 0..word_size {
    //     println!("\n |{}|--------------", i);
    //     println!(
    //         "oxygen len {} | scruber len {}",
    //         oxygen.len(),
    //         scruber.len()
    //     );
    //     let mut filter_oxygen = vec![];
    //     let mut filter_scruber = vec![];
    //     let mut less_common_value = '0';
    //     let mut most_common_value = '1';
    //     if ones[i] < zeroes[i] {
    //         less_common_value = '1';
    //     }
    //     if zeroes[i] < ones[i] {
    //         less_common_value = '0';
    //     }
    //     lcv.push(less_common_value);
    //     if ones[i] > zeroes[i] {
    //         most_common_value = '1';
    //     }
    //     if zeroes[i] > ones[i] {
    //         most_common_value = '0';
    //     }
    //     mcv.push(most_common_value);

    //     println!(
    //         "ones {} | zeroes {} | 0>1 {}",
    //         ones[i],
    //         zeroes[i],
    //         zeroes[i] > ones[i],
    //     );
    //     for j in 0..oxygen.len() {
    //         println!("oxy[{}] > {}", j, oxygen[j][i]);
    //         if oxygen[j][i] == most_common_value {
    //             filter_oxygen.push(oxygen[j].to_vec());
    //             print!(" ✔️ ");
    //         } else {
    //             print!(" x ");
    //         }
    //         println!("{}", oxygen[j].iter().collect::<String>())
    //     }
    //     for j in 0..scruber.len() {
    //         println!("scr[{}] > {}", j, scruber[j][i]);
    //         if scruber[j][i] == less_common_value {
    //             filter_scruber.push(scruber[j].to_vec());
    //             print!(" ✔️ ");
    //         } else {
    //             print!(" x ");
    //         }
    //         println!("{}", scruber[j].iter().collect::<String>())
    //     }
    //     if filter_oxygen.len() >= 1 {
    //         oxygen = filter_oxygen;
    //     }
    //     if filter_scruber.len() >= 1 {
    //         scruber = filter_scruber;
    //     }
    // }
    // println!("### len > ox {} | sc {}", oxygen.len(), scruber.len());
    // // println!(" + last bit for all entries left (11)");
    // // for i in 0..oxygen.len() {
    // //     println!("o {} => {}", i, oxygen[i][11]);
    // // }
    // let ox = &oxygen[0].iter().cloned().collect::<String>();
    // let oxval = isize::from_str_radix(ox, 2).unwrap();
    // let scr = &scruber[0].iter().cloned().collect::<String>();
    // let scrval = isize::from_str_radix(scr, 2).unwrap();
    // println!("d3s2 | ox {} | scr {} =  {}", oxval, scrval, oxval * scrval);
    // println!(
    //     "ox {} | {}",
    //     oxygen.len(),
    //     oxygen[0].iter().collect::<String>()
    // );
    // println!(
    //     "sc {} | {}",
    //     scruber.len(),
    //     scruber[0].iter().collect::<String>()
    // );
    // println!("mcv(oxy) {}", mcv.iter().collect::<String>());
    // println!("lcv(scr) {}", lcv.iter().collect::<String>());
    // println!("1\t0 {} {}", ones.len(), zeroes.len());

    // for i in 0..word_size {
    //     println!("{}\t{}", ones[i], zeroes[i]);
    // }
}
