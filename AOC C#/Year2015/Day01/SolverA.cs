

namespace AdventOfCode.Y2015.D01;

class SolverA : Solver
{
    public string Solve(string inputFilePath)
    {
        IEnumerable<string> inputLines = InputFileReader.readInputFile(inputFilePath);

        string output = "";

        foreach (string inputLine in inputLines)
        {
            output += inputLine + "\n";
        }

        return output;
    }
}