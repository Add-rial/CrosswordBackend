package model

type Cell struct{
    IsBlank bool
    NumberAssociated int
}

type Clues struct{
    Across []UnitClue
    Down []UnitClue
}

type Crossword struct{
    Clues Clues
    Rows int
    Columns int
    Grid [][]Cell
}
