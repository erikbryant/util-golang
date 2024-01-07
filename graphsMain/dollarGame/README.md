# The Dollar Game

The dollar game is a simple game of solitaire. Start with a connected graph where each vertex has some integer value (positive, negative, or zero). Attempt to make all vertices non-negative by repeatedly applying either of these steps:

* A given node requests aid. Each adjacent node gives it one dollar.
* A given node offers aid. It gives each adjacent node one dollar.

Once all nodes have non-zero values the game is won.

## Running the Game

```zsh
go run dollarGame.go | gvpack -u | dot -Tpng > test.png && open test.png
```

## Play Notes

Not all starting configurations are solvable. If the following inequality is true then the configuration will be solvable (it is not clear whether the converse is true):

Total of all values >= EdgeCount - VertexCount + 1

## Resources

* [Numberphile Video](https://www.youtube.com/watch?v=U33dsEcKgeQ) introducing and explaining the game
* The math behind [the Dollar Game](https://drive.google.com/file/d/1RMmfJ_0E6Gy-R6a59mkR_donKzSNHxJq/view)
* [Riemann-Roch](https://mattbaker.blog/2013/10/18/riemann-roch-for-graphs-and-applications/)
