using AdventOfCode;

if (args.Length == 0)
{
    Runner.RunLatest();
    return;
}

// create new solver for next puzzle
if (args[0].Equals("create"))
{
    FileGenerator.GenerateNextSolver();
    return;
}

// test
if (args[0].Equals("test"))
{
    Runner.RunLatest(InputType.Test);
    return;
}