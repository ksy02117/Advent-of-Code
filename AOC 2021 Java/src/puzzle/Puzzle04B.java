package puzzle;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

public class Puzzle04B extends Puzzle {

    private class BingoBoard {
        int[][] board = new int[5][5];
        boolean[][] mark = new boolean[5][5];

        private boolean Mark(int num) {
            // Mark Number
            int x = 0, y = 0;
            for (y = 0; y < 5; y++) {
                for (x = 0; x < 5; x++) {
                    if (board[y][x] == num) {
                        mark[y][x] = true;
                        break;
                    }
                }

                if (x != 5) {
                    break;
                }
            }

            if (y == 5)
                return false;

            // Is Bingo
            boolean vertical = true, horizontal = true;
            for (int i = 0; i < 5; i++) {
                if (!mark[i][x]) {
                    vertical = false;
                }
                if (!mark[y][i]) {
                    horizontal = false;
                }
                if (!vertical && !horizontal) {
                    return false;
                }
            }
            return true;
        }
    }

    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);
        List<BingoBoard> boards = new ArrayList<>();
        List<Integer> numbers = new ArrayList<>();

        // Numbers
        String[] numbersInString = inputLines.get(0).split(",");
        for (String numberInString : numbersInString) {
            numbers.add(Integer.parseInt(numberInString));
        }

        // Boards
        for (int i = 2; i < inputLines.size(); i++) {
            BingoBoard board = new BingoBoard();
            int y = 0;
            for (y = 0; y < 5; y++, i++) {
                if (i >= inputLines.size()) {
                    break;
                }
                numbersInString = inputLines.get(i).trim().split("\s+");

                for (int x = 0; x < 5; x++) {
                    board.board[y][x] = Integer.parseInt(numbersInString[x]);
                    board.mark[y][x] = false;
                }
            }

            if (y == 5) {
                boards.add(board);
            }
        }

        int i, j = 0;
        for (i = 0; i < numbers.size(); i++) {
            for (j = 0; j < boards.size();) {
                boolean isBingo = boards.get(j).Mark(numbers.get(i));
                if (isBingo) {
                    if (boards.size() == 1)
                        break;
                    else
                        boards.remove(j);
                } else {
                    j++;
                }
            }

            if (j != boards.size())
                break;
        }

        int sum = 0;
        for (int y = 0; y < 5; y++) {
            for (int x = 0; x < 5; x++) {
                if (!boards.get(j).mark[y][x])
                    sum += boards.get(j).board[y][x];
            }
        }

        return Integer.toString(sum * numbers.get(i));
    }

    @Test
    public void test() {
        String file_path = "resources/04/ex.txt";
        String result = solve(file_path);

        assertEquals("1924", result);
    }
}