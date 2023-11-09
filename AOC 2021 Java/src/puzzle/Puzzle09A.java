package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.List;

import org.junit.Test;

public class Puzzle09A extends Puzzle {

    private int[][] grid;

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);
        int x, y;

        grid = new int[inputLines.size()][inputLines.get(0).length()];
        y = 0;
        for (String inputLine : inputLines) {
            x = 0;
            for (char inputLetter : inputLine.toCharArray()) {
                grid[y][x] = inputLetter - '0';
                x++;
            }
            y++;
        }

        int sum = 0;
        for (y = 0; y < grid.length; y++) {
            for (x = 0; x < grid[y].length; x++) {
                // Top
                if (y != 0 && grid[y - 1][x] <= grid[y][x])
                    continue;

                // Bottom
                if (y != grid.length - 1 && grid[y + 1][x] <= grid[y][x])
                    continue;

                // Left
                if (x != 0 && grid[y][x - 1] <= grid[y][x])
                    continue;

                // Right
                if (x != grid[y].length - 1 && grid[y][x + 1] <= grid[y][x])
                    continue;

                System.out.print(grid[y][x]);
                if (y != 0)
                    System.out.print(grid[y - 1][x]);
                if (y != grid.length - 1)
                    System.out.print(grid[y + 1][x]);
                if (x != 0)
                    System.out.print(grid[y][x - 1]);
                if (x != grid[y].length - 1)
                    System.out.print(grid[y][x + 1]);

                System.out.println();
                sum += grid[y][x] + 1;
            }
        }

        return Integer.toString(sum);
    }

    @Test
    public void test() {
        String file_path = "resources/09/ex.txt";
        String result = solve(file_path);

        assertEquals("15", result);
    }
}
