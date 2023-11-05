package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

public class Puzzle05B extends Puzzle {

    private class Board {
        List<List<Integer>> grid;

        public Board() {
            grid = new ArrayList<>();
            grid.add(new ArrayList<>());
        }

        private void resizeBoard(int maxX, int maxY) {
            // resize based on maxY
            while (grid.size() <= maxY) {
                List<Integer> row = new ArrayList<>();
                for (int i = 0; i < grid.get(0).size(); i++) {
                    row.add(0);
                }
                grid.add(row);
            }

            // resize based on maxX
            for (List<Integer> row : grid) {
                while (row.size() <= maxX) {
                    row.add(0);
                }
            }
        }

        private void Add(Line line) {
            int maxY = line.sy > line.ey ? line.sy : line.ey;
            int maxX = line.sx > line.ex ? line.sx : line.ex;
            resizeBoard(maxX, maxY);

            // Horizontal
            if (line.sx == line.ex) {
                int from = (line.sy < line.ey) ? line.sy : line.ey;
                int to = (line.sy > line.ey) ? line.sy : line.ey;
                for (int y = from; y <= to; y++) {
                    grid.get(y).set(line.sx, grid.get(y).get(line.sx) + 1);
                }
                return;
            }

            // Vertical
            if (line.sy == line.ey) {
                int from = (line.sx < line.ex) ? line.sx : line.ex;
                int to = (line.sx > line.ex) ? line.sx : line.ex;
                for (int x = from; x <= to; x++) {
                    grid.get(line.sy).set(x, grid.get(line.sy).get(x) + 1);
                }
                return;
            }

            // Diagonal
            if (line.sx - line.ex == line.sy - line.ey || line.sx - line.ex == line.ey - line.sy) {
                int dx = line.sx > line.ex ? -1 : 1;
                int dy = line.sy > line.ey ? -1 : 1;
                int length = line.sx > line.ex ? line.sx - line.ex : line.ex - line.sx;

                for (int i = 0; i <= length; i++) {
                    int x = line.sx + dx * i;
                    int y = line.sy + dy * i;
                    grid.get(y).set(x, grid.get(y).get(x) + 1);
                }

            }
        }
    }

    private class Line {
        int sx, sy, ex, ey;

        public Line(String input) {
            String[] values = input.split("(,| -> )");

            sx = Integer.parseInt(values[0]);
            sy = Integer.parseInt(values[1]);
            ex = Integer.parseInt(values[2]);
            ey = Integer.parseInt(values[3]);
        }
    }

    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        List<Line> lines = new ArrayList<>();
        for (String inputLine : inputLines) {
            lines.add(new Line(inputLine));
        }

        Board board = new Board();
        for (Line line : lines) {
            board.Add(line);
        }

        int cnt = 0;
        for (int y = 0; y < board.grid.size(); y++) {
            for (int x = 0; x < board.grid.get(y).size(); x++) {
                if (board.grid.get(y).get(x) >= 2) {
                    cnt++;
                }
            }
        }

        return Integer.toString(cnt);
    }

    @Test
    public void test() {
        String file_path = "resources/05/ex.txt";
        String result = solve(file_path);

        assertEquals("12", result);
    }
}