package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.List;

import org.junit.Test;

public class Puzzle11A extends Puzzle {

    private int[][] grid = new int[10][10];
    private int count = 0;

    private void flash(int x, int y) {
        count++;

        grid[y][x] = 0;

        // Top Left
        if (y > 0 && x > 0 && grid[y - 1][x - 1] != 0) {
            grid[y - 1][x - 1]++;

            if (grid[y - 1][x - 1] > 9) {
                flash(x - 1, y - 1);
            }
        }

        // Top
        if (y > 0 && grid[y - 1][x] != 0) {
            grid[y - 1][x]++;

            if (grid[y - 1][x] > 9) {
                flash(x, y - 1);
            }
        }

        // Top Right
        if (y > 0 && x < 9 && grid[y - 1][x + 1] != 0) {
            grid[y - 1][x + 1]++;

            if (grid[y - 1][x + 1] > 9) {
                flash(x + 1, y - 1);
            }
        }

        // Right
        if (x < 9 && grid[y][x + 1] != 0) {
            grid[y][x + 1]++;

            if (grid[y][x + 1] > 9) {
                flash(x + 1, y);
            }
        }

        // Bottom Right
        if (y < 9 && x < 9 && grid[y + 1][x + 1] != 0) {
            grid[y + 1][x + 1]++;

            if (grid[y + 1][x + 1] > 9) {
                flash(x + 1, y + 1);
            }
        }

        // Bottom
        if (y < 9 && grid[y + 1][x] != 0) {
            grid[y + 1][x]++;

            if (grid[y + 1][x] > 9) {
                flash(x, y + 1);
            }
        }

        // Bottom Left
        if (x > 0 && y < 9 && grid[y + 1][x - 1] != 0) {
            grid[y + 1][x - 1]++;

            if (grid[y + 1][x - 1] > 9) {
                flash(x - 1, y + 1);
            }
        }

        // Left
        if (x > 0 && grid[y][x - 1] != 0) {
            grid[y][x - 1]++;

            if (grid[y][x - 1] > 9) {
                flash(x - 1, y);
            }
        }

    }

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        for (int y = 0; y < inputLines.size(); y++) {
            String inputLine = inputLines.get(y);

            for (int x = 0; x < inputLine.length(); x++) {
                grid[y][x] = inputLine.charAt(x) - '0';
            }
        }

        for (int i = 0; i < 100; i++) {
            // increment by 1
            for (int y = 0; y < grid.length; y++) {
                for (int x = 0; x < grid[y].length; x++) {
                    grid[y][x] += 1;
                }
            }

            // flash
            for (int y = 0; y < grid.length; y++) {
                for (int x = 0; x < grid[y].length; x++) {
                    if (grid[y][x] > 9) {
                        flash(x, y);
                    }
                }
            }

        }

        return Integer.toString(count);
    }

    @Test
    public void test() {
        String file_path = "resources/11/ex.txt";
        String result = solve(file_path);

        assertEquals("1656", result);
    }

}
