package puzzle;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class Puzzle03 extends Puzzle {

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
                bits.set(i, bits.get(i) + report.bits.get(i));
            }
        }
    }

    public List<Report> processInput(String file_path) {
        List<Report> out = new ArrayList<>();

        try (BufferedReader reader = new BufferedReader(new FileReader(file_path));) {
            for (String line = reader.readLine(); line != null; line = reader.readLine()) {
                out.add(new Report(line));
            }
        } catch (IOException e) {
            e.printStackTrace();
            System.exit(-1);
        }

        return out;
    }

    public String solveA(String file_path) {
        List<Report> input = processInput(file_path);

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

    public String solveB(String file_path) {
        List<Report> input1 = processInput(file_path);
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

        List<Report> input2 = processInput(file_path);

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

}