package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

public class Puzzle03A extends Puzzle {

    private class Report {
        List<Integer> bits;

        public Report(String input) {
            bits = new ArrayList<>();

            for (int i = 0; i < input.length(); i++) {
                switch (input.charAt(i)) {
                    case '0':
                    case '1':
                        bits.add(input.charAt(i) - '0');
                        break;
                }
            }
        }

        public Report(int initialSize) {
            bits = new ArrayList<>(initialSize);

            for (int i = 0; i < initialSize; i++)
                bits.add(0);
        }

        private void Add(Report report) {
            for (int i = 0; i < bits.size(); i++) {
                // bits[i] += report.bits[i];
                bits.set(i, bits.get(i) + report.bits.get(i));
            }
        }
    }

    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        List<Report> input = new ArrayList<>();
        for (String inputLine : inputLines) {
            input.add(new Report(inputLine));
        }

        Report counts = new Report(input.get(0).bits.size());

        for (Report report : input) {
            counts.Add(report);
        }

        int v1 = 0, v2 = 0;
        for (Integer count : counts.bits) {
            v1 *= 2;
            v2 *= 2;
            if (count > input.size() / 2) {
                v1++;
            } else {
                v2++;
            }
        }

        return Integer.toString(v1 * v2);
    }

    @Test
    public void test() {
        String file_path = "resources/03/ex.txt";
        String result = solve(file_path);

        assertEquals("198", result);
    }
}