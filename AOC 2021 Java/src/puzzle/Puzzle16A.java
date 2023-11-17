package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.ListIterator;

import org.junit.Test;

public class Puzzle16A extends Puzzle {

    private enum LengthType {
        TotalLength,
        PacketAmount
    }

    private int toDecimal(ListIterator<Integer> packet, int length) {
        int out = 0;
        for (int i = 0; i < length; i++) {
            out *= 2;
            out += packet.next();
        }
        return out;
    }

    private int parse(ListIterator<Integer> packet) {
        int out = 0;

        int version = toDecimal(packet, 3);
        out += version;

        int id = toDecimal(packet, 3);

        if (id == 4) {
            boolean isDone = false;
            while (!isDone) {
                isDone = packet.next() == 1 ? false : true;
                toDecimal(packet, 4);
            }
        } else {
            LengthType type = packet.next() == 0 ? LengthType.TotalLength : LengthType.PacketAmount;

            switch (type) {
                case TotalLength:
                    int length = toDecimal(packet, 15);
                    int startIdx = packet.nextIndex();
                    int currentIdx;
                    do {
                        out += parse(packet);
                        currentIdx = packet.nextIndex();
                    } while (length > currentIdx - startIdx);
                    break;
                case PacketAmount:
                    int amount = toDecimal(packet, 11);
                    for (int i = 0; i < amount; i++)
                        out += parse(packet);
                    break;
            }
        }

        return out;
    }

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        List<Integer> packet = new ArrayList<>();
        for (char letter : inputLines.get(0).toCharArray()) {
            switch (letter) {
                case '0':
                    packet.addAll(Arrays.asList(0, 0, 0, 0));
                    break;
                case '1':
                    packet.addAll(Arrays.asList(0, 0, 0, 1));
                    break;
                case '2':
                    packet.addAll(Arrays.asList(0, 0, 1, 0));
                    break;
                case '3':
                    packet.addAll(Arrays.asList(0, 0, 1, 1));
                    break;
                case '4':
                    packet.addAll(Arrays.asList(0, 1, 0, 0));
                    break;
                case '5':
                    packet.addAll(Arrays.asList(0, 1, 0, 1));
                    break;
                case '6':
                    packet.addAll(Arrays.asList(0, 1, 1, 0));
                    break;
                case '7':
                    packet.addAll(Arrays.asList(0, 1, 1, 1));
                    break;
                case '8':
                    packet.addAll(Arrays.asList(1, 0, 0, 0));
                    break;
                case '9':
                    packet.addAll(Arrays.asList(1, 0, 0, 1));
                    break;
                case 'A':
                    packet.addAll(Arrays.asList(1, 0, 1, 0));
                    break;
                case 'B':
                    packet.addAll(Arrays.asList(1, 0, 1, 1));
                    break;
                case 'C':
                    packet.addAll(Arrays.asList(1, 1, 0, 0));
                    break;
                case 'D':
                    packet.addAll(Arrays.asList(1, 1, 0, 1));
                    break;
                case 'E':
                    packet.addAll(Arrays.asList(1, 1, 1, 0));
                    break;
                case 'F':
                    packet.addAll(Arrays.asList(1, 1, 1, 1));
                    break;
            }
        }

        ListIterator<Integer> packetIterator = packet.listIterator();

        int result = parse(packetIterator);

        return Integer.toString(result);
    }

    @Test
    public void test() {
        String file_path = "resources/16/ex.txt";
        String result = solve(file_path);

        assertEquals("6", result);
    }

    @Test
    public void test2() {
        String file_path = "resources/16/ex2.txt";
        String result = solve(file_path);

        assertEquals("9", result);
    }

    @Test
    public void test3() {
        String file_path = "resources/16/ex3.txt";
        String result = solve(file_path);

        assertEquals("14", result);
    }

    @Test
    public void test4() {
        String file_path = "resources/16/ex4.txt";
        String result = solve(file_path);

        assertEquals("16", result);
    }

    @Test
    public void test5() {
        String file_path = "resources/16/ex5.txt";
        String result = solve(file_path);

        assertEquals("12", result);
    }

    @Test
    public void test6() {
        String file_path = "resources/16/ex6.txt";
        String result = solve(file_path);

        assertEquals("23", result);
    }

    @Test
    public void test7() {
        String file_path = "resources/16/ex7.txt";
        String result = solve(file_path);

        assertEquals("31", result);
    }

}
