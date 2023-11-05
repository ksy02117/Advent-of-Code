package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

public class Puzzle03B extends Puzzle {

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
    }

    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        List<Report> input1 = new ArrayList<>();
        for (String inputLine : inputLines) {
            input1.add(new Report(inputLine));
        }
        int reportSize = input1.get(0).bits.size();

        int cnt;
        for (int i = 0; i < reportSize; i++) {
            // Counting Stars
            cnt = 0;
            for (Report report : input1) {
                if (report.bits.get(i) == 1)
                    cnt++;
            }

            // Get Removing number
            int toRemove;
            if (cnt * 2 >= input1.size())
                toRemove = 0;
            else
                toRemove = 1;

            // Removing Indexes
            for (int j = 0; j < input1.size();) {
                if (input1.get(j).bits.get(i) == toRemove) {
                    input1.remove(j);
                } else {
                    j++;
                }
            }
        }

        int v1 = 0;
        for (int i = 0; i < input1.get(0).bits.size(); i++) {
            v1 *= 2;
            v1 += input1.get(0).bits.get(i);
        }

        List<Report> input2 = new ArrayList<>();
        for (String inputLine : inputLines) {
            input2.add(new Report(inputLine));
        }

        for (int i = 0; i < reportSize; i++) {
            // Counting Stars
            cnt = 0;
            for (Report report : input2) {
                if (report.bits.get(i) == 1)
                    cnt++;
            }

            // Get Removing number
            int toRemove;
            if (cnt * 2 < input2.size())
                toRemove = 0;
            else
                toRemove = 1;

            // Removing Indexes
            for (int j = 0; j < input2.size();) {
                if (input2.get(j).bits.get(i) == toRemove) {
                    input2.remove(j);
                } else {
                    j++;
                }
            }

            if (input2.size() == 1)
                break;
        }

        int v2 = 0;
        for (int i = 0; i < reportSize; i++) {
            v2 *= 2;
            v2 += input2.get(0).bits.get(i);
        }

        return Integer.toString(v1 * v2);
    }

    @Test
    public void test() {
        String file_path = "resources/03/ex.txt";
        String result = solve(file_path);

        assertEquals("230", result);
    }
}