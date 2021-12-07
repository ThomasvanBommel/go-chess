package main

func newPiece(name string, player int) Piece {
	switch name {
		case "Rook":
			return Rook{ PieceBase{ "R", player } }

		case "Knight":
			return Knight{ PieceBase{ "k", player } }

		case "Bishop":
			return Bishop{ PieceBase{ "B", player } }

		case "Queen":
			return Queen{ PieceBase{ "Q", player } }

		case "King":
			return King{ PieceBase{ "K", player } }

		default:
			return Pawn{ PieceBase{ "P", player } }
	}
}