package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.junit.Test;

public class Puzzle14A extends Puzzle {

    private class Rule {
        char letter1, letter2;
        char insert;

        public Rule(char letter1, char letter2, char insert) {
            this.letter1 = letter1;
            this.letter2 = letter2;
            this.insert = insert;
        }
    }

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        List<Character> polymer = new ArrayList<>();
        for (char letter : inputLines.get(0).toCharArray())
            polymer.add(letter);

        List<Rule> rules = new ArrayList<>();
        for (int lineIdx = 2; lineIdx < inputLines.size(); lineIdx++) {
            String inputLine = inputLines.get(lineIdx);

            rules.add(new Rule(inputLine.charAt(0), inputLine.charAt(1), inputLine.charAt(6)));
        }

        for (int step = 0; step < 10; step++) {
            List<Character> newPolymer = new ArrayList<>();
            for (int i = 0; i < polymer.size() - 1; i++) {
                newPolymer.add(polymer.get(i));

                for (Rule rule : rules) {
                    if (polymer.get(i).equals(rule.letter1) && polymer.get(i + 1).equals(rule.letter2)) {
                        newPolymer.add(rule.insert);
                        break;
                    }
                }
            }
            newPolymer.add(polymer.get(polymer.size() - 1));
            polymer = newPolymer;
        }

        Map<Character, Long> letterCount = new HashMap<>();
        for (Character character : polymer) {
            if (letterCount.containsKey(character))
                letterCount.replace(character, letterCount.get(character) + 1);
            else
                letterCount.put(character, 1l);
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

        assertEquals("1588", result);
    }

}
