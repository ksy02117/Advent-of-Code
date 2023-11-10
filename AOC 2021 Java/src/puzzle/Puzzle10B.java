package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.Collection;
import java.util.Collections;
import java.util.List;
import java.util.Stack;

import org.junit.Test;

public class Puzzle10B extends Puzzle {

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        List<Long> points = new ArrayList<>();
        lineLoop: for (String inputLine : inputLines) {
            Stack<Character> stack = new Stack<>();

            for (char inputLetter : inputLine.toCharArray()) {
                switch (inputLetter) {
                    case '(':
                    case '[':
                    case '{':
                    case '<':
                        stack.push(inputLetter);
                        break;
                    case ')':
                        if (stack.pop() != '(') {
                            continue lineLoop;
                        }
                        break;
                    case ']':
                        if (stack.pop() != '[') {
                            continue lineLoop;
                        }
                        break;
                    case '}':
                        if (stack.pop() != '{') {
                            continue lineLoop;
                        }
                        break;
                    case '>':
                        if (stack.pop() != '<') {
                            continue lineLoop;
                        }
                        break;
                }
            }

            long point = 0;
            while (!stack.empty()) {
                point *= 5;
                switch (stack.pop()) {
                    case '(':
                        point += 1;
                        break;
                    case '[':
                        point += 2;
                        break;
                    case '{':
                        point += 3;
                        break;
                    case '<':
                        point += 4;
                        break;
                }
            }

            points.add(point);
        }

        Collections.sort(points);

        long answer = points.get(points.size() / 2);

        return Long.toString(answer);
    }

    @Test
    public void test() {
        String file_path = "resources/10/ex.txt";
        String result = solve(file_path);

        assertEquals("288957", result);
    }

}
