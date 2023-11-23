package puzzle;

import static org.junit.Assert.assertEquals;

import java.text.StringCharacterIterator;
import java.util.List;

import org.junit.Test;

public class Puzzle18B extends Puzzle {

    private abstract class Element {
        public Pair parent;

        public abstract int magnitude();
    }

    private class Pair extends Element {
        public Element left;
        public Element right;

        @Override
        public int magnitude() {
            return left.magnitude() * 3 + right.magnitude() * 2;
        }

        @Override
        public String toString() {
            return "[" + left.toString() + "," + right.toString() + "]";
        }

        public void explodeLeft() {
            if (left instanceof Pair)
                System.exit(-1);

            int value = ((Value) left).value;
            Element currentElement = this;

            // find deepest common parent of exploding node and adjacent node
            while (true) {
                if (currentElement.parent == null)
                    return;

                if (currentElement.parent.left != currentElement)
                    break;

                currentElement = currentElement.parent;
            }

            currentElement = ((Pair) currentElement).parent.left;

            while (currentElement instanceof Pair) {
                currentElement = ((Pair) currentElement).right;
            }

            ((Value) currentElement).value += value;
            return;
        }

        public void explodeRight() {
            if (right instanceof Pair)
                System.exit(-1);

            int value = ((Value) right).value;
            Element currentElement = this;

            // find deepest common parent of exploding node and adjacent node
            while (true) {
                if (currentElement.parent == null)
                    return;

                if (currentElement.parent.right != currentElement)
                    break;

                currentElement = currentElement.parent;
            }

            currentElement = ((Pair) currentElement).parent.right;

            while (currentElement instanceof Pair) {
                currentElement = ((Pair) currentElement).left;
            }

            ((Value) currentElement).value += value;
            return;
        }

        public boolean explodeMostLeft(int depth) {
            if (depth >= 4) {
                explodeLeft();
                explodeRight();
                if (this == parent.left) {
                    parent.left = new Value(0);
                    parent.left.parent = parent;
                } else {
                    parent.right = new Value(0);
                    parent.right.parent = parent;
                }

                return true;
            }

            if (left instanceof Pair) {
                Pair leftPair = (Pair) left;
                if (leftPair.explodeMostLeft(depth + 1))
                    return true;
            }

            if (right instanceof Pair) {
                Pair rightPair = (Pair) right;
                if (rightPair.explodeMostLeft(depth + 1))
                    return true;
            }

            return false;
        }

        public boolean splitMostLeft() {
            if (left instanceof Pair) {
                if (((Pair) left).splitMostLeft())
                    return true;
            } else {
                int value = ((Value) left).value;
                if (value >= 10) {
                    int leftValue = value / 2;
                    int rightValue = value / 2 + (value % 2);

                    Pair newPair = new Pair();
                    newPair.parent = this;
                    newPair.left = new Value(leftValue);
                    newPair.left.parent = newPair;
                    newPair.right = new Value(rightValue);
                    newPair.right.parent = newPair;

                    left = newPair;

                    return true;
                }
            }

            if (right instanceof Pair) {
                if (((Pair) right).splitMostLeft())
                    return true;
            } else {
                int value = ((Value) right).value;
                if (value >= 10) {
                    int leftValue = value / 2;
                    int rightValue = value / 2 + (value % 2);

                    Pair newPair = new Pair();
                    newPair.parent = this;
                    newPair.left = new Value(leftValue);
                    newPair.left.parent = newPair;
                    newPair.right = new Value(rightValue);
                    newPair.right.parent = newPair;

                    right = newPair;

                    return true;
                }
            }

            return false;
        }
    }

    private class Value extends Element {
        private int value;

        public Value(int value) {
            this.value = value;
        }

        @Override
        public int magnitude() {
            return value;
        }

        @Override
        public String toString() {
            return Integer.toString(value);
        }
    }

    private Element parseElement(StringCharacterIterator input) {

        Pair out = new Pair();

        // left
        char letter = input.next();
        if (letter == '[') {
            out.left = parseElement(input);
            input.next(); // comma
        } else {
            int value = 0;
            for (; letter >= '0' && letter <= '9'; letter = input.next()) {
                value *= 10;
                value += letter - '0';
            }
            out.left = new Value(value);
        }
        out.left.parent = out;

        // right
        letter = input.next();
        if (letter == '[') {
            out.right = parseElement(input);
            input.next(); // comma
        } else {
            int value = 0;
            for (; letter >= '0' && letter <= '9'; letter = input.next()) {
                value *= 10;
                value += letter - '0';
            }
            out.right = new Value(value);
        }
        out.right.parent = out;

        return out;
    }

    private Element addElements(Element left, Element right) {
        Pair result = new Pair();
        result.left = left;
        result.left.parent = result;
        result.right = right;
        result.right.parent = result;

        // loop until reduced
        while (true) {
            // explosion
            if (result.explodeMostLeft(0))
                continue;

            // split
            if (result.splitMostLeft())
                continue;

            break;
        }

        return result;
    }

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        int max = 0;

        for (int i = 0; i < inputLines.size(); i++) {
            for (int j = 0; j < inputLines.size(); j++) {
                if (i == j)
                    continue;

                Element element1 = parseElement(new StringCharacterIterator(inputLines.get(i)));
                Element element2 = parseElement(new StringCharacterIterator(inputLines.get(j)));
                Element result = addElements(element1, element2);
                int magnitude = result.magnitude();

                if (max < magnitude)
                    max = magnitude;
            }
        }

        return Integer.toString(max);
    }

    @Test
    public void test() {
        String file_path = "resources/18/ex.txt";
        String result = solve(file_path);

        assertEquals("3993", result);
    }
}
