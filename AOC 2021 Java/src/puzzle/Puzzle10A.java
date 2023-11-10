package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.List;
import java.util.Stack;

import org.junit.Test;

public class Puzzle10A extends Puzzle {

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        int point = 0;
        for (String inputLine : inputLines) {
            Stack<Character> stack = new Stack<>();

            letterLoop: for (char inputLetter : inputLine.toCharArray()) {
                switch (inputLetter) {
                    case '(':
                    case '[':
                    case '{':
                    case '<':
                        stack.push(inputLetter);
                        break;
                    case ')':
                        if (stack.pop() != '(') {
                            point += 3;
                            break letterLoop;
                        }
                        break;
                    case ']':
                        if (stack.pop() != '[') {
                            point += 57;
                            break letterLoop;
                        }
                        break;
                    case '}':
                        if (stack.pop() != '{') {
                            point += 1197;
                            break letterLoop;
                        }
                        break;
                    case '>':
                        if (stack.pop() != '<') {
                            point += 25137;
                            break letterLoop;
                        }
                        break;
                }
            }

        }

        return Integer.toString(point);
    }

    @Test
    public void test() {
        String file_path = "resources/10/ex.txt";
        String result = solve(file_path);

        assertEquals("26397", result);
    }

}
