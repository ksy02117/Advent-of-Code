package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.junit.Test;

public class Puzzle06B extends Puzzle {

    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);
        Map<Integer, Long> fishes = new HashMap();

        String[] lines = inputLines.get(0).split(",");
        for (int i = 0; i < lines.length; i++) {
            int timer = Integer.parseInt(lines[i]);

            if (fishes.containsKey(timer)) {
                fishes.replace(timer, fishes.get(timer) + 1);
            } else {
                fishes.put(timer, 1l);
            }
        }

        for (int t = 0; t < 256; t++) {
            // age
            for (int i = 0; i < 9; i++) {
                if (fishes.containsKey(i)) {
                    fishes.put(i - 1, fishes.get(i));
                } else {
                    fishes.remove(i - 1);
                }
            }
            fishes.remove(8);

            // reset and produce
            if (fishes.containsKey(-1)) {
                if (fishes.containsKey(6)) {
                    fishes.replace(6, fishes.get(6) + fishes.get(-1));
                } else {
                    fishes.put(6, fishes.get(-1));
                }
                fishes.put(8, fishes.get(-1));
            }
            fishes.remove(-1);

        }

        long cnt = 0;
        for (int i = 0; i < 9; i++) {
            cnt += fishes.get(i);
        }

        return Long.toString(cnt);
    }

    @Test
    public void test() {
        String file_path = "resources/06/ex.txt";
        String result = solve(file_path);

        assertEquals("26984457539", result);
    }
}