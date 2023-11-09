

namespace AdventOfCode;

public interface Solver
{
    public string Solve(string inputFilePath);
}

public enum Part
{
    A, B
}

public enum InputType
{
    Test,
    Input
}