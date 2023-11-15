package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

public class Puzzle15B extends Puzzle {

    private int[][] grid;
    private int[][] risk;
    private int width;
    private int height;

    private class Coordinate {
        int x, y;

        public Coordinate(int x, int y) {
            this.x = x;
            this.y = y;
        }
    }

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        height = inputLines.size();
        width = inputLines.get(0).length();
        grid = new int[height * 5][width * 5];
        risk = new int[height * 5][width * 5];

        for (int y = 0; y < height * 5; y++) {
            String inputLine = inputLines.get(y % height);

            for (int x = 0; x < width * 5; x++) {
                grid[y][x] = (inputLine.charAt(x % width) - '0' + (y / height) + (x / width) - 1) % 9 + 1;
                risk[y][x] = Integer.MAX_VALUE;
            }
        }

        risk[0][0] = 0;
        List<Coordinate> queue = new ArrayList<>();
        queue.add(new Coordinate(0, 0));
        while (!queue.isEmpty()) {
            Coordinate coordinate = queue.remove(0);
            int x = coordinate.x;
            int y = coordinate.y;

            // Left
            if (x > 0 && risk[y][x - 1] > risk[y][x] + grid[y][x - 1]) {
                risk[y][x - 1] = risk[y][x] + grid[y][x - 1];
                queue.add(new Coordinate(x - 1, y));
            }

            // Up
            if (y > 0 && risk[y - 1][x] > risk[y][x] + grid[y - 1][x]) {
                risk[y - 1][x] = risk[y][x] + grid[y - 1][x];
                queue.add(new Coordinate(x, y - 1));
            }

            // Right
            if (x < width * 5 - 1 && risk[y][x + 1] > risk[y][x] + grid[y][x + 1]) {
                risk[y][x + 1] = risk[y][x] + grid[y][x + 1];
                queue.add(new Coordinate(x + 1, y));
            }

            // Down
            if (y < height * 5 - 1 && risk[y + 1][x] > risk[y][x] + grid[y + 1][x]) {
                risk[y + 1][x] = risk[y][x] + grid[y + 1][x];
                queue.add(new Coordinate(x, y + 1));
            }
        }

        int result = risk[height * 5 - 1][width * 5 - 1];
        return Integer.toString(result);
    }

    @Test
    public void test() {
        String file_path = "resources/15/ex.txt";
        String result = solve(file_path);

        assertEquals("315", result);
    }

}
