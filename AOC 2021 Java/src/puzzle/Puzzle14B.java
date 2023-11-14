package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.junit.Test;

public class Puzzle14B extends Puzzle {

    private Map<Character, Long> letterCount = new HashMap<>();
    private Map<String, Long> patternCount = new HashMap<>();

    private class Rule {
        String pattern;
        char insert;

        public Rule(String pattern, char insert) {
            this.pattern = pattern;
            this.insert = insert;
        }
    }

    private void incrementLetterCount(char input, long amount) {
        if (letterCount.containsKey(input))
            letterCount.replace(input, letterCount.get(input) + amount);
        else
            letterCount.put(input, amount);
    }

    private void incrementPatternCount(String input, long amount) {
        if (patternCount.containsKey(input))
            patternCount.replace(input, patternCount.get(input) + amount);
        else
            patternCount.put(input, amount);
    }

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        String inputLine = inputLines.get(0);
        for (int i = 0; i < inputLine.length(); i++) {
            incrementLetterCount(inputLine.charAt(i), 1);

            if (i == inputLine.length() - 1)
                break;

            String pattern = inputLine.substring(i, i + 2);
            incrementPatternCount(pattern, 1);
        }

        List<Rule> rules = new ArrayList<>();
        for (int lineIdx = 2; lineIdx < inputLines.size(); lineIdx++) {
            inputLine = inputLines.get(lineIdx);

            rules.add(new Rule(inputLine.substring(0, 2), inputLine.charAt(6)));
        }

        for (int step = 0; step < 40; step++) {
            Map<String, Long> previousPatternCount = patternCount;
            patternCount = new HashMap<>();

            for (var entry : previousPatternCount.entrySet()) {
                for (Rule rule : rules) {
                    if (rule.pattern.equals(entry.getKey())) {
                        incrementLetterCount(rule.insert, entry.getValue());
                        String pattern1 = rule.pattern.substring(0, 1) + rule.insert;
                        incrementPatternCount(pattern1, entry.getValue());
                        String pattern2 = rule.insert + rule.pattern.substring(1, 2);
                        incrementPatternCount(pattern2, entry.getValue());
                        break;
                    }
                }
            }
        }

        long max = 0, min = Long.MAX_VALUE;
        for (var entry : letterCount.entrySet()) {
            if (max < entry.getValue().longValue())
                max = entry.getValue().longValue();
            if (min > entry.getValue().longValue())
                min = entry.getValue().longValue();
        }

        return Long.toString(max - min);
    }

    @Test
    public void test() {
        String file_path = "resources/14/ex.txt";
        String result = solve(file_path);

        assertEquals("2188189693529", result);
    }

}
