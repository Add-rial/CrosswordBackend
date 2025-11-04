package model

type Cell struct{
    IsBlank bool
    NumberAssociated int
}

type UnitClue struct{
    ClueID int 
    ClueText string
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
