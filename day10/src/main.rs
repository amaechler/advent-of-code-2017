fn main() {
    let input: Vec<usize> = (0..256).collect();
    let hash_lengths: Vec<usize> = vec![
        83, 0, 193, 1, 254, 237, 187, 40, 88, 27, 2, 255, 149, 29, 42, 100
    ];

    let mut first_input = input.clone();
    hash_knots(first_input.as_mut_slice(), hash_lengths.as_slice(), 1);
    println!("Solution 1: {:?}", first_input[0] * first_input[1]);

    // build input for second solution
    let mut ascii_hash_lengths = "83,0,193,1,254,237,187,40,88,27,2,255,149,29,42,100"
        .chars()
        .map(|x| x as usize)
        .collect::<Vec<_>>();

    let mut prefix_vec: Vec<usize> = vec![17, 31, 73, 47, 23];
    ascii_hash_lengths.append(&mut prefix_vec);

    let mut second_input = input.clone();
    hash_knots(
        second_input.as_mut_slice(),
        ascii_hash_lengths.as_slice(),
        64);

    let dense_hash = dense_hash(&second_input, 16);

    println!("Solution 2: {:?}", dense_hash);
}

// hashes the input array in-memory
pub fn hash_knots(knot_vec: &mut [usize], hash_lengths: &[usize], number_of_rounds: usize) {
    let knot_len = knot_vec.len();

    let mut current_position = 0;
    let mut skip_size = 0;
    
    for _r in 0..number_of_rounds {
        for hl in hash_lengths {
            // reverse the relevant parts of the knot
            for x in 0..*hl / 2 {
                knot_vec.swap(
                    (current_position + x) % knot_len,
                    (current_position + (*hl - 1 - x)) % knot_len,
                )
            }

            // update current position and skip size
            current_position = (current_position + *hl + skip_size) % knot_len;
            skip_size += 1;
        }
    }
}

pub fn dense_hash(input: &[usize], chunk_size: usize) -> String {
    let mut dense_hash: String = String::new();
    for slice in input.chunks(chunk_size) {
        let slice_scanner = slice
            .iter()
            .scan(0, |state, &x| {
                if x == slice[0] { 
                    *state = x;
                    return Some(*state);
                } else {
                    *state = *state ^ x;
                    return Some(*state);
                }
            })
            .last();

        dense_hash = format!("{}{:x}", dense_hash, slice_scanner.unwrap());
    }

    return dense_hash;
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn example_from_day10() {
        let mut input: Vec<usize> = (0..5).collect();
        let lenghts = vec![3, 4, 1, 5];

        hash_knots(&mut input, &lenghts, 1);

        assert_eq!(vec![3, 4, 2, 1, 0], input);
    }
}
