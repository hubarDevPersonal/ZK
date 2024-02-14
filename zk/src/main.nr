use dep::std;

fn validate_moves(board : [u2; 9]) {
    let mut user_moves: [u4; 3] = [0; 3];

    for i in 0..9 {
        user_moves[board[i]] += 1;
    }

    assert(user_moves[0] >= 3);
    assert(user_moves[1] >= 2);
    assert(user_moves[0] >= user_moves[1]);
    assert(user_moves[0] - user_moves[1] <= 1);
}

fn determine_winner(board : [u2; 9]) -> u2 {
    let mut winner = 2;

    // check rows
    for i in 0..3 {
        if ((board[i + (i * 2)] != 2) & (board[i + (i * 2)] == board[i + (i * 2) + 1]) & (board[i + (i * 2) + 1] == board[i + (i * 2) + 2])) {
            winner = board[i + (i * 2)];
            std::println(f"row number {i} make user {winner} win!");
        }
    }

    // check columns
    for i in 0..3 {
        if ((board[i] != 2) & (board[i] == board[i + 3]) & (board[i + 3] == board[i + 6])) {
            winner = board[i];
            std::println(f"column number {i} make user {winner} win!");
        }
    }

    // check diagonals
    if ((board[0] != 2) & (board[0] == board[4]) & (board[4] == board[8])) {
        winner = board[0];
        std::println(f"diagonal number 1 make user {winner} win!");
    } else if ((board[2] != 2) & (board[2] == board[4]) & (board[4] == board[6])) {
        winner = board[2];
        std::println(f"diagonal number 2 make user {winner} win!");
    }

    winner
}

fn main(board : [u2; 9], game : pub u16, winner : pub u2) {
    // board is an array of cells where:
    // 1st user is 0, 2nd user is 1, no move is 2
    assert(winner < 2);

    // validate if the board situation is theoretically possible
    validate_moves(board);

    // returns the winner or 2 if nobody won
    let determined = determine_winner(board) as u2;

    std::println(f"the winner is: {determined}!");

    assert(winner == determined);
}


