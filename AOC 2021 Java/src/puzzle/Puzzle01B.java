package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

public class Puzzle01B extends Puzzle {

    public String solve(String file_path) {
        List<String> inputs = processInput(file_path);

        List<Integer> depths = new ArrayList<>();
        for (String input : inputs) {
            depths.add(Integer.parseInt(input));
        }

        int cnt = 0;
        for (int i = 0; i < depths.size() - 3; i++)
            if (depths.get(i + 3) > depths.get(i))
                cnt++;

        return Integer.toString(cnt);
    }

    @Test
    public void test() {
        String file_path = "resources/01/ex.txt";
        String result = solve(file_path);

        assertEquals("5", result);
    }
}