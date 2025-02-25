const fenMap = {
  p: "/images/bP.png",
  r: "/images/bR.png",
  n: "/images/bN.png",
  b: "/images/bB.png",
  q: "/images/bQ.png",
  k: "/images/bK.png",
  P: "/images/wP.png",
  R: "/images/wR.png",
  N: "/images/wN.png",
  B: "/images/wB.png",
  Q: "/images/wQ.png",
  K: "/images/wK.png",
};

const keyValueArray = [
  [0, ""],
  [1, "p"],
  [2, "n"],
  [3, "b"],
  [4, "r"],
  [5, "q"],
  [6, "k"],
  [7, "white"],
  [8, "black"],
];

const myMap = new Map(keyValueArray);

function getPiece(type, color) {
  if (color === 7) {
    return fenMap[type.toLocaleUpperCase()] ?? "";
  }
  return fenMap[type] ?? "";
}
async function fetchPossibleMoves() {
  const possibleMoves = await fetch("/possible-moves");
  const moves = await possibleMoves.json();
  return moves;
}

function loadBoard(board) {
  const chessboard = document.getElementById("chessboard");
  for (let row of board) {
    const oddEve = board.indexOf(row) % 2;
    for (let piece of row) {
      const cellPiece = piece?.CellPiece;
      const div = document.createElement("div");
      const divImg = document.createElement("div");
      divImg.className = "image-div"
      div.className = "square";
      let imgElement = document.createElement("img");
      imgElement.className = "piece";
      div.id = piece?.CellCode;
      if (cellPiece?.PieceType !== 0) {
        div.classList.add(`${myMap.get(cellPiece?.Color)}-piece`);
      }
      if (piece) {
        imgElement.src = getPiece(
          myMap.get(cellPiece?.PieceType),
          cellPiece?.Color
        );
      }
      div.classList.add(
        row.indexOf(piece) % 2 == oddEve ? "white-cell" : "black-cell"
      );
      if (cellPiece?.PieceType !== 0) {
        divImg.appendChild(imgElement)
        div.appendChild(divImg);
      }
      chessboard.appendChild(div);
    }
  }
}

async function fetchBoard() {
  const response = await fetch("/initialise-board");
  const board = await response.json();
  loadBoard(board?.data);
  const moves = await fetchPossibleMoves();
  console.log("Helloooooooooooooooooooooooooooooooooo\nhellooo");
  addEventListners(moves?.data);
  // console.log(moves);
}

window.onload = fetchBoard;
