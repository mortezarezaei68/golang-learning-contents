# Coding Challenge - Maximize Shelf Stability Score

Since the objective was to maximize the score, defined as the sum of each product's weight multiplied by
the level it's assigned to, I decided to start putting heaviest items on the top shelf because they contribute
significantly more to the score when placed on a higher shelf.

Therefore, I first sorted products in descending order of weight and shelves in descending order of level, defind as
`model.Level.Order` in the app. Then, starting from the heaviest product, I tried to place it on the highest shelf
where it fit without exceeding capacity. If it couldn't fit, I moved it down to the next shelf, repeating this process
until a valid placement was found. If a valid placement couldn't be found for a product, an error was returned stating that
`there isn't enough capacity to place all the products`.

Finally, a list of products was achieved which was used to calculate the final score, which was supposed to be the
maximum score possible without breaking constraints.

# Edit 1
Before, I was assigning heaviest products to the top shelves. Now, because a new constraint has been issued, I've changed
the shelves sorting based on their capacities not their orders. Now, heaviest products are assigned to the shelves with
more capacities.

# Edit 2
Now, I am assigning products based on the available capacity of the shelves. If the remaining capacity of a shelf is higher
than the capacity of the next shelf (shelves are sorted in descenidng order of capacity), it will be assigned to that shelf.
If not, it will be assigned to the next shelf. This process is repeated for other products. Heavier products will find larger
remaining capacities and reside there.