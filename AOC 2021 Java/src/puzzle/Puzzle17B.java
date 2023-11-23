package puzzle;

import static org.junit.Assert.assertEquals;

import org.junit.Test;

public class Puzzle17B extends Puzzle {

    @Override
    public String solve(String file_path) {
        String inputLines = processInput(file_path).get(0);

        inputLines = inputLines.substring(15);

        String[] coordinatesInStrings = inputLines.split("(, y=)|(\\.\\.)");

        int targetMinX = Integer.parseInt(coordinatesInStrings[0]);
        int targetMaxX = Integer.parseInt(coordinatesInStrings[1]);
        int targetMinY = Integer.parseInt(coordinatesInStrings[2]);
        int targetMaxY = Integer.parseInt(coordinatesInStrings[3]);

        int cnt = 0;

        for (int initialDY = targetMinY; initialDY < -targetMinY; initialDY++) {
            for (int initialDX = 0; initialDX <= targetMaxX; initialDX++) {
                int dy = initialDY;
                int dx = initialDX;
                int y = 0, x = 0;

                while (true) {
                    if (y < targetMinY)
                        break;

                    if (y >= targetMinY && y <= targetMaxY && x <= targetMaxX && x >= targetMinX) {
                        cnt++;
                        break;
                    }

                    y += dy;
                    x += dx;
                    dy += -1;
                    if (dx > 0)
                        dx += -1;
                }
            }
        }

        return Integer.toString(cnt);
    }

    @Test
    public void test() {
        String file_path = "resources/17/ex.txt";
        String result = solve(file_path);

        assertEquals("112", result);
    }

}
