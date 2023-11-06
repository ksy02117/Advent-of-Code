package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.Arrays;
import java.util.Comparator;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.Map;
import java.util.HashMap;

import org.junit.Test;

public class Puzzle08B extends Puzzle {

    private enum Segment {
        Top,
        TopLeft,
        TopRight,
        Middle,
        BottomLeft,
        BottomRight,
        Bottom
    }

    private class Digit {
        public static final Set<Segment> ZERO = new HashSet<>(Arrays.asList(Segment.Top, Segment.TopLeft,
                Segment.TopRight, Segment.BottomLeft, Segment.BottomRight, Segment.Bottom));
        public static final Set<Segment> ONE = new HashSet<>(Arrays.asList(Segment.TopRight, Segment.BottomRight));
        public static final Set<Segment> TWO = new HashSet<>(
                Arrays.asList(Segment.Top, Segment.TopRight, Segment.Middle, Segment.BottomLeft, Segment.Bottom));
        public static final Set<Segment> THREE = new HashSet<>(
                Arrays.asList(Segment.Top, Segment.TopRight, Segment.Middle, Segment.BottomRight, Segment.Bottom));
        public static final Set<Segment> FOUR = new HashSet<>(
                Arrays.asList(Segment.TopLeft, Segment.TopRight, Segment.Middle, Segment.BottomRight));
        public static final Set<Segment> FIVE = new HashSet<>(
                Arrays.asList(Segment.Top, Segment.TopLeft, Segment.Middle, Segment.BottomRight, Segment.Bottom));
        public static final Set<Segment> SIX = new HashSet<>(Arrays.asList(Segment.Top, Segment.TopLeft, Segment.Middle,
                Segment.BottomLeft, Segment.BottomRight, Segment.Bottom));
        public static final Set<Segment> SEVEN = new HashSet<>(
                Arrays.asList(Segment.Top, Segment.TopRight, Segment.BottomRight));
        public static final Set<Segment> EIGHT = new HashSet<>(Arrays.asList(Segment.Top, Segment.TopLeft,
                Segment.TopRight, Segment.Middle, Segment.BottomLeft, Segment.BottomRight, Segment.Bottom));
        public static final Set<Segment> NINE = new HashSet<>(Arrays.asList(Segment.Top, Segment.TopLeft,
                Segment.TopRight, Segment.Middle, Segment.BottomRight, Segment.Bottom));
    }
    // cnt 8687497
    // 1 : --c--f- 2
    // 7 : a-c--f- 3
    // 4 : -bcd-f- 4
    // 2 : a-cde-g 5
    // 3 : a-cd-fg 5
    // 5 : ab-d-fg 5
    // 0 : abc-efg 6
    // 6 : ab-defg 6
    // 9 : abcd-fg 6
    // 8 : abcdefg 7

    // cnt 8877
    // 1 : -c-- 2
    // 7 : ac-- 3
    // 4 : -cd- 4
    // 2 : acdg 5
    // 3 : acdg 5
    // 5 : a-dg 5
    // 0 : ac-g 6
    // 6 : a-dg 6
    // 9 : acdg 6
    // 8 : acdg 7

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        int sum = 0;
        for (String inputLine : inputLines) {
            String[] inputDivided = inputLine.split(" \\| ");

            // Process All Digits
            String[] allDigits = inputDivided[0].split(" ");

            Arrays.sort(allDigits, new Comparator<String>() {
                public int compare(String a, String b) {
                    return a.length() - b.length();
                }
            });

            // Get TopLeft & BottomLeft & BottomRight based on their count
            Map<Character, Integer> letterCount = new HashMap<>();
            letterCount.put('a', 0);
            letterCount.put('b', 0);
            letterCount.put('c', 0);
            letterCount.put('d', 0);
            letterCount.put('e', 0);
            letterCount.put('f', 0);
            letterCount.put('g', 0);

            for (String oneDigit : allDigits) {
                for (char letter : oneDigit.toCharArray()) {
                    letterCount.replace(letter, letterCount.get(letter) + 1);
                }
            }

            Map<Character, Segment> charToSegment = new HashMap<>();
            for (Map.Entry<Character, Integer> entry : letterCount.entrySet()) {
                switch (entry.getValue()) {
                    case 4:
                        charToSegment.put(entry.getKey(), Segment.BottomLeft);
                        break;
                    case 6:
                        charToSegment.put(entry.getKey(), Segment.TopLeft);
                        break;
                    case 7:
                        // Get Middle and Bottom (4 has only Middle)
                        if (allDigits[2].contains(entry.getKey().toString()))
                            charToSegment.put(entry.getKey(), Segment.Middle);
                        else
                            charToSegment.put(entry.getKey(), Segment.Bottom);
                        break;
                    case 8:
                        // Get Top & TopRight (1 has only TopRight)
                        if (allDigits[0].contains(entry.getKey().toString()))
                            charToSegment.put(entry.getKey(), Segment.TopRight);
                        else
                            charToSegment.put(entry.getKey(), Segment.Top);
                        break;
                    case 9:
                        charToSegment.put(entry.getKey(), Segment.BottomRight);
                        break;
                }
            }

            // Get Answer
            String[] answerDigits = inputDivided[1].split(" ");
            int number = 0;
            for (String AnswerDigit : answerDigits) {
                Set<Segment> digitSegments = new HashSet<>();

                for (char letter : AnswerDigit.toCharArray()) {
                    digitSegments.add(charToSegment.get(letter));
                }

                number *= 10;
                if (digitSegments.equals(Digit.ONE))
                    number += 1;
                if (digitSegments.equals(Digit.TWO))
                    number += 2;
                if (digitSegments.equals(Digit.THREE))
                    number += 3;
                if (digitSegments.equals(Digit.FOUR))
                    number += 4;
                if (digitSegments.equals(Digit.FIVE))
                    number += 5;
                if (digitSegments.equals(Digit.SIX))
                    number += 6;
                if (digitSegments.equals(Digit.SEVEN))
                    number += 7;
                if (digitSegments.equals(Digit.EIGHT))
                    number += 8;
                if (digitSegments.equals(Digit.NINE))
                    number += 9;
                if (digitSegments.equals(Digit.ZERO))
                    number += 0;
            }

            sum += number;
        }

        return Integer.toString(sum);
    }

    @Test
    public void test() {
        String file_path = "resources/08/ex.txt";
        String result = solve(file_path);

        assertEquals("61229", result);
    }
}
