package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

import org.junit.Test;

public class Puzzle07A extends Puzzle {

    @Override
    public String solve(String file_path) {
        String[] inputs = processInput(file_path).get(0).split(",");

        List<Integer> positions = new ArrayList<>();
        for (String input : inputs) {
            positions.add(Integer.parseInt(input));
        }

        Collections.sort(positions);

        int target = positions.get(positions.size() / 2);

        int sum = 0;
        for (int position : positions) {
            int dif = target - position;

            if (dif < 0)
                dif *= -1;

            sum += dif;
        }

        return Integer.toString(sum);
    }

    @Override
    @Test
    public void test() {
        String file_path = "resources/07/ex.txt";
        String result = solve(file_path);

        assertEquals("37", result);
    }

}
