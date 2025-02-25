let newX = 0,
  newY = 0,
  startX = 0,
  startY = 0;

function addEventListners(moves) {
  const boardCells = document.querySelectorAll(".square");
  const pieces = document.querySelectorAll(".image-div");

  pieces?.forEach((piece) => {
    piece.addEventListener("mousedown", (e) => {
      piece.style.position = "absolute";
      startX = e.clientX;
      startY = e.clientY;

      initialX = piece.style.left;
      initialY = piece.style.top;
      console.log(initialX, initialY);
      document.addEventListener("mousemove", onMouseMove);
      document.addEventListener("mouseup", onMouseUp);
    });
    function onMouseMove(e) {
      newX = startX - e.clientX;
      newY = startY - e.clientY;

      startX = e.clientX;
      startY = e.clientY;

      piece.style.top = piece.offsetTop - newY + "px";
      piece.style.left = piece.offsetLeft - newX + "px";
    }

    function onMouseUp(event) {
      const chessboard = document.getElementById("chessboard");

      const chessboardRect = chessboard.getBoundingClientRect();

      const pieceRect = piece.getBoundingClientRect();

      if (
        pieceRect.left < chessboardRect.left ||
        pieceRect.right > chessboardRect.right ||
        pieceRect.top < chessboardRect.top ||
        pieceRect.bottom > chessboardRect.bottom
      ) {
        piece.style.left = initialX;
        piece.style.top = initialY;
        document.removeEventListener("mousemove", onMouseMove);
        document.removeEventListener("mouseup", onMouseUp);
        return;
      }
      document.removeEventListener("mousemove", onMouseMove);
      document.removeEventListener("mouseup", onMouseUp);

      const x = event.clientX - chessboardRect.left;
      const y = event.clientY - chessboardRect.top;

      const row = Math.floor(y / 100);
      const col = Math.floor(x / 100);
      const targetCell = boardCells[row * 8 + col];

      targetCell.appendChild(piece);
      piece.style.position = "static";
    }
  });

  //for highlighting
  boardCells.forEach((cell) => {
    const cellId = cell.id;
    const possibleMoves = moves?.moves[cellId];
    if (possibleMoves && possibleMoves.length > 0) {
      cell.addEventListener("mouseover", () => {
        possibleMoves.forEach((m) => {
          document.querySelector(`#${m}`).style.backgroundColor = "grey";
        });
      });
      cell.addEventListener("mouseout", () => {
        possibleMoves.forEach((m) => {
          document.querySelector(`#${m}`).style.backgroundColor = "";
        });
      });
    }
  });
}
