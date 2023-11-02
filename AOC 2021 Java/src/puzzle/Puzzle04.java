package puzzle;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class Puzzle04 extends Puzzle {
    private List<BingoBoard> boards;
    private List<Integer> numbers;

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

    public void processInput(String file_path) {
        boards = new ArrayList<>();
        numbers = new ArrayList<>();

        try (BufferedReader reader = new BufferedReader(new FileReader(file_path));) {

            // Numbers
            String line = reader.readLine();
            String[] numbersInString = line.split(",");
            for (String numberInString : numbersInString) {
                numbers.add(Integer.parseInt(numberInString));
            }

            // Boards
            for (line = reader.readLine(); line != null; line = reader.readLine()) {
                BingoBoard board = new BingoBoard();
                for (int y = 0; y < 5; y++) {
                    line = reader.readLine();
                    if (line == null) {
                        return;
                    }
                    line = line.trim();
                    numbersInString = line.split("\s+");

                    for (int x = 0; x < 5; x++) {
                        board.board[y][x] = Integer.parseInt(numbersInString[x]);
                        board.mark[y][x] = false;
                    }
                }

                boards.add(board);
            }
        } catch (IOException e) {
            e.printStackTrace();
            System.exit(-1);
        }
    }

    public String solveA(String file_path) {
        processInput(file_path);

        int i, j = 0;
        for (i = 0; i < numbers.size(); i++) {
            for (j = 0; j < boards.size(); j++) {
                boolean isBingo = boards.get(j).Mark(numbers.get(i));
                if (isBingo)
                    break;
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

    public String solveB(String file_path) {
        processInput(file_path);

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

}