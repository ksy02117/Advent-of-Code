package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

public class Puzzle07B extends Puzzle {

    @Override
    public String solve(String file_path) {
        String[] inputs = processInput(file_path).get(0).split(",");

        List<Integer> positions = new ArrayList<>();
        for (String input : inputs) {
            positions.add(Integer.parseInt(input));
        }

        double total = 0;
        for (int position : positions) {
            total += position;
        }

        int target = (int) (total / positions.size());

        int sum1 = 0, sum2 = 0;
        for (int position : positions) {
            int dif1 = target - position;
            int dif2 = target + 1 - position;

            if (dif1 < 0)
                dif1 *= -1;
            if (dif2 < 0)
                dif2 *= -1;

            sum1 += dif1 * (dif1 + 1) / 2;
            sum2 += dif2 * (dif2 + 1) / 2;
        }

        int answer = sum1 < sum2 ? sum1 : sum2;

        return Integer.toString(answer);
    }

    @Override
    @Test
    public void test() {
        String file_path = "resources/07/ex.txt";
        String result = solve(file_path);

        assertEquals("168", result);
    }

}
