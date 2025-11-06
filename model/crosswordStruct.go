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
    CrosswordID int
}

type TokenSample struct{
    Token string `json:"token" example:"your_token"`
}

type CrosswordSolution struct{
    Sol []UnitClue `json:"sol"`
    Id uint `json:"crosswordid"`
}