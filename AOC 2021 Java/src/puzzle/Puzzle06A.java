package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

public class Puzzle06A extends Puzzle {

    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);
        String[] lines = inputLines.get(0).split(",");

        List<Integer> timers = new ArrayList<>();
        for (int i = 0; i < lines.length; i++) {
            timers.add(Integer.parseInt(lines[i]));
        }

        int t = 0;
        for (t = 0; t < 80; t++) {
            for (int i = 0; i < timers.size(); i++) {
                if (timers.get(i) == 0) {
                    timers.set(i, 6);
                    timers.add(9);
                } else {
                    timers.set(i, timers.get(i) - 1);
                }
            }
        }

        return Integer.toString(timers.size());
    }

    @Test
    public void test() {
        String file_path = "resources/06/ex.txt";
        String result = solve(file_path);

        assertEquals("5934", result);
    }
}