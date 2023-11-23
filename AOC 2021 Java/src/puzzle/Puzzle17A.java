package puzzle;

import static org.junit.Assert.assertEquals;

import org.junit.Test;

public class Puzzle17A extends Puzzle {

    @Override
    public String solve(String file_path) {
        String inputLines = processInput(file_path).get(0);

        inputLines = inputLines.substring(15);

        String[] coordinatesInStrings = inputLines.split("(, y=)|(\\.\\.)");

        int targetMinY = Integer.parseInt(coordinatesInStrings[2]);

        int result = 0;

        for (int i = 0; i < targetMinY * -1; i++) {
            result += i;
        }

        return Integer.toString(result);
    }

    @Test
    public void test() {
        String file_path = "resources/17/ex.txt";
        String result = solve(file_path);

        assertEquals("45", result);
    }

}
