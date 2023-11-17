package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.ListIterator;

import org.junit.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.CsvSource;

public class Puzzle16B extends Puzzle {

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

    private long parse(ListIterator<Integer> packet) {

        toDecimal(packet, 3); // version
        int id = toDecimal(packet, 3);

        if (id == 4) {
            boolean isDone = false;
            long out = 0;
            while (!isDone) {
                isDone = packet.next() == 1 ? false : true;
                out *= 16;
                out += toDecimal(packet, 4);
            }
            return out;
        } else {
            LengthType type = packet.next() == 0 ? LengthType.TotalLength : LengthType.PacketAmount;
            List<Long> subPackets = new ArrayList<>();

            switch (type) {
                case TotalLength:
                    int length = toDecimal(packet, 15);
                    List<Integer> subPacket = new ArrayList<>();
                    for (int i = 0; i < length; i++) {
                        subPacket.add(packet.next());
                    }
                    ListIterator<Integer> subPacketIterator = subPacket.listIterator();
                    while (subPacketIterator.hasNext()) {
                        subPackets.add(parse(subPacketIterator));
                    }
                    break;
                case PacketAmount:
                    int amount = toDecimal(packet, 11);
                    for (int i = 0; i < amount; i++)
                        subPackets.add(parse(packet));
                    break;
            }
            Long a = 1l;

            long out;
            switch (id) {
                case 0: // sum
                    out = 0;
                    for (Long value : subPackets) {
                        out += value;
                    }
                    return out;
                case 1: // product
                    out = 1;
                    for (Long value : subPackets) {
                        out *= value;
                    }
                    return out;
                case 2: // minimum
                    out = Collections.min(subPackets);
                    return out;
                case 3: // maximum
                    out = Collections.max(subPackets);
                    return out;
                case 5: // greater than
                    out = subPackets.get(0) > subPackets.get(1) ? 1 : 0;
                    return out;
                case 6: // less than
                    out = subPackets.get(0) < subPackets.get(1) ? 1 : 0;
                    return out;
                case 7: // equal to
                    out = subPackets.get(0).longValue() == subPackets.get(1).longValue() ? 1 : 0;
                    return out;
                default:
                    System.out.println("null");
            }
        }

        System.out.println("errroororrororo");
        return -1;
    }

    private long getResult(String inputLine) {
        List<Integer> packet = new ArrayList<>();
        for (char letter : inputLine.toCharArray()) {
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
                default:
                    System.out.println("error");
            }
        }

        ListIterator<Integer> packetIterator = packet.listIterator();

        return parse(packetIterator);
    }

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        Long result = getResult(inputLines.get(0));

        return Long.toString(result);
    }

    @Test
    public void test() {
        String file_path = "resources/16/ex.txt";
        String result = solve(file_path);

        assertEquals("2021", result);
    }

    @ParameterizedTest
    @CsvSource({
            "3, C200B40A82",
            "54,04005AC33890",
            "7, 880086C3E88112",
            "9, CE00C43D881120",
            "1, D8005AC2A8F0",
            "0, F600BC2D8F",
            "0, 9C005AC2F8F0",
            "1, 9C0141080250320F1802104A08"
    })
    public void testResult(long expected, String input) {
        long result = getResult(input);

        assertEquals(expected, result);
    }

}
